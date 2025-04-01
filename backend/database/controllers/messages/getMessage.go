package messages

import (
	"my-real-time-forum/backend/database/initialize"
	"my-real-time-forum/backend/database/structure"
)

func GetMessages(username string) ([]structure.Message, error) {
	var messages []structure.Message
	db, err := initialize.OpenDB()
	if err != nil {
		return nil, err
	}
	defer db.Close()

	datas, err := db.Query("SELECT sender, recipient, message, timestamp, read_status FROM `message` WHERE sender=? OR recipient=? ORDER BY timestamp ASC", username, username)
	if err != nil {
		return nil, err
	}
	defer datas.Close()

	for datas.Next() {
		var message structure.Message
		err = datas.Scan(&message.Sender, &message.Recipient, &message.Message, &message.Timestamp, &message.ReadStatus)
		if err != nil {
			return nil, err
		}
		messages = append(messages, message)
	}

	return messages, nil
}

func GetMessagesWithOffset(sender, recipient string, offset, limit int) ([]structure.Message, error) {
	var messages []structure.Message
	db, err := initialize.OpenDB()
	if err != nil {
		return nil, err
	}
	defer db.Close()

	datas, err := db.Query("SELECT sender, recipient, message, timestamp, read_status FROM `message` WHERE (sender=? AND recipient=?) OR (sender=? AND recipient=?) ORDER BY timestamp DESC LIMIT ? OFFSET ?", sender, recipient, recipient, sender, limit, offset)
	if err != nil {
		return nil, err
	}
	defer datas.Close()

	for datas.Next() {
		var message structure.Message
		err = datas.Scan(&message.Sender, &message.Recipient, &message.Message, &message.Timestamp, &message.ReadStatus)
		if err != nil {
			return nil, err
		}
		messages = append(messages, message)
	}

	return messages, nil
}

func GetAllUnreadMessage(username string) ([]structure.Message, error) {
	db, err := initialize.OpenDB()
	if err != nil {
		return nil, err
	}
	defer db.Close()

	var AllUnreadMessages []structure.Message

	datas, err := db.Query("SELECT sender, recipient, message, timestamp, read_status FROM `message` WHERE recipient=? AND read_status=?", username, false)
	if err != nil {
		return nil, err
	}
	defer datas.Close()

	for datas.Next() {
		var message structure.Message
		err = datas.Scan(&message.Sender, &message.Recipient, &message.Message, &message.Timestamp, &message.ReadStatus)
		if err != nil {
			return nil, err
		}
		AllUnreadMessages = append(AllUnreadMessages, message)
	}

	return AllUnreadMessages, err
}
