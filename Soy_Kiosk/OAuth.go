package main

import (
    "fmt"
	"net/http"
	"log"
	"encoding/json"
	"os"
	"regexp"
	"io"
	/*"bytes"*/
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

func relURLtoAbsURL(inStr string) string {
	fmt.Println(inStr)
	return "http://homeassistant.local:8123" + inStr
}

func oAuthHandler(w http.ResponseWriter, r *http.Request) {
	//See if local ip
	//use ip, _, err := net.SplitHostPort(r.RemoteAddr)


    //title := r.URL.Path[len("/Kiosk/"):]
	//queryParams := r.URL.Query()
	var dashboard string 
	//dashboard = string.Join(queryParams["dashboard"])
	if( dashboard == "" ) {
		content, err := os.ReadFile("/data/options.json")
		if err != nil {
			fmt.Println("Error Parsing Json", err)
			os.Exit(1)
		}
		var data haConfig
		json.Unmarshal(content, &data)
		dashboard = data.default_dashboard
	}
	
	//Send Dashboard
	url := "http://homeassistant.local:8123/dashboard-kiosk/" // Replace with the desired URL
	//get page
	resp, err := http.Get(url)
	if err != nil {
	   //handle error or something
	}
	
	//We Read the response body on the line below.
	body, err := io.ReadAll(resp.body)
	if err != nil {
		//handle error
	}
	//Convert the body to type string
	sb := string(body)


	//strip side bar
	stripBar, err := regexp.Compile(`<ha-sidebar>*</ha-sidebar>`)
	if err != nil {
      // handle error
    }
	sb = stripBar.ReplaceAllString(sb, " ")
	
	//rewrite page with absolute links
	re, err := regexp.Compile(`/^[^\/]+\/[^\/].*$|^\/[^\/].*`)
    if err != nil {
      // handle error
    }
	sb = re.ReplaceAllStringFunc(sb, relURLtoAbsURL)
	
	
	//write it out
	fmt.Fprintf(w,"%s",sb)
	
}

func main() {
    http.HandleFunc("/Kiosk/", oAuthHandler)
    log.Fatal(http.ListenAndServe(":8080", nil))
}

//Future OAuth Stuff
	//need authorize
	//need token
	//need user info