package main

import (
	"encoding/json"
	"log"
	"net/http"
	"strings"
	"time"

	"github.com/redis/go-redis/v9"
)

func home(w http.ResponseWriter, r *http.Request) {
	var v viewData
	v.Stream = postDBChron
	exeTmpl(w, r, &v, "main.tmpl")
}

func getByRanked(w http.ResponseWriter, r *http.Request) {
	var v viewData
	v.Stream = postDBRank
	exeTmpl(w, r, &v, "main.tmpl")
}

func viewPost(w http.ResponseWriter, r *http.Request) {
	parts := strings.Split(r.RequestURI, "/")
	var p post
	rdb.HGetAll(rdx, parts[len(parts)-1]).Scan(&p)
	getAllChidren(&p, "RANK")
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
		// popLast()
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
