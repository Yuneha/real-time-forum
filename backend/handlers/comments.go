package handlers

import (
	"encoding/json"
	"fmt"
	"my-real-time-forum/backend/database/controllers/comments"
	"my-real-time-forum/backend/database/structure"
	"net/http"
)

type response struct {
	Comments []structure.Comment `json:"comments"`
	Users    []structure.User    `json:"users"`
}

func CommentsHandler(w http.ResponseWriter, r *http.Request) {
	var post structure.Post
	if err := json.NewDecoder(r.Body).Decode(&post); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	var allComments, users, err = comments.GetAllCommentsOfPostWithUser(post.PostId)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	response := response{
		Comments: allComments,
		Users:    users,
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}
