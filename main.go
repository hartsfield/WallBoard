package main // viewData represents the root model used to dynamically update the app

import (
	"context"
	"fmt"
	"log"
	"math/rand"
	"os"
	"time"

	"github.com/redis/go-redis/v9"
)

var (
	// connect to redis
	redisIP = os.Getenv("redisIP")
	rdb     = redis.NewClient(&redis.Options{
		Addr:     redisIP + ":6379",
		Password: "",
		DB:       1,
	})

	// this context is used for the client/server connection. It's useful
	// for passing the token/credentials around.
	rdx = context.Background()

	postDBChron []*post
	postDBRank  []*post
)

func init() {
	log.SetFlags(log.LstdFlags | log.Lshortfile)
	rand.Seed(time.Now().UTC().UnixNano())
}

func main() {
	if len(logFilePath) > 1 {
		logFile := setupLogging()
		defer logFile.Close()
	}
	beginCache()

	ctx, srv := bolt()

	fmt.Println("Server started @ http://localhost" + srv.Addr)
	log.Println("Server started @ " + srv.Addr)

	<-ctx.Done()
}
