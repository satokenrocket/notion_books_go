package main

import (
  "fmt"
  "net/http"
  "io/ioutil"
  "strconv"
  "encoding/json"
//   "log"
)

type Person struct {
	// Kind     string `json:"kind""`
    Items       []Items    `json:"items"`
    
    // Birthday string `json:"birthday"`
}
type Items struct {
	Kind     string `json:"kind"`
    Id    string    `json:"id"`
	Etag   string    `json:"etag"`
    
    // Birthday string `json:"birthday"`
}



type Items struct {
	Kind     string `json:"Authorization"`
    Id    string    `json:"Content-Type"`
	Etag   string    `json:"etag"`
    
}









func main() {
//   url := "http://google.co.jp"
  var (
		num1 int=9784043636037
		s string
		NOTION_ACCESS_TOKEN int =
  )

  	s = strconv.Itoa(num1)
  	url := "https://www.googleapis.com/books/v1/volumes?q=isbn:"+s


  	resp, _ := http.Get(url)
  	defer resp.Body.Close()

  	byteArray, _ := ioutil.ReadAll(resp.Body)
	jsonBytes := ([]byte)(byteArray)
	data := new(Person)
	if err := json.Unmarshal(jsonBytes, data); err != nil {
        fmt.Println("JSON Unmarshal error:", err)
        return
    }
 
    fmt.Println(data.Items[0].Kind)

	// notion -----------------------------------------------------
	url := "https://api.notion.com/v1/pages"



}



	