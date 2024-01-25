package entities

type Comment struct {
	CommentID  int    `db:"commentid" json:"commentid"`
	UserID     int    `db:"userid" json:"userid"`
	PostID     string `db:"postid" json:"postid"`
	Reputation string `db:"reputation" json:"reputation"`
	Body       string `db:"body" json:"body"`
}
