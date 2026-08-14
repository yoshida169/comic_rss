package handlers

import (
	"crypto/rand"
	"database/sql"
	"encoding/hex"
	"encoding/json"
	"net/http"
	"sync"

	"golang.org/x/crypto/bcrypt"
)

// DB is the database handle used by auth handlers. It must be set (via SetDB)
// before LoginHandler is used.
var DB *sql.DB

// SetDB wires up the database handle used by auth handlers.
func SetDB(database *sql.DB) {
	DB = database
}

var (
	sessionsMu sync.Mutex
	sessions   = map[string]string{} // token -> user_name
)

type loginRequest struct {
	UserName string `json:"user_name"`
	Password string `json:"password"`
}

type loginResponse struct {
	Token    string `json:"token"`
	UserName string `json:"user_name"`
}

// POST /login
func LoginHandler(w http.ResponseWriter, req *http.Request) {
	var reqBody loginRequest
	if err := json.NewDecoder(req.Body).Decode(&reqBody); err != nil {
		http.Error(w, "fail to decode json\n", http.StatusBadRequest)
		return
	}

	if reqBody.UserName == "" || reqBody.Password == "" {
		http.Error(w, "user_name and password are required\n", http.StatusBadRequest)
		return
	}

	var id int
	var passwordHash string
	err := DB.QueryRow(
		"SELECT id, password_hash FROM users WHERE user_name = ?",
		reqBody.UserName,
	).Scan(&id, &passwordHash)

	if err == sql.ErrNoRows {
		http.Error(w, "invalid user_name or password\n", http.StatusUnauthorized)
		return
	}
	if err != nil {
		http.Error(w, "internal server error\n", http.StatusInternalServerError)
		return
	}

	if err := bcrypt.CompareHashAndPassword([]byte(passwordHash), []byte(reqBody.Password)); err != nil {
		http.Error(w, "invalid user_name or password\n", http.StatusUnauthorized)
		return
	}

	token, err := generateToken()
	if err != nil {
		http.Error(w, "internal server error\n", http.StatusInternalServerError)
		return
	}

	sessionsMu.Lock()
	sessions[token] = reqBody.UserName
	sessionsMu.Unlock()

	json.NewEncoder(w).Encode(loginResponse{
		Token:    token,
		UserName: reqBody.UserName,
	})
}

func generateToken() (string, error) {
	b := make([]byte, 32)
	if _, err := rand.Read(b); err != nil {
		return "", err
	}
	return hex.EncodeToString(b), nil
}
