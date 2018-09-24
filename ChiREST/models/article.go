package models
type Article struct {
	ID     string `json:"id"`
	UserID int64  `json:"user_id"`
	Title  string `json:"title"`
	Slug   string `json:"slug"`
}