package handlers

import (
	"encoding/json"
	"my-real-time-forum/backend/database/structure"
	"net/http"
)

func MainPageHandler(w http.ResponseWriter, r *http.Request) {
	//reveice information and create a user
	var user structure.User
	if err := json.NewDecoder(r.Body).Decode(&user); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	// Retrieve session ID from Cookie
	cookie, err := r.Cookie(user.Username)
	if err != nil {
		http.Error(w, "Unauthorized", http.StatusUnauthorized)
		return
	}

	sessionMutex.Lock()
	authenticated, ok := sessions[cookie.Value]
	sessionMutex.Unlock()

	// Check if the session is valid
	if !ok || !authenticated {
		http.Error(w, "Forbidden", http.StatusForbidden)
		return
	}

}
