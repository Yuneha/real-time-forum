package handlers

import (
	"encoding/json"
	"my-real-time-forum/backend/database/controllers/posts"
	"my-real-time-forum/backend/database/structure"
	"net/http"
)

type responsesPost struct {
	User             structure.User `json:"user"`
	CreatePostObject structure.Post `json:"post"`
}

func CreatePostHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	// Parse the JSON request body
	var response responsesPost
	err := json.NewDecoder(r.Body).Decode(&response)
	if err != nil {
		http.Error(w, "Bad request", http.StatusBadRequest)
		return
	}

	user := response.User
	post := response.CreatePostObject

	posts.AddPost(user.UserId, post.Title, post.Message, post.Categorie)

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(user)
}
