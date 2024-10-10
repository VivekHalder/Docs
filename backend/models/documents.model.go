package models

type Document struct {
	ID      string `json:"id" db:"id"`
	Title   string `json:"title" db:"title"`
	Content string `json:"content" db:"content"`
	Version int    `json:"version" db:"version"`
}
