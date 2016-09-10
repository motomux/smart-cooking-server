package main

import (
	"fmt"
	"log"
	"net/http"
	"time"

	tarantool "github.com/tarantool/go-tarantool"
)

func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hi there, This is smart cooking server!")
}

func main() {

	server := "smart-cooking-db:3301"
	opts := tarantool.Opts{
		Timeout:       500 * time.Millisecond,
		Reconnect:     1 * time.Second,
		MaxReconnects: 3,
	}

	client, err := tarantool.Connect(server, opts)
	if err != nil {
		log.Fatalf("Failed to connect: %s", err.Error())
	}
	log.Println("Connected to tarantool")

	http.HandleFunc("/", handler)
	log.Fatalln(http.ListenAndServe(":80", nil))
}
