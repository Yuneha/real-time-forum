package posts

import (
	"my-real-time-forum/backend/database/initialize"
	"time"
)

func AddPost(userId int, title, message, categorie string) error {
	db, err := initialize.OpenDB()
	if err != nil {
		return err
	}
	defer db.Close()

	_, err = db.Exec("INSERT INTO `post`(`user_id`, `title`, `publication_date`, `categorie`, `message`, `nb_reply`, `nb_like`, `nb_dislike`) VALUES(?,?,?,?,?,?,?,?)", userId, title, time.Now().Format("02/01/2006 15:04:05"), categorie, message, 0, 0, 0)
	return err
}

func RemovePost() {

}

func UpdatePost() {

}
