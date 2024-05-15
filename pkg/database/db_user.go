package database

func CreateUserTable() error {
	db := GetDB()
	defer db.Close()

	_, err := db.Exec("CREATE TABLE IF NOT EXISTS users (id INTEGER PRIMARY KEY, username TEXT, password TEXT)")
	if err != nil {
		return err
	}

	return nil
}

func AddUserToDB(id string, username string, password string) error {
	db := GetDB()
	defer db.Close()

	_, err := db.Exec("INSERT INTO users (id, username, password) VALUES (?, ?, ?)", id, username, password)
	if err != nil {
		if err.Error() == "UNIQUE constraint failed: users.id" {
			return nil
		}
		return err
	}

	return nil
}
