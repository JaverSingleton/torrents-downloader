package resource

import (
    "encoding/json"
    "io/ioutil"
)

const (
  Resources_All = 9999999
  Resources_NormalPriority = 4
  Resources_HighPriority = 2
)

func GetResources(priority int) ([]Resource, error) {
    byteArray, err := ioutil.ReadFile("resources.json")
    if err != nil {
        return []Resource {}, err
    }
    var resources = Resources {}
    if err = json.Unmarshal(byteArray, &resources); err != nil {
        return []Resource {}, err
    }

    var maxSize int
    if priority > len(resources.Resources) - 1 {
        maxSize = len(resources.Resources) - 1
    } else {
        maxSize = priority
    }

    return resources.Resources[:maxSize], nil
}
type Resources struct {
    Resources []Resource `json:"resources"`
}

type Resource struct {
    Url string `json:"url"`
    Name string `json:"name"`
}
