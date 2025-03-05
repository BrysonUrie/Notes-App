package database

import (
	"crypto/rand"
	"crypto/sha256"
	"encoding/hex"
)

func generateSalt(length int) (string, error) {
	// Create a slice to store the random salt
	salt := make([]byte, length)

	// Fill the slice with random bytes
	_, err := rand.Read(salt)
	if err != nil {
		return "", err
	}

	return hex.EncodeToString(salt), nil
}

func hashPassword(password string, salt string) (string, error) {
	hash := sha256.New()
	hash.Write([]byte(password + salt))
	return hex.EncodeToString(hash.Sum(nil)), nil
}

func CreateUserTable() error {
	db := GetDB()
	defer db.Close()

	_, err := db.Exec("CREATE TABLE IF NOT EXISTS users (id INTEGER PRIMARY KEY, username TEXT, password TEXT, salt TEXT)")
	if err != nil {
		return err
	}

	return nil
}

func AddUserToDB(id string, username string, password string) error {
	salt, err := generateSalt(16)
	if err != nil {
		return err
	}
	hashedPassword, err := hashPassword(password, salt)
	if err != nil {
		return err
	}

	db := GetDB()
	defer db.Close()

	_, err = db.Exec("INSERT INTO users (id, username, password, salt) VALUES (?, ?, ?, ?)", id, username, hashedPassword, salt)
	if err != nil {
		if err.Error() == "UNIQUE constraint failed: users.id" {
			return nil
		}
		return err
	}

	return nil
}

func GetUser(username string, password string) error {

	return nil
}
