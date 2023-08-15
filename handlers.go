package main

import (
	"bytes"
	"encoding/json"
	"log"
	"net/http"
	"net/url"
	"strconv"
	"strings"
	"time"

	"github.com/redis/go-redis/v9"
)

func home(w http.ResponseWriter, r *http.Request) {
	var v viewData
	// var count int = 0
	v.Stream = postDBChron[:20]
	exeTmpl(w, r, &v, "main.tmpl")
}
func getByChron(w http.ResponseWriter, r *http.Request) {
	// var v viewData
	// // var count int = 0
	// v.Stream = postDBChron
	// exeTmpl(w, r, &v, "main.tmpl")
	var perPage int = 20
	var v viewData
	var count int = 20
	log.Println(strings.Split(r.RequestURI, "?"))
	if len(strings.Split(r.RequestURI, "?")) > 1 {
		params, err := url.ParseQuery(strings.Split(r.RequestURI, "?")[1])
		if err != nil {
			log.Println(err)
		}
		if params["count"] == nil {
			params["count"] = append(params["count"], "0")
		}

		if params["count"][0] != "None" {
			count, err = strconv.Atoi(params["count"][0])
			if err != nil {
				log.Println(err)
			}

			var nextCount string
			if len(postDBChron) < count+perPage {
				v.Stream = postDBChron[len(postDBChron)-(count+perPage-len(postDBChron)):]
				// if len(postDBChron) < count {
				// }
				nextCount = "None"
			} else {
				v.Stream = postDBChron[count-1 : count+perPage]
				log.Println(v.Stream)
				nextCount = strconv.Itoa(count + perPage)
			}
			var bb bytes.Buffer
			err = templates.ExecuteTemplate(&bb, "stream.tmpl", v)
			if err != nil {
				log.Println(err)
			}
			ajaxResponse(w, map[string]string{
				"success":  "true",
				"template": bb.String(),
				"count":    nextCount,
			})
		}
	} else {
		if len(postDBChron) < perPage {
			v.Stream = postDBRank[:]
		} else {
			v.Stream = postDBChron[:perPage]
		}
		exeTmpl(w, r, &v, "main.tmpl")
	}

}
func getByRanked(w http.ResponseWriter, r *http.Request) {
	var perPage int = 20
	var v viewData
	var count int = 20
	log.Println(strings.Split(r.RequestURI, "?"))
	if len(strings.Split(r.RequestURI, "?")) > 1 {
		params, err := url.ParseQuery(strings.Split(r.RequestURI, "?")[1])
		if err != nil {
			log.Println(err)
		}
		if params["count"] == nil {
			params["count"] = append(params["count"], "0")
		}

		if params["count"][0] != "None" {
			count, err = strconv.Atoi(params["count"][0])
			if err != nil {
				log.Println(err)
			}

			var nextCount string
			if len(postDBRank) < count && len(postDBRank) < count+perPage {
				v.Stream = postDBRank[len(postDBRank)-(count-len(postDBRank)):]
				nextCount = "None"
			} else {
				v.Stream = postDBRank[count+1 : count+perPage]
				log.Println(v.Stream)
				nextCount = strconv.Itoa(count + perPage)
			}
			var bb bytes.Buffer
			err = templates.ExecuteTemplate(&bb, "stream.tmpl", v)
			if err != nil {
				log.Println(err)
			}
			ajaxResponse(w, map[string]string{
				"success":  "true",
				"template": bb.String(),
				"count":    nextCount,
			})
		}
	} else {
		if len(postDBRank) < perPage {
			v.Stream = postDBRank[:]
		} else {
			v.Stream = postDBRank[:perPage]
		}
		exeTmpl(w, r, &v, "main.tmpl")
	}
}

func viewPost(w http.ResponseWriter, r *http.Request) {
	parts := strings.Split(r.RequestURI, "/")
	var p post
	rdb.HGetAll(rdx, parts[len(parts)-1]).Scan(&p)
	log.Println(len(p.Id))
	if len(p.Id) == 11 {
		getAllChidren(&p, "RANK")
	} else {
		p.BodyText = "This post was automatically deleted."
	}
	var v viewData
	v.Stream = nil
	v.Stream = append(v.Stream, &p)
	v.ViewType = "post"
	exeTmpl(w, r, &v, "post.tmpl")

}

func handleForm(w http.ResponseWriter, r *http.Request) {
	data, err := marshalPostData(r)
	if err != nil {
		log.Println(err)
	}
	log.Println(data.Parent)
	parentExists, err := rdb.Exists(rdx, data.Parent).Result()
	if err != nil {
		log.Println(err)
	}

	if parentExists == 0 && data.Parent != "root" {
		log.Println(" no parent")
		ajaxResponse(w, map[string]string{
			"success":   "false",
			"replyID":   "",
			"timestamp": data.FTS,
		})
		return
	}
	if len(data.Title) > 1 {
		// return
	}
	if len(data.BodyText) < 5 || len(data.BodyText) > 1000 {
		ajaxResponse(w, map[string]string{"success": "false"})
		return
	}
	data.Id = genPostID(10)
	data.TS = time.Now()
	data.FTS = data.TS.Format("2006-01-02 03:04:05 pm")
	rdb.HSet(
		rdx, data.Id,
		"name", data.Author,
		"title", data.Title,
		"bodytext", data.BodyText,
		"id", data.Id,
		"ts", data.TS,
		"fts", data.FTS,
		"parent", data.Parent,
		"childCount", "0",
	)
	if data.Parent != "root" {
		rdb.ZAdd(rdx, data.Parent+":CHILDREN:CHRON", redis.Z{Score: float64(time.Now().UnixMilli()), Member: data.Id})
		rdb.ZAdd(rdx, data.Parent+":CHILDREN:RANK", redis.Z{Score: 0, Member: data.Id})
		bubbleUp(data)
	} else {
		rdb.ZAdd(rdx, "ANON:POSTS:CHRON", redis.Z{Score: float64(time.Now().UnixMilli()), Member: data.Id})
		rdb.ZAdd(rdx, "ANON:POSTS:RANK", redis.Z{Score: 0, Member: data.Id})
		popLast()
	}
	ajaxResponse(w, map[string]string{
		"success":   "true",
		"replyID":   data.Id,
		"timestamp": data.FTS,
	})
	beginCache()
}

func marshalPostData(r *http.Request) (*post, error) {
	t := &post{}
	decoder := json.NewDecoder(r.Body)
	defer r.Body.Close()
	err := decoder.Decode(t)
	if err != nil {
		return t, err
	}
	return t, nil
}
