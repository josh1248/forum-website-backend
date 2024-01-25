package models

import (
	"github.com/josh1248/forum-website-backend/internal/auth"
	"github.com/josh1248/forum-website-backend/internal/entities"
)

// Use StructScan here if memory becomes an issue.
func FindAllUsers() ([]entities.User, error) {
	var users []entities.User
	err := db.Select(&users, "SELECT * FROM users")
	return users, err
}

// username is unique, so return type is a single item
func FindUserByName(username string) (entities.User, error) {
	var user entities.User

	err := db.Get(&user, "SELECT * FROM users WHERE name = ?", username)
	//sql.ErrNoRows to seperate different errors.
	return user, err
}

// Create a new user with a given name.
// username must be unique to prevent confusion in forums.
func CreateUser(formInput entities.InputUser) error {
	//Using the behaviour of id as an INTEGER PRIMARY KEY row,
	//SQLite can generate incremental IDs for us if we leave it empty.
	//Documentation: https://www.sqlite.org/autoinc.html
	_, err := db.NamedExec(`INSERT INTO users (name, reputation, password) 
					        VALUES (:name, 0, :password)`,
		formInput)
	return err

	// createStmt.LastInsertId() possible here?
}

// given a username and plaintext password, return if login details are correct (plus error).
func AuthenticateUser(formInput entities.InputUser) (bool, error) {
	user, err := FindUserByName(formInput.Name)
	if err != nil {
		return false, err
	}

	return auth.VerifyPassword(user.Password, formInput.Password), nil
}
