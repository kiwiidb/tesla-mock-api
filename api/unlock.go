package handler

import (
	"fmt"
	"net/http"
	"time"
)

//UnlockHander njlocks the teslaj
func UnLockHandler(w http.ResponseWriter, r *http.Request) {
	time.Sleep(3 * time.Second)
	fmt.Fprintf(w, lockResult)
	w.Header().Set("Content-Type", "application/json")
}
