package main

import (
	"fmt"
	"log"
	"math/rand"
	"net/http"
)

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/green", func(writer http.ResponseWriter, request *http.Request) {
		fmt.Fprintf(writer, "Hello Green World! %d, token: %s", rand.Int(), request.Header.Get("Authorization"))
	})

	if err := http.ListenAndServe(":80", mux); err != nil {
		log.Fatal(err)
	}
}
