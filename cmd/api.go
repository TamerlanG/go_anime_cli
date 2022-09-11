package cmd

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

type Response struct {
	Data       []Anime     `json:"data"`
}

type Anime struct {
	Title         string   `json:"title"`
	Type          string   `json:"type"`
	Episodes      int      `json:"episodes"`
	Status        string   `json:"status"`
}

type Config struct {
  Limit string
}

func Search(searchQuery string, config Config) []Anime { client := http.Client{}
  request, error := http.NewRequest("GET", "https://api.jikan.moe/v4/anime", nil)
  
  query := request.URL.Query()
  
  query.Add("q", searchQuery)
  query.Add("sfw", "true")
  query.Add("limit", config.Limit)
  
  request.URL.RawQuery = query.Encode()

  if error != nil {
    log.Fatal(error)
  } 
  
  res, error := client.Do(request)
  
  if error != nil {
    log.Fatal(error)
  }

  defer res.Body.Close()

  // Read the Response Body
  body, error := ioutil.ReadAll(res.Body)
  
  if error != nil {
    log.Fatal(error)
  }

  var result Response
  if err := json.Unmarshal(body, &result); err != nil {   
    fmt.Println(err)
  }
 
  return result.Data

}
 
