package main

import (
	"fmt"
	"net/http"
	"encoding/json"
    "log"
    "strconv"

	"github.com/JaverSingleton/torrents-downloader/tparser"
)

func search(w http.ResponseWriter, r *http.Request) {
	var query string
	if array, ok := r.URL.Query()["query"]; ok && len(array) > 0 {    
		query = array[0]
	}
	var priority string
	if array, ok := r.URL.Query()["priority"]; ok && len(array) > 0 {    
		priority = array[0]
	}
	payload := struct {
		Torrents []tparser.Torrent `json:"torrents"`
	} {}
	priorityNumber, err := strconv.Atoi(priority)
	torrents, err := tparser.Find(query, priorityNumber)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	payload.Torrents = torrents
    log.Println("Torrents:", payload)

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

	http.HandleFunc("/search", search)

	http.ListenAndServe(":3000", nil)
}