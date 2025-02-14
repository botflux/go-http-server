package main

import (
	"github.com/botflux/go-http-server/http"
	"log"
)

func main() {
	server := http.Server{
		ListenAddr: ":8887",
	}

	if err := server.Listen(); err != nil {
		log.Fatal(err)
	}
}
