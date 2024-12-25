package main

import (
	"context"
	"encoding/json"
	"html/template"
	"log"
	"os"

	"github.com/redis/go-redis/v9"
)

// ckey/ctxkey is used as the key for the HTML context and is how we retrieve
// token information and pass it around to handlers
// type ckey int

// const ctxkey ckey = iota

type env map[string]string

type config struct {
	App    app    `json:"app"`
	GCloud gcloud `json:"gcloud"`
}

type app struct {
	Name       string `json:"name"`
	DomainName string `json:"domain_name"`
	Version    string `json:"version"`
	Env        env    `json:"env"`
	Port       string `json:"port"`
	AlertsOn   bool   `json:"alertsOn"`
	TLSEnabled bool   `json:"tls_enabled"`
	Repo       string `json:"repo"`
}

type gcloud struct {
	Command   string `json:"command"`
	Zone      string `json:"zone"`
	Project   string `json:"project"`
	User      string `json:"user"`
	LiveDir   string `json:"livedir"`
	ProxyConf string `json:"proxyConf"`
}

var (
	// servicePort is the port this program will run on

	appConf     *config            = readConf()
	servicePort                    = ":" + appConf.App.Port
	logFilePath                    = appConf.App.Env["logFilePath"]
	templates   *template.Template = template.New("")
	AppName     string             = appConf.App.Name
	companyName string             = appConf.App.Name
	// connect to redis
	redisIP = os.Getenv("redisIP")
	rdb     = redis.NewClient(&redis.Options{
		Addr:     redisIP + ":6379",
		Password: "",
		DB:       1,
	})
	// redis context
	rdx = context.Background()

	// Database caches
	postDBChron []*post
	postDBRank  []*post
)

func readConf() *config {
	b, err := os.ReadFile("./bolt.conf.json")
	if err != nil {
		log.Println(err)
	}
	c := config{}
	err = json.Unmarshal(b, &c)
	if err != nil {
		log.Println(err)
	}
	return &c
}
