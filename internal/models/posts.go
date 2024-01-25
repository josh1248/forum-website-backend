package models

import (
	"github.com/josh1248/forum-website-backend/internal/entities"
)

// Use StructScan here if memory becomes an issue.
func FindAllPosts() ([]entities.Post, error) {
	var posts []entities.Post
	err := db.Select(&posts, "SELECT * FROM posts")
	return posts, err
}

// Create a new user with a given name.
// username must be unique to prevent confusion in forums.
func CreatePost(formInput entities.Post) error {
	//Using the behaviour of id as an INTEGER PRIMARY KEY row,
	//SQLite can generate incremental IDs for us if we leave it empty.
	//Documentation: https://www.sqlite.org/autoinc.html
	_, err := db.NamedExec(`
		INSERT INTO posts (id, madeby, reputation, title, content) 
		VALUES (:id, :madeby, :reputation, :title, :content)`,
		formInput)
	return err
}
