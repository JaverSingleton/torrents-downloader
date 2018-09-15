package tparser

import (
    "encoding/json"
    "io/ioutil"
)

const (
  Resources_All = 0
  Resources_NormalPriority = 4
  Resources_HighPriority = 2
)

func GetResources(priority int) ([]Resource, error) {
    byteArray, err := ioutil.ReadFile("resources.json")
    if err != nil {
        return []Resource {}, err
    }
    var resources = []Resource {}
    if err = json.Unmarshal(byteArray, &resources); err != nil {
        return []Resource {}, err
    }

    var maxSize int
    if priority == Resources_All {
        maxSize = len(resources)
    } else {
        maxSize = priority
    }

    return resources[:maxSize], nil
}


type Resource struct {
    Url string `json:"url"`
    Name int `json:"name"`
}
