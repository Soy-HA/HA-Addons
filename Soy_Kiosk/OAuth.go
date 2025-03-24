package main

import (
    "fmt"
	"net/http"
	"log"
	"encoding/json"
	"os"
	"io"
	"bytes"
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
	//Authethenicate user
	access_token := "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJpc3MiOiI2YzQ4MmI1MTIxZmU0NGVhOTM2ZDQxOTRkMWEzZWExZSIsImlhdCI6MTc0Mjc3Nzk5NSwiZXhwIjoyMDU4MTM3OTk1fQ.MlljcluzEDUmMQgGoB6-IjOu16aJmB-MuMF2_nNgTLA"
	url := "http://homeassistant.local:8123/api/websocket"
	
	
	jsonStr := map[string]string{"type": "auth", "acess_token": access_token}
	jsonValue, _ := json.Marshal(values)
	req, err := http.NewRequest("POST", url, bytes.NewBuffer(jsonStr))
	//req.Header.Set("X-Custom-Header", "myvalue")
    req.Header.Set("Content-Type", "application/json")

    client := &http.Client{}
    resp, err := client.Do(req)
    if err != nil {
        panic(err)
    }
    defer resp.Body.Close()

    fmt.Println("response Status:", resp.Status)
    fmt.Println("response Headers:", resp.Header)
    body, _ := io.ReadAll(resp.Body)
    fmt.Println("response Body:", string(body))
	
	
	
	//Send Dashboard
		//TEMP for funsies. Just seeing what happens if I jsut throw the dashboard in here lololol
	url := "http://homeassistant.local:8123/" // Replace with the desired URL

	fmt.Fprintf(w,"<iframe src='%s' title='ha-main-window' id='dashboard' name='dashboard'></iframe>",url)
	
}

func main() {
    http.HandleFunc("/Kiosk/", oAuthHandler)
    log.Fatal(http.ListenAndServe(":8080", nil))
}

//Future OAuth Stuff
	//need authorize
	//need token
	//need user info