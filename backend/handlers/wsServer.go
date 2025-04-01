package handlers

import (
	"log"
	"my-real-time-forum/backend/database/controllers/messages"
	"my-real-time-forum/backend/database/controllers/users"
	"my-real-time-forum/backend/database/structure"
	"net/http"
	"time"

	"github.com/gorilla/websocket"
)

var (
	upgrader = websocket.Upgrader{
		ReadBufferSize:  1024,
		WriteBufferSize: 1024,
	}
)

var usersArr = make(map[string]*websocket.Conn)
var broadcastMessage = make(chan structure.Message)
var broadcastUsers = make(chan structure.UserList)
var broadcastTypingStatus = make(chan Typing)

type User struct {
	Conn     *websocket.Conn
	Username string
}

type Typing struct {
	Sender    string
	Recipient string
	Type      string
	Status    string
}

var userConnected = make(map[*User]bool) // Track connected clients
var test string

func ConnectionsHandler(w http.ResponseWriter, r *http.Request) {
	// Begin by upgrading the HTTP request
	ws, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Println(err)
		return
	}
	defer ws.Close()

	//on open get the logged user username
	var username string
	err = ws.ReadJSON(&username)
	if err != nil {
		log.Printf("error: %v", err)
		return
	}

	usersArr[username] = ws

	user := &User{Username: username, Conn: ws}
	userConnected[user] = true

	broadcastUserList(w, username)

	for {
		var request map[string]interface{}
		err := ws.ReadJSON(&request)
		if err != nil {
			log.Printf("error: %v", err)
			delete(usersArr, username)
			break
		}

		if request["type"] == "message" {
			msg := structure.Message{
				Sender:     request["sender"].(string),
				Message:    request["message"].(string),
				Recipient:  request["recipient"].(string),
				Timestamp:  time.Now().Format(time.RFC3339),
				ReadStatus: false,
			}
			messages.AddMessage(msg)
			broadcastMessage <- msg
		} else if request["type"] == "mark_read" {
			msg := structure.Message{
				Sender: request["sender"].(string),
			}
			messages.MarkMessageAsRead(user.Username, msg.Sender)
		}

		if request["type"] == "updateUserList" {
			for user := range userConnected {
				response := structure.UserList{
					Type:      "UserList",
					UserList:  getUserList(w, user.Username),
					Sender:    request["sender"].(string),
					Recipient: request["recipient"].(string),
				}
				if user.Username == response.Sender || user.Username == response.Recipient {
					user.Conn.WriteJSON(response)
				}
			}
		}

		if request["type"] == "notification" {
			response := structure.Notifications{
				Type:      "notif",
				Recipient: request["recipient"].(string),
			}
			for user := range userConnected {
				if user.Username == response.Recipient {
					count, notification := getUnreadMess(response.Recipient)
					response.Count = count
					response.Notification = notification
					user.Conn.WriteJSON(response)
				}
			}
		}

		if request["type"] == "typing" {
			response := Typing{
				Sender:    request["sender"].(string),
				Recipient: request["recipient"].(string),
				Type:      "typing",
				Status:    request["status"].(string),
			}
			broadcastTypingStatus <- response
		}
	}

	// Handle user disconnection
	defer func() {
		// Set user to disconnected in db
		test = user.Username
		users.SetConnectedByUsername(test, 0)
		delete(userConnected, user)
		broadcastUserListTest(w)
	}()
}

func HandleMessages() {
	for {
		select {
		// Connected / Disconnected handling
		case msg := <-broadcastUsers:
			for user := range userConnected {
				if msg.Sender == user.Username {
					err := user.Conn.WriteJSON(msg)
					if err != nil {
						log.Printf("error: %v", err)
						user.Conn.Close()
						delete(userConnected, user)
					}
				}
			}
		// Messages handling
		case msg := <-broadcastMessage:
			senderConn, senderExists := usersArr[msg.Sender]
			recipientConn, recipientExists := usersArr[msg.Recipient]

			if senderExists {
				err := senderConn.WriteJSON(msg)
				if err != nil {
					log.Printf("error: %v", err)
					senderConn.Close()
					delete(usersArr, msg.Sender)
				}
			}

			if recipientExists && recipientConn != senderConn {
				err := recipientConn.WriteJSON(msg)
				if err != nil {
					log.Printf("error: %v", err)
					recipientConn.Close()
					delete(usersArr, msg.Recipient)
				}
			}

		case msg := <-broadcastTypingStatus:
			senderConn := usersArr[msg.Sender]
			recipientConn, recipientExists := usersArr[msg.Recipient]
			if recipientExists && recipientConn != senderConn {
				err := recipientConn.WriteJSON(msg)
				if err != nil {
					log.Printf("error: %v", err)
					recipientConn.Close()
					delete(usersArr, msg.Recipient)
				}
			}
		}
	}
}

func broadcastUserList(w http.ResponseWriter, username string) {
	for user := range userConnected {
		if user.Username != username {
			response := structure.UserList{
				Type:     "UserList",
				UserList: getUserList(w, user.Username),
				Sender:   user.Username,
			}
			broadcastUsers <- response
		}
	}
}

func broadcastUserListTest(w http.ResponseWriter) {
	for user := range userConnected {
		if user.Username != test {
			response := structure.UserList{
				Type:     "UserList",
				UserList: getUserList(w, user.Username),
				Sender:   user.Username,
			}
			broadcastUsers <- response
		}
	}
}

func getUserList(w http.ResponseWriter, username string) []structure.User {
	var users, err = users.GetAllUsersByAsc()
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return nil
	}

	var userMessages, err1 = messages.GetMessages(username)
	if err1 != nil {
		http.Error(w, err1.Error(), http.StatusBadRequest)
		return nil
	}

	for _, msg := range userMessages {
		for index, user := range users {
			if user.Username == msg.Sender && msg.Recipient == username || msg.Sender == username && msg.Recipient == user.Username {
				var element = users[index]
				users = append(users[:index], users[index+1:]...)
				users = append([]structure.User{element}, users...)
			}
		}
	}
	return users
}

func getUnreadMess(username string) (int, map[string]int) {
	var allUnreadMessages, _ = messages.GetAllUnreadMessage(username)
	var count int
	notification := make(map[string]int)
	for _, message := range allUnreadMessages {
		count++
		notification[message.Sender] = notification[message.Sender] + 1
	}
	return count, notification
}
