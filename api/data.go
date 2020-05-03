package handler

import (
  "fmt"
  "net/http"
  "io/ioutil"
)


func Handler(w http.ResponseWriter, r *http.Request) {
  data, _ = ioutil.ReadFile("./mock_data.json")
  fmt.Fprintf(w, string(data)
}

