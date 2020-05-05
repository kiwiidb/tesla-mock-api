package handler

import (
	"encoding/json"
	"fmt"
	"net/http"
)

var oauthResult = `
{
	"access_token": "test_token",
	"created_at": 1451528865,
	"expires_in": 7776000,
	"token_type": "bearer"
	}
`
var clientId = "test_id"
var clientSecret = "test_secret"
var password = "battmobiel"
var email = "kdebacker@sofico.be"
var grantType = "password"

var accessToken = "test_token"

var unAuthorizedResp = `
{
    "response": "authorization_required_for_txid_e8035ba33093100265d75f31e28e8584"
}`

type AutoGenerated struct {
	GrantType    string `json:"grant_type"`
	ClientID     string `json:"client_id"`
	ClientSecret string `json:"client_secret"`
	Email        string `json:"email"`
	Password     string `json:"password"`
}

//OAuthHandler locks the teslaj
func OAuthHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Need POST request", http.StatusBadRequest)
		return
	}
	var credentials AutoGenerated
	err := json.NewDecoder(r.Body).Decode(&credentials)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	if credentials.Password != password || credentials.ClientID != clientId ||
		credentials.ClientSecret != clientSecret || credentials.Email != email ||
		credentials.GrantType != grantType {
		http.Error(w, unAuthorizedResp, http.StatusBadRequest)
		return
	}
	fmt.Fprintf(w, oauthResult)
	w.Header().Set("Content-Type", "application/json")
}
