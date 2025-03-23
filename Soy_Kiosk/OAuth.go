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
    title := r.URL.Path[len("/Kiosk/"):]
	queryParams := r.URL.Query()
	
	
    fmt.Fprintf(w, "<h1>%s</h1><div>%s</div>", title, queryParams["dashboard"])
}

func main() {
    http.HandleFunc("/Kiosk/", oAuthHandler)
    log.Fatal(http.ListenAndServe(":8080", nil))
}

//Future OAuth Stuff
	//need authorize
	//need token
	//need user info