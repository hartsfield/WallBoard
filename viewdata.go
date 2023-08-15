package main

import "time"

// post is the structure of a user post. Posts are created by users and stored
// in redis.
type post struct {
	Title      string    `json:"title" redis:"title"`
	Id         string    `json:"id" redis:"id"`
	Author     string    `json:"author,name" redis:"author"`
	TS         time.Time `json:"ts" redis:"ts"`
	FTS        string    `json:"fts" redis:"fts"`
	BodyText   string    `json:"bodytext" redis:"bodytext"`
	Nonce      string    `json:"nonce" redis:"nonce"`
	Children   []*post   `json:"children" redis:"children"`
	ChildCount int       `json:"childCount" redis:"childCount"`
	Parent     string    `json:"parent" redis:"parent"`
	PostCount  string    `json:"postCount" redis:"postCount"`
}

// viewData represents the root model used to dynamically update the app
type viewData struct {
	ViewType    string `json:"viewType" redis:"viewType"`
	PageTitle   string
	CompanyName string
	Stream      []*post
	Nonce       string
	Order       string `json:"order" redis:"order"`
}
