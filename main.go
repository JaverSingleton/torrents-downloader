package main

import (
	"fmt"
	"net/http"
	"encoding/json"
    "log"
    "strconv"

	"github.com/JaverSingleton/torrents-downloader/tparser"
	"github.com/JaverSingleton/torrents-downloader/downloader"
)

func search(w http.ResponseWriter, r *http.Request) {
	var query string
	if array, ok := r.URL.Query()["query"]; ok && len(array) > 0 {    
		query = array[0]
	}
	var amount string
	if array, ok := r.URL.Query()["amount"]; ok && len(array) > 0 {    
		amount = array[0]
	}
	payload := struct {
		Torrents []tparser.Torrent `json:"torrents"`
	} {}
	priorityNumber, err := strconv.Atoi(amount)
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

func download(w http.ResponseWriter, r *http.Request) {
	var magnet string
	if array, ok := r.URL.Query()["magnet"]; ok && len(array) > 0 {    
		magnet = array[0]
	}
	var target string
	if array, ok := r.URL.Query()["target"]; ok && len(array) > 0 {    
		target = array[0]
	}
	payload := struct {
		Result string `json:"result"`
	} {}
	link, err := tparser.GetLink(magnet)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	downloadErr := downloader.Download(link, target)
	if downloadErr != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	payload.Result = "Ok"
    log.Println("Link:", link)

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
	http.HandleFunc("/download", download)

	http.ListenAndServe(":3000", nil)
}