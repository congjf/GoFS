package main

import (
	"log"
	"net/http"
	"time"
)

func main() {
	log.Println("GoFS:")

	var router Router

	server := &http.Server{
		Addr:           ":8888",
		Handler:        router,
		ReadTimeout:    10 * time.Second,
		WriteTimeout:   10 * time.Second,
		MaxHeaderBytes: 1 << 20,
	}
	log.Fatal(server.ListenAndServe())
}

type Router struct {
}

func (r Router) ServeHTTP(rw http.ResponseWriter, req *http.Request) {
	h := http.FileServer(http.Dir("."))
	h.ServeHTTP(rw, req)
}
