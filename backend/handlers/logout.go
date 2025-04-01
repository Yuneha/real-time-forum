package handlers

import (
	"encoding/json"
	"my-real-time-forum/backend/database/controllers/users"
	"my-real-time-forum/backend/database/structure"
	"net/http"
)

func LogoutHandler(w http.ResponseWriter, r *http.Request) {
	//reveice information and create a user
	var user structure.User
	if err := json.NewDecoder(r.Body).Decode(&user); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	// Retrieve session ID from Cookie and delete the session on the server side
	cookie, err := r.Cookie(user.Username)
	if err == nil {
		sessionMutex.Lock()
		delete(sessions, cookie.Value)
		sessionMutex.Unlock()
	}

	http.SetCookie(w, &http.Cookie{
		Name:   user.Username,
		Value:  "",
		Path:   "/",
		MaxAge: -1, // Delete the Cookie
	})

	users.SetConnected(user.UserId, 0)
}
