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

func Search(query string) []Anime {
  url := fmt.Sprintf("https://api.jikan.moe/v4/anime?sfw=true&q=%s", query);

  res, error := http.Get(url);

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
 
