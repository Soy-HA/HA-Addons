package main

import (
    "fmt"
	"net/http"
	"log"
)

func handler(w http.ResponseWriter, r *http.Request) {
    w.WriteHeader(http.StatusNotFound)
	fmt.Fprint(w, "Page not found")
}

func oAuthHandler(w http.ResponseWriter, r *http.Request) {
    title := r.URL.Path[len("/OAuth/"):]
	//need authorize
	//need token
	//need user info
    fmt.Fprintf(w, "<h1>test</h1><div>%s</div>", title)
}

func main() {
    http.HandleFunc("/", handler)
    http.HandleFunc("/OAuth/", oAuthHandler)
    log.Fatal(http.ListenAndServe(":8080", nil))
}