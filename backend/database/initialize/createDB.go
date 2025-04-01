package initialize

func CreateDB() error {
	db, err := OpenDB()
	if err != nil {
		return err
	}
	defer db.Close()

	_, err = db.Exec("CREATE TABLE IF NOT EXISTS `user` (`user_id` INTEGER PRIMARY KEY AUTOINCREMENT UNIQUE NOT NULL, `username` TEXT UNIQUE NOT NULL, `email` TEXT UNIQUE NOT NULL, `password` TEXT NOT NULL, `first_name` TEXT NOT NULL, `last_name` TEXT NOT NULL, `gender` TEXT NOT NULL, `date_of_birth` TEXT NOT NULL, `age` INTEGER NOT NULL, `registration_date` TEXT NOT NULL, `profile_picture` TEXT, `is_connected` INTEGER NOT NULL, `role` INTEGER NOT NULL);")
	if err != nil {
		return err
	}

	_, err = db.Exec("CREATE TABLE `comment` (`comment_id` INTEGER PRIMARY KEY AUTOINCREMENT UNIQUE NOT NULL, `post_id` INTEGER NOT NULL REFERENCES `post`(`post_id`), `user_id` INTEGER NOT NULL REFERENCES `user`(`user_id`), `publication_date` TEXT NOT NULL, `message` TEXT NOT NULL);")
	if err != nil {
		return err
	}

	_, err = db.Exec("CREATE TABLE IF NOT EXISTS `post` (`post_id` INTEGER PRIMARY KEY AUTOINCREMENT UNIQUE NOT NULL, `user_id` INTEGER NOT NULL REFERENCES `user`(`user_id`), `title` TEXT NOT NULL, `publication_date` TEXT NOT NULL, `categorie` TEXT NOT NULL, `message` TEXT NOT NULL, `nb_reply` INTEGER NOT NULL, `nb_like` INTEGER NOT NULL, `nb_dislike` INTEGER NOT NULL);")
	if err != nil {
		return err
	}

	_, err = db.Exec("CREATE TABLE `categorie` (`categorie_id` INTEGER PRIMARY KEY AUTOINCREMENT UNIQUE NOT NULL, `name` TEXT NOT NULL);")
	if err != nil {
		return err
	}

	_, err = db.Exec("CREATE TABLE IF NOT EXISTS `message` (`message_id` INTEGER PRIMARY KEY AUTOINCREMENT, `sender` TEXT, `recipient` TEXT, `message` TEXT, `timestamp` TEXT NOT NULL, read_status BOOLEAN DEFAULT FALSE);")
	if err != nil {
		return err
	}

	return err
}
