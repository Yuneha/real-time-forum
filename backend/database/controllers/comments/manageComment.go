package comments

import (
	"my-real-time-forum/backend/database/initialize"
	"time"
)

func AddComment(userId, postId int, message string) error {
	db, err := initialize.OpenDB()
	if err != nil {
		return err
	}
	defer db.Close()

	_, err = db.Exec("INSERT INTO `comment`(`post_id`, `user_id`, `publication_date`, `message`) VALUES(?,?,?,?)", postId, userId, time.Now().Format("02/01/2006 15:04:05"), message)
	return err
}
