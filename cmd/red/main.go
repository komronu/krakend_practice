package main

import (
	"fmt"
	"log"
	"math/rand"
	"net/http"
)

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/red", func(writer http.ResponseWriter, request *http.Request) {
		fmt.Fprintf(writer, "Hello Red World! %d, token %s", rand.Int(), request.Header.Get("Authorization"))
	})

	if err := http.ListenAndServe(":80", mux); err != nil {
		log.Fatal(err)
	}
}
