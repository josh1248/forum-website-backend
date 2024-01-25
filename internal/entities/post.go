package entities

type Post struct {
	PostID     int    `db:"postid" json:"postid"`
	UserID     int    `db:"userid" json:"userid"`
	Reputation int    `db:"reputation" json:"reputation"`
	Title      string `db:"title" json:"title"`
	Content    string `db:"content" json:"content"`
}
