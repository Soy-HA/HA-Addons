package main

import (
    "fmt"
    "os"
	"net/http"
)

func handler(w http.ResponseWriter, r *http.Request) {
    fmt.Fprintf(w, "Hi there, I love %s!", r.URL.Path[1:])
}

func oAuthHandler(w http.ResponseWriter, r *http.Request) {
    title := r.URL.Path[len("/OAuth/"):]
    p, err := loadPage(title)
	if err != nil {
        http.Redirect(w, r, "/edit/"+title, http.StatusFound)
        return
    }
    fmt.Fprintf(w, "<h1>%s</h1><div>%s</div>", p.Title, p.Body)
}

func main() {
    http.HandleFunc("/", handler)
    http.HandleFunc("/OAuth/", oAuthHandler)
    log.Fatal(http.ListenAndServe(":8080", nil))
}