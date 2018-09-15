package main

import (
	"fmt"
	"net/http"
	"encoding/json"

	"github.com/JaverSingleton/torrents-downloader/tparser"
)

func search(w http.ResponseWriter, r *http.Request) {
	var query string
	if array, ok := r.URL.Query()["query"]; ok && len(array) > 0 {    
		code = array[0]
	}
	payload := struct {
		torrents string `json:"torrents"`
	} {}
	payload.torrents = tparser.find(query)

	js, err := json.Marshal(payload)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
  	w.Header().Set("Content-Type", "application/json")
	w.Write(js)
}

func main() {
	fmt.Println("Listening on port :3000")

	http.Handle("/assets/", http.StripPrefix("/assets/", http.FileServer(http.Dir("./assets/"))))
	http.HandleFunc("/search", search)

	http.ListenAndServe(":3000", nil)
}