package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
)

func main() {
	port := os.Getenv("PORT")

	if port == "" {
		fmt.Println(`"PORT" env not found. switching to default`)
		port = "7060"
	}

	servePath := os.Getenv("SERVE_PATH")

	if servePath == "" {
		panic(`"SERVE_PATH" env not found.`)
	}

	fs := http.FileServer(http.Dir(servePath))
	http.Handle("/", fs)

	log.Println("Listening server on :" + port)
	err := http.ListenAndServe(":"+port, nil)
	if err != nil {
		log.Fatal(err)
	}
}
