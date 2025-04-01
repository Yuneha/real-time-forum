package handlers

import (
	"encoding/json"
	"net/http"
	"strings"

	"my-real-time-forum/backend/database/controllers/users"
	"my-real-time-forum/backend/database/functions"
	"my-real-time-forum/backend/database/structure"

	"golang.org/x/crypto/bcrypt"
)

type ErrorResponse struct {
	Message string `json:"message"`
}

func RegisterHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	// reveice information and create a user
	var user structure.User
	if err := json.NewDecoder(r.Body).Decode(&user); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	// Check length of password
	if len(user.Password) < 8 {
		json.NewEncoder(w).Encode(ErrorResponse{Message: "your password have less than 8 caractere"})
		return
	}

	// convert the format of the birthdate
	formattedDate := functions.FormattedDate(user.DateOfBirth)
	// get the Age with the birthdate
	age := functions.GetAge(formattedDate)
	// crypt the password
	cryptPassword, _ := bcrypt.GenerateFromPassword([]byte(user.Password), 12)

	Emailallreadyuse, _ := users.GetUserByEmail(user.Email)
	Usernameallreadyuse, _ := users.GetUserByUsername(user.Username)
	if Emailallreadyuse.Email != "" {
		if strings.EqualFold(Emailallreadyuse.Email, user.Email) {
			if Usernameallreadyuse.Username != "" {
				if strings.EqualFold(Usernameallreadyuse.Username, user.Username) {
					json.NewEncoder(w).Encode(ErrorResponse{Message: "Username and Email already use."})
				}
			} else {
				json.NewEncoder(w).Encode(ErrorResponse{Message: "Email already use."})
			}
			return
		}
	}

	if Usernameallreadyuse.Username != "" {
		if strings.EqualFold(Usernameallreadyuse.Username, user.Username) {
			json.NewEncoder(w).Encode(ErrorResponse{Message: "Username already use."})
			return
		}
	}

	// add user to DB
	users.AddUser(user.Username, user.Email, string(cryptPassword), user.FirstName, user.LastName, user.Gender, formattedDate, age)

	logUser, err := users.GetUserByEmail(user.Email)
	if err != nil {
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(logUser)
}
