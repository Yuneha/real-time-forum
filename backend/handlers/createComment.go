package handlers

import (
	"encoding/json"
	"my-real-time-forum/backend/database/controllers/comments"
	"my-real-time-forum/backend/database/structure"
	"net/http"
)

type responsesComment struct {
	User          structure.User `json:"user"`
	Post          structure.Post `json:"post"`
	CommentObject structure.Post `json:"comment"`
}

func CreateCommentHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	// Parse the JSON request body
	var response responsesComment
	err := json.NewDecoder(r.Body).Decode(&response)
	if err != nil {
		http.Error(w, "Bad request", http.StatusBadRequest)
		return
	}

	user := response.User
	post := response.Post
	comment := response.CommentObject

	comments.AddComment(user.UserId, post.PostId, comment.Message)

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}
