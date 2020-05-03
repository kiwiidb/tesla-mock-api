package handler

import (
	"fmt"
	"net/http"
	"time"
)

var lockResult = `
{
	"response": {
	"result": true,
	"reason": "unlocked"
	}
	}
`

//LockHander locks the teslaj
func LockHandler(w http.ResponseWriter, r *http.Request) {
	time.Sleep(3)
	fmt.Fprintf(w, lockResult)
	w.Header().Set("Content-Type", "application/json")
}
