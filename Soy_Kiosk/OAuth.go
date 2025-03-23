package main

import (
    "fmt"
	"net/http"
	"log"
	"encoding/json"
	"os"
)

type haConfig struct {
	home_assistant_location string `json:"key1"`
    allow_outside_connections bool `json:"key2"`
    default_dashboard string `json:"key3"`
    kiosk_account_access_token string `json:"key4"`
}


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
		var data haConfig
		err = json.Unmarshal(content, &data)
		dashboard := data.default_dashboard
	}
    fmt.Fprintf(w, "<h1>%s</h1><div>%s</div>", title, dashboard)
}

func main() {
    http.HandleFunc("/Kiosk/", oAuthHandler)
    log.Fatal(http.ListenAndServe(":8080", nil))
}

//Future OAuth Stuff
	//need authorize
	//need token
	//need user info