package handler

import (
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/sirupsen/logrus"
)

func Handler(w http.ResponseWriter, r *http.Request) {
	data, err := ioutil.ReadFile("./mock_data.json")
	if err != nil {
		logrus.WithError(err).Error("error reading file")
	}
	fmt.Fprintf(w, string(data))
}
