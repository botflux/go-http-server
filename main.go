package main

import (
	"github.com/botflux/go-http-server/http"
	"github.com/botflux/go-http-server/routing"
	"log"
)

func main() {
	r := &routing.Router{}

	r.Add("GET", "/", func(request http.Request) http.Response {
		return http.Response{
			StatusCode: 200,
			Headers: map[string]string{
				"Content-Type": "text/html; charset=utf-8",
			},
			Body: "<html><body><h1>Hello world</h1></body></html>",
		}
	})

	server := http.Server{
		ListenAddr: ":8887",
		Router:     r,
	}

	if err := server.Listen(); err != nil {
		log.Fatal(err)
	}
}
