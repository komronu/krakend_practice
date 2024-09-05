package main

import (
	"fmt"
	"log"
	"math/rand"
	"net/http"
)

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/red", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Hello Red World! %d, token %s, user id %s", rand.Int(), r.Header.Get("Authorization"), r.Header.Get("x-subject-id"))
	})

	if err := http.ListenAndServe(":80", mux); err != nil {
		log.Fatal(err)
	}
}
