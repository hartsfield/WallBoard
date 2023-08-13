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

	postDB []*post
)

func init() {
	log.SetFlags(log.LstdFlags | log.Lshortfile)
	rand.Seed(time.Now().UTC().UnixNano())
}

func buildDB() {
	var ids []string
	opts := &redis.ZRangeBy{
		Min:    "-inf",
		Max:    "+inf",
		Offset: 0,
		Count:  -1,
	}
	err := rdb.ZRevRangeByScore(rdx, "ANON:POSTS:CHRON", opts).ScanSlice(&ids)
	if err != nil {
		log.Println(err)
	}
	for _, id := range ids {
		var p post
		rdb.HGetAll(rdx, id).Scan(&p)
		getAllChidren(&p)
		postDB = append(postDB, &p)
	}
}

func getAllChidren(po *post) {
	var ids []string
	opts := &redis.ZRangeBy{
		Min:    "-inf",
		Max:    "+inf",
		Offset: 0,
		Count:  -1,
	}
	err := rdb.ZRevRangeByScore(rdx, po.Id+":CHILDREN", opts).ScanSlice(&ids)
	if err != nil {
		log.Println(err)
	}

	for _, id := range ids {
		var p post
		rdb.HGetAll(rdx, id).Scan(&p)
		getAllChidren(&p)
		po.Children = append(po.Children, &p)
	}
}

func main() {
	if len(logFilePath) > 1 {
		logFile := setupLogging()
		defer logFile.Close()
	}

	ctx, srv := bolt()

	fmt.Println("Server started @ http://localhost" + srv.Addr)
	log.Println("Server started @ " + srv.Addr)

	buildDB()
	<-ctx.Done()
}
