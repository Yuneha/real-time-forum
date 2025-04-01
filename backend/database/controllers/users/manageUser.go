package users

import (
	"my-real-time-forum/backend/database/initialize"
	"time"
)

func AddUser(username, email, password, firstName, lastName, gender, dateOfBirth string, age int) error {
	db, err := initialize.OpenDB()
	if err != nil {
		return err
	}
	defer db.Close()

	_, err = db.Exec("INSERT INTO `user`(`username`, `email`, `password`, `first_name`, `last_name`, `gender`, `date_of_birth`, `age`, `registration_date`, `profile_picture`, `is_connected`, `role`) VALUES(?,?,?,?,?,?,?,?,?,?,?,?)", username, email, password, firstName, lastName, gender, dateOfBirth, age, time.Now().Format("02/01/2006"), "null", 0, 1)
	return err
}

func RemoveUser() {

}

func UpdateUSer() {

}

func SetConnected(userId, isConnected int) error {
	db, err := initialize.OpenDB()
	if err != nil {
		return err
	}
	defer db.Close()

	_, err = db.Exec("UPDATE `user` SET is_connected=? WHERE user_Id = ?", isConnected, userId)
	return err
}

func SetConnectedByUsername(username string, isConnected int) error {
	db, err := initialize.OpenDB()
	if err != nil {
		return err
	}
	defer db.Close()

	_, err = db.Exec("UPDATE `user` SET is_connected=? WHERE username = ?", isConnected, username)
	return err
}
