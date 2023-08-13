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
	v.Stream = postDB
	exeTmpl(w, r, &v, "main.tmpl")
}

func viewPost(w http.ResponseWriter, r *http.Request) {
	parts := strings.Split(r.RequestURI, "/")
	var p post
	rdb.HGetAll(rdx, parts[len(parts)-1]).Scan(&p)
	getAllChidren(&p)
	// for _, id := range childrenIDs {
	// 	var po post
	// 	rdb.HGetAll(rdx, id).Scan(&po)
	// 	p.Children = append(p.Children, getAllChidren(&po))
	// }
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
	if len(data.Title) > 1 {
		// return
	}
	if len(data.BodyText) < 5 || len(data.BodyText) > 1000 {
		ajaxResponse(w, map[string]string{"success": "false"})
		return
	}
	data.Id = genPostID(10)
	data.TS = time.Now()
	rdb.HSet(
		rdx, data.Id,
		"name", data.Author,
		"title", data.Title,
		"bodytext", data.BodyText,
		"id", data.Id,
		"ts", data.TS,
		"parent", data.Parent,
	)
	log.Println(data.Parent)
	if data.Parent != "root" {
		rdb.ZAdd(rdx, data.Parent+":CHILDREN", redis.Z{Score: 0, Member: data.Id})
	} else {

		postDB = append(postDB, data)
		rdb.ZAdd(rdx, "ANON:POSTS:CHRON", redis.Z{Score: 0, Member: data.Id})
	}
	ajaxResponse(w, map[string]string{"success": "true"})
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
