package models

type Reaction struct {
	ID        string
	UserID    int
	PostID    int
	CommentID int
	Vote      int
}
