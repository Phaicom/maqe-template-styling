package models

type Post struct {
	ID        int64  `json:"id"`
	AuthorID  int64  `json:"author_id"`
	Title     string `json:"title"`
	Body      string `json:"body"`
	ImageURL  string `json:"image_url"`
	CreatedAt string `json:"created_at"`
}
