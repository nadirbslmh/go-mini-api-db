package models

type PostRequest struct {
	Title   string `json:"title"`
	Content string `json:"content"`
}
