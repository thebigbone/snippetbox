package main

import (
	"flag"
	"log"
	"net/http"
	"os"

	"github.com/fatih/color"
)

var infoLog = log.New(os.Stdout, color.GreenString("[INFO]\t"), log.Ldate|log.Ltime)

var errorlog = log.New(os.Stderr, color.RedString("[ERROR]\t"), log.Ldate|log.Ltime|log.Lshortfile)

func main() {
	mux := http.NewServeMux()

	address := flag.String("p", ":8080", "port to listen on")
	flag.Parse()

	fileserver := http.FileServer(http.Dir("./ui/static/"))

	mux.HandleFunc("/", home)

	mux.Handle("/static/", http.StripPrefix("/static/", fileserver))
	mux.HandleFunc("/snippet/view", snippetView)
	mux.HandleFunc("/snippet/create", snippetCreate)

	infoLog.Printf(color.GreenString("Starting server at localhost %s", *address))

	err := http.ListenAndServe(*address, mux)
	errorlog.Fatal(color.RedString("%s", err))
}
