package main

import (
	"fmt"
	"log"
	"math/rand"
	"net/http"
)

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/hello", hello)
	mux.HandleFunc("/ping", ping)
	mux.HandleFunc("/v1/test1", test1)
	mux.HandleFunc("/v1/test2", test2)

	if err := http.ListenAndServe(":80", mux); err != nil {
		log.Fatal(err)
	}
}

func hello(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello World! %d, token %s, user id %s", rand.Int(), r.Header.Get("Authorization"), r.Header.Get("x-subject-id"))
}

func ping(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "pong")
}

func test1(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "/v1/test1")
}

func test2(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "/v1/test2")
}
