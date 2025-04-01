package handlers

import (
	"encoding/json"
	"fmt"
	"my-real-time-forum/backend/database/controllers/posts"
	"my-real-time-forum/backend/database/structure"
	"net/http"
)

type Response struct {
	Posts []structure.Post `json:"posts"`
	Users []structure.User `json:"users"`
}

func PostcardHandler(w http.ResponseWriter, r *http.Request) {
	//reveice information and create a user
	var user structure.User
	if err := json.NewDecoder(r.Body).Decode(&user); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	var allPost, users, err = posts.GetAllPostsWithUser()
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	response := Response{
		Posts: allPost,
		Users: users,
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}
