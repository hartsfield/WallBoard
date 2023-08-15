package main

import (
	"context"
	"html/template"
	"os"

	"github.com/redis/go-redis/v9"
)

// ckey/ctxkey is used as the key for the HTML context and is how we retrieve
// token information and pass it around to handlers
type ckey int

const ctxkey ckey = iota

var (
	// servicePort is the port this program will run on
	servicePort                    = ":" + os.Getenv("servicePort")
	logFilePath                    = os.Getenv("logFilePath")
	templates   *template.Template = template.New("main")
	companyName string             = "BoltApp"
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
