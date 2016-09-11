package main

import (
	"flag"
	"log"
	"net/http"
	"time"

	"github.com/motomux/smart-cooking-server/handler"
	tarantool "github.com/tarantool/go-tarantool"
)

func main() {
	port := flag.String("port", "80", "port of server")
	db := flag.String("db", "smart-cooking-db:3301", "host of db server")
	flag.Parse()

	opts := tarantool.Opts{
		Timeout:       500 * time.Millisecond,
		Reconnect:     1 * time.Second,
		MaxReconnects: 3,
	}

	client, err := tarantool.Connect(*db, opts)
	if err != nil {
		log.Fatalf("Failed to connect: %s, %s", err.Error(), *db)
	}
	log.Println("Connected to tarantool")

	env := &handler.Env{
		Client: client,
	}
	// Handler
	mux := handler.NewHandler(env)

	// Run server
	log.Println("Starting web server on", port)
	log.Fatalln(http.ListenAndServe(":"+*port, mux))
}
