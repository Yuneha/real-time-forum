package messages

import (
	"my-real-time-forum/backend/database/initialize"
	"my-real-time-forum/backend/database/structure"
	"time"
)

func AddMessage(message structure.Message) error {
	db, err := initialize.OpenDB()
	if err != nil {
		return err
	}
	defer db.Close()

	_, err = db.Exec("INSERT INTO `message` (sender, recipient, message, timestamp, read_status) VALUES (?, ?, ?, ?, ?)", message.Sender, message.Recipient, message.Message, time.Now(), false)
	return err
}

func MarkMessageAsRead(username, sender string) error {
	db, err := initialize.OpenDB()
	if err != nil {
		return err
	}
	defer db.Close()

	_, err = db.Exec("UPDATE 'message' SET read_status = ? WHERE recipient = ? AND sender = ? AND read_status = ?", true, username, sender, false)
	return err
}
