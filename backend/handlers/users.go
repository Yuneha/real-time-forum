package handlers

import (
	"encoding/json"
	"my-real-time-forum/backend/database/controllers/messages"
	"my-real-time-forum/backend/database/controllers/users"
	"my-real-time-forum/backend/database/structure"
	"net/http"
)

func UsersHandler(w http.ResponseWriter, r *http.Request) {
	//reveice information and create a user
	var logUser structure.User
	if err := json.NewDecoder(r.Body).Decode(&logUser); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	var users, err = users.GetAllUsersByAsc()
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	var userMessages, err1 = messages.GetMessages(logUser.Username)
	if err1 != nil {
		http.Error(w, err1.Error(), http.StatusBadRequest)
		return
	}

	for _, msg := range userMessages {
		for index, user := range users {
			if user.Username == msg.Sender && msg.Recipient == logUser.Username || msg.Sender == logUser.Username && msg.Recipient == user.Username {
				var element = users[index]
				users = append(users[:index], users[index+1:]...)
				users = append([]structure.User{element}, users...)
			}
		}
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(users)
}
