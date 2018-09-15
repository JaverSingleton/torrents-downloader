package tparser

import (
	"net/http"
	"encoding/json"
    "net/url"
    "io/ioutil"
    "strconv"
	"log"
    "time"
    "reflect"
)

func Find(query string) ([]Torrent, error) { 
	result, err := find(query)
    if err != nil {
    	return []Torrent {}, err
    }
    
	return result, nil
}

func find(query string) (TparserResult, error) {
	var Url *url.URL
    Url, err := url.Parse("http://tparser.org")
    if err != nil {
    	return TparserResult {}, err
    }

    Url.Path += "/" + query

	req, err := http.NewRequest("GET", Url.String(), nil)
	if (err != nil) {
    	return TparserResult {}, err
	}
	client := &http.Client {}

	response, err := client.Do(req)
    if err != nil {
    	return TparserResult {}, err
    }
    defer response.Body.Close()

    contents, err := ioutil.ReadAll(response.Body)
    if err != nil {
    	return TparserResult {}, err
    }
    var result = TparserResult {}

	if err = json.Unmarshal(contents, &result); err != nil {
		return TparserResult {}, err
	}
	return result, nil
}

type Torrent struct {
	Id string `json:"id"`
    Name string `json:"name"`
    Seed string `json:"seed"`
    Leech string `json:"leech"`
    Size string `json:"size"` 
    Magnet string `json:"magnet"`
}

type TparserResult struct {
    Items []TparserItem `json:"sr"`
}

type TparserItem struct {
    Id string `json:"d"`
    Img string `json:"img"`
    Name string `json:"name"`
    Seed string `json:"s"`
    Leech string `json:"l"`
    Size string `json:"size"` 
    Unit string `json:"t"` 
}
