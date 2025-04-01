package handlers

import (
	"encoding/json"
	"my-real-time-forum/backend/database/controllers/users"
	"my-real-time-forum/backend/database/structure"
	"net/http"
	"sync"

	"golang.org/x/crypto/bcrypt"
)

var (
	// Map for storing session information
	sessions = make(map[string]bool)
	// Mutex to synchronize access to session information
	sessionMutex = &sync.Mutex{}
)

type Data struct {
	UserId      int    `json:"user_id"`
	Username    string `json:"username"`
	Email       string `json:"email"`
	IsConnected int    `json:"is_connected"`
}

func LoginHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	var emptyUser structure.User

	//reveice information and create a user
	var user structure.User
	if err := json.NewDecoder(r.Body).Decode(&user); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	var newUser structure.User

	var err error
	if user.Username != "" {
		newUser, err = users.GetUserByUsername(user.Username)
	} else {
		newUser, err = users.GetUserByEmail(user.Email)
	}
	if err != nil {
		http.Error(w, "Internal server error", http.StatusInternalServerError)
		return
	}

	if newUser == emptyUser {
		http.Error(w, "User not found", http.StatusNotFound)
		return
	}

	if bcrypt.CompareHashAndPassword([]byte(newUser.Password), []byte(user.Password)) != nil {
		http.Error(w, "Wrong Password", http.StatusSeeOther)
		return
	}

	// Save the session on the server side
	sessionMutex.Lock()
	sessions[newUser.Username] = true
	sessionMutex.Unlock()

	// Send the session ID to the client as a Cookie
	http.SetCookie(w, &http.Cookie{
		Name:   newUser.Username,
		Value:  newUser.Username,
		Path:   "/",
		MaxAge: 60,
	})

	users.SetConnected(newUser.UserId, 1)

	if newUser.IsConnected == 1 {
		http.Error(w, "Already connected", http.StatusSeeOther)
		return
	}

	logUser := Data{
		UserId:      newUser.UserId,
		Username:    newUser.Username,
		Email:       newUser.Email,
		IsConnected: newUser.IsConnected,
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(logUser)
}
