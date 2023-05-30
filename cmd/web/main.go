package main

import (
	"flag"
	"log"
	"net/http"
)

func main() {
	mux := http.NewServeMux()

	address := flag.String("p", ":8080", "port to listen on")
	flag.Parse()

	fileserver := http.FileServer(http.Dir("./ui/static/"))

	mux.HandleFunc("/", home)

	mux.Handle("/static/", http.StripPrefix("/static/", fileserver))
	mux.HandleFunc("/snippet/view", snippetView)
	mux.HandleFunc("/snippet/create", snippetCreate)
	log.Printf("Starting server on localhost%s", *address)
	err := http.ListenAndServe(*address, mux)
	log.Fatal(err)
}
