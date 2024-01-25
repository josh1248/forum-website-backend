package models

import (
	"log"

	"github.com/josh1248/forum-website-backend/internal/auth"
	"github.com/josh1248/forum-website-backend/internal/entities"
)

// Development-only function.
// Clears junk data that may have been inputted for testing.
// Helps to avoid the need for repeated db meddling in sqlite.
// need to break abstraction here for testing purposes.
func resetDB() {
	resetUsers()
	resetPosts()
	//resetComments()
}

var testUsers []entities.User = []entities.User{
	{ID: 1, Name: "Jojo", Reputation: 200, Password: `letmein`},
	{ID: 2, Name: "PasswordGuy", Reputation: -30, Password: `1e!"E#@5yu6V52~\n42`},
	{ID: 3, Name: "Momo", Reputation: 1500, Password: `qwerty123`},
	{ID: 4, Name: "Geegee", Reputation: 0, Password: `letmein`},
}

func resetUsers() {
	_, err := db.Exec("DROP TABLE users")
	if err != nil {
		log.Fatal(err)
	}

	_, err = db.Exec(`
		CREATE TABLE users (
			id 			INTEGER PRIMARY KEY,
			name 		TEXT 	UNIQUE NOT NULL,
			reputation 	INT		NOT NULL,
			password 	TEXT 	NOT NULL
		)
	`)
	if err != nil {
		log.Fatal(err)
	}

	for _, testUser := range testUsers {
		processedPW, err := auth.ProcessPassword(testUser.Password)
		if err != nil {
			log.Fatal(err)
		}

		testUser.Password = processedPW

		_, err = db.NamedExec(`
			INSERT INTO users (id, name, reputation, password) 
			VALUES (:id, :name, :reputation, :password)`,
			testUser)
		if err != nil {
			log.Fatal(err)
		}
	}

	log.Println("Users table in database restarted.")
}

var testPosts []entities.Post = []entities.Post{
	{PostID: 1, UserID: 1, Reputation: 340, Title: `Golden boy`, Content: `Hello world!`},
	{PostID: 2, UserID: 1, Reputation: -10000, Title: `Food tastes`, Content: `I love pineapple pizza.`},
	{PostID: 3, UserID: 3, Reputation: 20, Title: `Passwords`, Content: `This website gotta have 2FA, man.`},
	{PostID: 4, UserID: 4, Reputation: 5, Title: `Web programming is hard.`, Content: `It really is.`},
}

func resetPosts() {
	_, err := db.Exec("DROP TABLE posts")
	if err != nil {
		log.Fatal(err)
	}

	_, err = db.Exec(`
		CREATE TABLE posts (
			postid 		INTEGER PRIMARY KEY,
			userid 		TEXT NOT NULL,
			reputation 	INT NOT NULL,
			title 		TEXT NOT NULL,
			content 	TEXT NOT NULL
		)
	`)
	if err != nil {
		log.Fatal(err)
	}

	for _, testPost := range testPosts {

		_, err = db.NamedExec(`
			INSERT INTO posts (postid, userid, reputation, title, content) 
			VALUES (:postid, :userid, :reputation, :title, :content)`,
			testPost)
		if err != nil {
			log.Fatal(err)
		}
	}

	log.Println("Posts table in database restarted.")
}

/*
var testComments []entities.Comment = []entities.Comment{}

func resetComments() {
}
*/
