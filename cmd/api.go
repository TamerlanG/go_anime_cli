package cmd

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

func Search(query string) string {
  url := fmt.Sprintf("https://api.jikan.moe/v4/anime?sfw=true&q=%s", query);

  fmt.Println(url);

  res, error := http.Get(url);

  if error != nil {
    log.Fatal(error)
  }

  // Read the Response Body
  body, error := ioutil.ReadAll(res.Body)

  if error != nil {
    log.Fatal(error)
  } 

  // Convert it to string 
  stringBody := string(body); 

  return stringBody;
}

