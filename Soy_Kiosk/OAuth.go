package main

import (
    "fmt"
	"net/http"
	"log"
)

func handler(w http.ResponseWriter, r *http.Request) {
    fmt.Fprintf(w, "Hi there, I love %s!", r.URL.Path[1:])
}

func oAuthHandler(w http.ResponseWriter, r *http.Request) {
    title := r.URL.Path[len("/OAuth/"):]
    fmt.Fprintf(w, "<h1>test</h1><div>test</div>")
}

func main() {
    http.HandleFunc("/", handler)
    http.HandleFunc("/OAuth/", oAuthHandler)
    log.Fatal(http.ListenAndServe(":8080", nil))
}