package handlers

import (
	"encoding/json"
	"fmt"
	"media_management_go/backend/common"
	"media_management_go/backend/database"
	"net/http"
	"strings"
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
	// Validate token and return 200 if valid
	if claims, ok := requireAuth(w, r); ok {
		writeJSON(w, struct {
			Subject   string    `json:"subject"`
			IssuedAt  time.Time `json:"issued_at"`
			ExpiresAt time.Time `json:"expires_at"`
		}{
			Subject:   claims.Subject,
			IssuedAt:  claims.IssuedAt.Time,
			ExpiresAt: claims.ExpiresAt.Time,
		}, http.StatusOK)
	}
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
		writeJSONError(w, "Failed to generate token", http.StatusInternalServerError)
		return
	}

	// persist the token in sqlite for later introspection / revocation
	if _, err := database.AddToken(signed); err != nil {
		writeJSONError(w, "Failed to persist token", http.StatusInternalServerError)
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


// validateToken extracts and validates the JWT from Authorization header.
// Returns the parsed claims if token is valid, or error if validation fails.
func validateToken(r *http.Request) (*jwt.RegisteredClaims, error) {
	authHeader := r.Header.Get("Authorization")
	if authHeader == "" {
		return nil, fmt.Errorf("missing Authorization header")
	}

	const prefix = "Bearer "
	if !strings.HasPrefix(authHeader, prefix) {
		return nil, fmt.Errorf("invalid Authorization header format")
	}

	tokenStr := strings.TrimPrefix(authHeader, prefix)
	if tokenStr == "" {
		return nil, fmt.Errorf("empty token")
	}

	// Parse and validate JWT signature
	claims := &jwt.RegisteredClaims{}
	token, err := jwt.ParseWithClaims(tokenStr, claims, func(t *jwt.Token) (interface{}, error) {
		if t.Method.Alg() != jwt.SigningMethodHS256.Alg() {
			return nil, fmt.Errorf("unexpected signing method: %v", t.Method.Alg())
		}
		return []byte(common.GetConfig().JWT_KEY), nil
	})

	if err != nil || !token.Valid {
		return nil, fmt.Errorf("invalid token: %v", err)
	}

	// Verify token exists in database
	storedToken, err := database.GetToken(tokenStr)
	if err != nil {
		return nil, fmt.Errorf("database error: %v", err)
	}
	if storedToken == nil {
		return nil, fmt.Errorf("token not found in database")
	}

	return claims, nil
}

// requireAuth is a helper that validates the token and writes error response if invalid
func requireAuth(w http.ResponseWriter, r *http.Request) (*jwt.RegisteredClaims, bool) {
	claims, err := validateToken(r)
	if err != nil {
		writeJSONError(w, err.Error(), http.StatusUnauthorized)
		return nil, false
	}
	return claims, true
}
