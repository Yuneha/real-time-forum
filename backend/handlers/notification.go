package handlers

import (
	"encoding/json"
	"my-real-time-forum/backend/database/controllers/messages"
	"my-real-time-forum/backend/database/structure"
	"net/http"
)

func NotificationHandler(w http.ResponseWriter, r *http.Request) {
	//reveice information and create a user
	var user structure.User
	if err := json.NewDecoder(r.Body).Decode(&user); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	var allUnreadMessages, _ = messages.GetAllUnreadMessage(user.Username)
	var count int
	notification := make(map[string]int)
	for _, message := range allUnreadMessages {
		count++
		notification[message.Sender] = notification[message.Sender] + 1
	}
	response := structure.Notifications{
		Notification: notification,
		Count:        count,
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}
