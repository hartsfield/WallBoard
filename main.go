package main

import (
	"fmt"
	"log"
	"math/rand"
	"time"
)

// init sets up formatting for logging, and seeds math/rand for generating post
// IDs
func init() {
	log.SetFlags(log.LstdFlags | log.Lshortfile)
	rand.Seed(time.Now().UTC().UnixNano())
}

func main() {
	if len(logFilePath) > 1 {
		logFile := setupLogging()
		defer logFile.Close()
	}

	// cache the database
	beginCache()

	ctx, srv := bolt()

	fmt.Println("Server started @ http://localhost" + srv.Addr)
	log.Println("Server started @ http://localhost" + srv.Addr)

	<-ctx.Done()
}
