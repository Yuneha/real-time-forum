package posts

import (
	"my-real-time-forum/backend/database/initialize"
	"my-real-time-forum/backend/database/structure"
)

func GetAllPosts() ([]structure.Post, error) {
	var allPosts []structure.Post
	db, err := initialize.OpenDB()
	if err != nil {
		return nil, err
	}
	defer db.Close()

	datas, err := db.Query("SELECT * FROM `post`")
	if err != nil {
		return nil, err
	}
	defer datas.Close()

	for datas.Next() {
		var post structure.Post
		err = datas.Scan(&post.PostId, &post.UserId, &post.Title, &post.PublicationDate, &post.Categorie, &post.Message, &post.NbReply, &post.NbLike, &post.NbDislike)
		if err != nil {
			return nil, err
		}
		allPosts = append(allPosts, post)
	}

	return allPosts, nil
}

func GetAllPostsWithUser() ([]structure.Post, []structure.User, error) {
	var allPosts []structure.Post
	var users []structure.User
	db, err := initialize.OpenDB()
	if err != nil {
		return nil, nil, err
	}
	defer db.Close()

	datas, err := db.Query("SELECT * FROM `post` JOIN user ON post.user_id = user.user_id")
	if err != nil {
		return nil, nil, err
	}
	defer datas.Close()

	for datas.Next() {
		var post structure.Post
		var user structure.User
		err = datas.Scan(&post.PostId, &post.UserId, &post.Title, &post.PublicationDate, &post.Categorie, &post.Message, &post.NbReply, &post.NbLike, &post.NbDislike, &user.UserId, &user.Username, &user.Email, &user.Password, &user.FirstName, &user.LastName, &user.Gender, &user.DateOfBirth, &user.Age, &user.RegistrationDate, &user.ProfilePicture, &user.IsConnected, &user.Role)
		if err != nil {
			return nil, nil, err
		}
		allPosts = append(allPosts, post)
		users = append(users, user)
	}

	return allPosts, users, nil
}

func GetPostById(id int) {

}
