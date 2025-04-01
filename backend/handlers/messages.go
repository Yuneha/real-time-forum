package handlers

import (
	"encoding/json"
	"my-real-time-forum/backend/database/controllers/messages"
	"net/http"
	"strconv"
)

func MessagesHandler(w http.ResponseWriter, r *http.Request) {
	sender := r.URL.Query().Get("sender")
	recipient := r.URL.Query().Get("recipient")
	offset, _ := strconv.Atoi(r.URL.Query().Get("offset"))
	limit, _ := strconv.Atoi(r.URL.Query().Get("limit"))

	messages, err := messages.GetMessagesWithOffset(sender, recipient, offset, limit)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(messages)
}
