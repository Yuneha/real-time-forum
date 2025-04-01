package comments

import (
	"my-real-time-forum/backend/database/initialize"
	"my-real-time-forum/backend/database/structure"
)

func GetAllCommentsOfPost(postId int) ([]structure.Comment, error) {
	var allComments []structure.Comment
	db, err := initialize.OpenDB()
	if err != nil {
		return nil, err
	}
	defer db.Close()

	datas, err := db.Query("SELECT * FROM `comment` WHERE `post_id`=?", postId)
	if err != nil {
		return nil, err
	}
	defer datas.Close()

	for datas.Next() {
		var comment structure.Comment
		err = datas.Scan(&comment.CommentId, &comment.PostId, &comment.UserId, &comment.PublicationDate, &comment.Message)
		if err != nil {
			return nil, err
		}
		allComments = append(allComments, comment)
	}

	return allComments, nil
}

func GetAllCommentsOfPostWithUser(postId int) ([]structure.Comment, []structure.User, error) {
	var allComments []structure.Comment
	var users []structure.User
	db, err := initialize.OpenDB()
	if err != nil {
		return nil, nil, err
	}
	defer db.Close()

	datas, err := db.Query("SELECT * FROM `comment` JOIN user ON comment.user_id = user.user_id WHERE `post_id`=? ", postId)
	if err != nil {
		return nil, nil, err
	}
	defer datas.Close()

	for datas.Next() {
		var comment structure.Comment
		var user structure.User
		err = datas.Scan(&comment.CommentId, &comment.PostId, &comment.UserId, &comment.PublicationDate, &comment.Message, &user.UserId, &user.Username, &user.Email, &user.Password, &user.FirstName, &user.LastName, &user.Gender, &user.DateOfBirth, &user.Age, &user.RegistrationDate, &user.ProfilePicture, &user.IsConnected, &user.Role)
		if err != nil {
			return nil, nil, err
		}
		allComments = append(allComments, comment)
		users = append(users, user)
	}

	return allComments, users, nil
}
