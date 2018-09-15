package tparser

import (
	"net/http"
    "net/url"
    "io/ioutil"
    "strconv"
    "strings"
	"encoding/json"

	"github.com/k3a/html2text"
	"github.com/JaverSingleton/torrents-downloader/resource"
)



func Find(query string, priority int) ([]Torrent, error) { 
	resources, err := resource.GetResources(priority)
    if err != nil {
    	return []Torrent {}, err
    }

	var result []Torrent 
	for _, url := range resources {
		result = append(result, find(url, query)...)
	}

	return result
}

func find(url string, query string) []Torrent {
	result, err := runSearch(url, query)
    if err != nil {
    	return []Torrent {}
    }
	torrents := make([]Torrent, len(result.Items))
	for index, item := range result.Items {
		torrents[index] = convert(item)
	}

	return torrents
}

func runSearch(resourceUrl string, query string) (TparserResult, error) {
	var Url *url.URL
    Url, err := url.Parse(resourceUrl)
    if err != nil {
    	return TparserResult {}, err
    }

	Url.Path += ""
    parameters := url.Values{}
    parameters.Add("callback", "one")
    parameters.Add("jsonpx", query)
    parameters.Add("s", strconv.Itoa(1))
    Url.RawQuery = parameters.Encode()

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

    byteArray, err := ioutil.ReadAll(response.Body)
    if err != nil {
    	return TparserResult {}, err
    }
    content := string(byteArray[:])
    content = strings.TrimPrefix(content, "one(")
    content = strings.TrimSuffix(content, ")")
    content = strings.Replace(content, "\"", " ", -1)
    content = strings.Replace(content, "'", "\"", -1)

    var result = TparserResult {}

	if err = json.Unmarshal([]byte(content), &result); err != nil {
		return TparserResult {}, err
	}
	return result, nil
}

func convert(item TparserItem) Torrent {
	return Torrent {
		Id: item.Id,
		Name: html2text.HTML2Text(item.Name),
		Seed: item.Seed,
		Leech: item.Leech,
		Size: item.Size + " " + item.Unit,
		Magnet: strconv.Itoa(len(item.Img)) + item.Img + item.Id,
	}
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
