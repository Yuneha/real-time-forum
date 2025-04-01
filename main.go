package main

import (
	"fmt"
	"log"
	"my-real-time-forum/backend/database/controllers/users"
	"my-real-time-forum/backend/database/initialize"
	"my-real-time-forum/backend/handlers"
	"net/http"
	"os"
	"os/signal"
	"syscall"

	_ "github.com/mattn/go-sqlite3"
)

func init() {
	// os.Remove("./backend/database/database.db")
	initialize.CreateDB()

	fs := http.FileServer(http.Dir("frontend/static"))
	http.Handle("/frontend/static/", http.StripPrefix("/frontend/static", fs))
}

func main() {
	fmt.Println("Server running at: http://localhost:5656")
	setupRoutes()
	go handlers.HandleMessages()
	log.Fatal(http.ListenAndServe(":5656", nil))

	// Set up channel to listen for system signals (e.g., Ctrl+C or kill)
	sigs := make(chan os.Signal, 1)
	signal.Notify(sigs, syscall.SIGINT, syscall.SIGTERM)

	go func() {
		<-sigs
		log.Println("Server is shutting down...")
		logoutAllUsers()
		os.Exit(0)
	}()
}

func setupRoutes() {
	http.Handle("/", http.FileServer(http.Dir("./frontend")))
	http.HandleFunc("/ws", handlers.ConnectionsHandler)
	http.HandleFunc("/register", handlers.RegisterHandler)
	http.HandleFunc("/login", handlers.LoginHandler)
	http.HandleFunc("/posts", handlers.PostcardHandler)
	http.HandleFunc("/createPost", handlers.CreatePostHandler)
	http.HandleFunc("/profile", handlers.ProfileHandler)
	http.HandleFunc("/createComment", handlers.CreateCommentHandler)
	http.HandleFunc("/comments", handlers.CommentsHandler)
	http.HandleFunc("/getUsers", handlers.UsersHandler)
	http.HandleFunc("/getMessages", handlers.MessagesHandler)
	http.HandleFunc("/mainPage", handlers.MainPageHandler)
	http.HandleFunc("/notification", handlers.NotificationHandler)
	http.HandleFunc("/logout", handlers.LogoutHandler)
}

func logoutAllUsers() {
	allUsers, _ := users.GetAllUsers()
	for _, user := range allUsers {
		users.SetConnected(user.UserId, 0)
	}
}
