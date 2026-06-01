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
	mux.Handle("/app/", http.StripPrefix("/app", fileserver))
	mux.HandleFunc("/healthz", handleReadiness)

	srv := &http.Server{
		Addr:    ":" + port,
		Handler: mux,
	}

	log.Printf("Serving from %s on port: %s\n", fileroot, port)
	log.Fatal(srv.ListenAndServe())
}

func handleReadiness(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("Content-Type", "text/plain; charset=utf-8")
	w.WriteHeader(http.StatusOK)
	w.Write([]byte(http.StatusText(http.StatusOK)))
}
