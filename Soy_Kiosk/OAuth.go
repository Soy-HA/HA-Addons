package main

import (
    "fmt"
	"net/http"
	"log"
	"encoding/json"
	"os"
)

func handler(w http.ResponseWriter, r *http.Request) {
    w.WriteHeader(http.StatusNotFound)
	fmt.Fprint(w, "Page not found")
}

func oAuthHandler(w http.ResponseWriter, r *http.Request) {
	//See if local ip
	//use ip, _, err := net.SplitHostPort(r.RemoteAddr)


    title := r.URL.Path[len("/Kiosk/"):]
	queryParams := r.URL.Query()
	dashboard := queryParams["dashboard"]
	if( dashboard == nil ) {
		//change this to pull in from config file
		content, err := os.ReadFile("/data/options.json")
		var data Data
		err = json.Unmarshal(content, &data)
		dashboard := data.default_dashboard
	}
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