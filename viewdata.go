package main

import "time"

type post struct {
	Title    string    `json:"title" redis:"title"`
	Id       string    `json:"id" redis:"id"`
	Author   string    `json:"author,name" redis:"author"`
	TS       time.Time `json:"ts" redis:"ts"`
	FTS      string    `json:"fts" redis:"fts"`
	BodyText string    `json:"bodytext" redis:"bodytext"`
	Nonce    string    `json:"nonce" redis:"nonce"`
	Children []*post   `json:"children" redis:"children"`
	Parent   string    `json:"parent" redis:"parent"`
}

type userData struct {
	Name string `json:"name" redis:"name"`
}

type viewData struct {
	ViewType    string `json:"viewType" redis:"viewType"`
	PageTitle   string
	CompanyName string
	Stream      []*post
	UserData    userData
	Nonce       string
}
