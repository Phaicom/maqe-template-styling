package models

type Forum struct {
	Author Author `json:"author"`
	Post   Post   `json:"post"`
}
