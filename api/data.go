package handler

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

func Handler(w http.ResponseWriter, r *http.Request) {
	data, _ = ioutil.ReadFile("./mock_data.json")
	fmt.Fprintf(w, string(data))
}
