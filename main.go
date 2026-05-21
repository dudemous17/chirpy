package main

import (
	"log"
	"net/http"
)

func main() {
	const fileroot = "."
	const port = "8080"

	mux := http.NewServeMux()
	fileserver := http.FileServer(http.Dir(fileroot))
	mux.Handle("/", fileserver)

	srv := &http.Server{
		Addr:    ":" + port,
		Handler: mux,
	}

	log.Printf("Serving from %s on port: %s\n", fileroot, port)
	log.Fatal(srv.ListenAndServe())
}
