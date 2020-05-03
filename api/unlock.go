package handler

import (
	"fmt"
	"net/http"
	"time"
)

//UnlockHander njlocks the teslaj
func UnLockHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Need POST request", http.StatusBadRequest)
		return
	}
	if !checkAuth(r) {
		http.Error(w, "Unauthorized", http.StatusUnauthorized)
		return
	}
	time.Sleep(3 * time.Second)
	fmt.Fprintf(w, lockResult)
	w.Header().Set("Content-Type", "application/json")
}
