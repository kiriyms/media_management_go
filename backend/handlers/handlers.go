package handlers

import (
	"encoding/json"
	"media_management_go/backend/common"
	"net/http"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

type PostLoginRequest struct {
	Key string `json:"key"`
}

type PostLoginResponse struct {
	Token string `json:"token"`
}

func HandleGetLogin(w http.ResponseWriter, r *http.Request) {
	// validate session token or cookie
}

func HandlePostLogin(w http.ResponseWriter, r *http.Request) {
	// process login key, create session or cookie
	cfg := common.GetConfig()
	var req PostLoginRequest

	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		writeJSONError(w, "Invalid request payload", http.StatusBadRequest)
		return
	}

	if req.Key != cfg.USER_KEY {
		writeJSONError(w, "Unauthorized", http.StatusUnauthorized)
		return
	}

	agent := r.UserAgent()
	exp := time.Now().Add(168 * time.Hour)
	claims := jwt.RegisteredClaims{
		ExpiresAt: jwt.NewNumericDate(exp),
		IssuedAt:  jwt.NewNumericDate(time.Now()),
		Subject:   agent,
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	signed, err := token.SignedString([]byte(cfg.JWT_KEY))
	if err != nil {
		http.Error(w, "Failed to generate token", http.StatusInternalServerError)
		return
	}

	resp := PostLoginResponse{
		Token: signed,
	}
	writeJSON(w, resp, http.StatusOK)

}

func HandleGetLink(w http.ResponseWriter, r *http.Request) {
	// return all links data for user
}

func HandlePostLink(w http.ResponseWriter, r *http.Request) {
	// process new link data for user
}

func HandleGetNote(w http.ResponseWriter, r *http.Request) {
	// return all notes data for user
}

func HandlePostNote(w http.ResponseWriter, r *http.Request) {
	// process new note data for user
}

func writeJSON(w http.ResponseWriter, v interface{}, status int) {
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	w.WriteHeader(status)
	_ = json.NewEncoder(w).Encode(v)
}

func writeJSONError(w http.ResponseWriter, msg string, status int) {
	type errResp struct {
		Error string `json:"error"`
	}
	writeJSON(w, errResp{Error: msg}, status)
}
