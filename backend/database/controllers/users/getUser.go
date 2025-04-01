package users

import (
	"my-real-time-forum/backend/database/initialize"
	"my-real-time-forum/backend/database/structure"
)

func GetAllUsers() ([]structure.User, error) {
	db, err := initialize.OpenDB()
	if err != nil {
		return nil, err
	}
	defer db.Close()

	datas, err := db.Query("SELECT * FROM `user`")
	if err != nil {
		return nil, err
	}
	defer datas.Close()

	var users []structure.User
	var user structure.User
	for datas.Next() {
		err = datas.Scan(&user.UserId, &user.Username, &user.Email, &user.Password, &user.FirstName, &user.LastName, &user.Gender, &user.DateOfBirth, &user.Age, &user.RegistrationDate, &user.ProfilePicture, &user.IsConnected, &user.Role)
		if err != nil {
			return nil, err
		}
		users = append(users, user)
	}

	return users, nil
}

func GetAllUsersByAsc() ([]structure.User, error) {
	db, err := initialize.OpenDB()
	if err != nil {
		return nil, err
	}
	defer db.Close()

	datas, err := db.Query("SELECT * FROM `user` ORDER BY LOWER(username) ASC")
	if err != nil {
		return nil, err
	}
	defer datas.Close()

	var users []structure.User
	var user structure.User
	for datas.Next() {
		err = datas.Scan(&user.UserId, &user.Username, &user.Email, &user.Password, &user.FirstName, &user.LastName, &user.Gender, &user.DateOfBirth, &user.Age, &user.RegistrationDate, &user.ProfilePicture, &user.IsConnected, &user.Role)
		if err != nil {
			return nil, err
		}
		users = append(users, user)
	}

	return users, nil
}

func GetUSerById(id int) (structure.User, error) {
	db, err := initialize.OpenDB()
	if err != nil {
		return structure.User{}, err
	}
	defer db.Close()

	datas, err := db.Query("SELECT * FROM `user` WHERE `id`=?", id)
	if err != nil {
		return structure.User{}, err
	}
	defer datas.Close()

	var user structure.User
	for datas.Next() {
		err = datas.Scan(&user.UserId, &user.Username, &user.Email, &user.Password, &user.FirstName, &user.LastName, &user.Gender, &user.DateOfBirth, &user.Age, &user.RegistrationDate, &user.ProfilePicture, &user.IsConnected, &user.Role)
		if err != nil {
			return structure.User{}, err
		}
	}

	return user, nil
}

func GetUserByEmail(email string) (structure.User, error) {
	db, err := initialize.OpenDB()
	if err != nil {
		return structure.User{}, err
	}
	defer db.Close()

	datas, err := db.Query("SELECT * FROM `user` WHERE LOWER(`email`) = LOWER(?)", email)
	if err != nil {
		return structure.User{}, err
	}
	defer datas.Close()

	var user structure.User

	for datas.Next() {
		err = datas.Scan(&user.UserId, &user.Username, &user.Email, &user.Password, &user.FirstName, &user.LastName, &user.Gender, &user.DateOfBirth, &user.Age, &user.RegistrationDate, &user.ProfilePicture, &user.IsConnected, &user.Role)
		if err != nil {
			return structure.User{}, err
		}
	}

	return user, nil
}

func GetUserByUsername(username string) (structure.User, error) {
	db, err := initialize.OpenDB()
	if err != nil {
		return structure.User{}, err
	}
	defer db.Close()

	datas, err := db.Query("SELECT * FROM `user` WHERE LOWER(`username`) = LOWER(?)", username)
	if err != nil {
		return structure.User{}, err
	}
	defer datas.Close()

	var user structure.User

	for datas.Next() {
		err = datas.Scan(&user.UserId, &user.Username, &user.Email, &user.Password, &user.FirstName, &user.LastName, &user.Gender, &user.DateOfBirth, &user.Age, &user.RegistrationDate, &user.ProfilePicture, &user.IsConnected, &user.Role)
		if err != nil {
			return structure.User{}, err
		}
	}

	return user, nil
}

func GetAllConnectedUser() ([]structure.User, error) {
	db, err := initialize.OpenDB()
	if err != nil {
		return nil, err
	}
	defer db.Close()

	datas, err := db.Query("SELECT * FROM `user` WHERE is_connected=?", 1)
	if err != nil {
		return nil, err
	}
	defer datas.Close()

	var users []structure.User
	var user structure.User
	for datas.Next() {
		err = datas.Scan(&user.UserId, &user.Username, &user.Email, &user.Password, &user.FirstName, &user.LastName, &user.Gender, &user.DateOfBirth, &user.Age, &user.RegistrationDate, &user.ProfilePicture, &user.IsConnected, &user.Role)
		if err != nil {
			return nil, err
		}
		users = append(users, user)
	}

	return users, nil
}
