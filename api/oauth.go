package handler

import (
	"fmt"
	"net/http"
	"time"
)

var oauthResult = `
{
	"access_token": "8s2wfclhyp5iiikowm3ocnfalt7qfl7es8xhuda3ttusslssx6c14hq7yocp62c5",
	"created_at": 1451528865,
	"expires_in": 7776000,
	"token_type": "bearer"
	}
`

//OAuthHandler locks the teslaj
func OAuthHandler(w http.ResponseWriter, r *http.Request) {
	time.Sleep(3 * time.Second)
	fmt.Fprintf(w, oauthResult)
	w.Header().Set("Content-Type", "application/json")
}
