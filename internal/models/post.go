package models

type Post struct {
	ID        uint   `json:"id" gorm:"primaryKey"`
	Title     string `json:"title"`
	Content   string `json:"content"`
	CreatedAt string `json:"created_at"`
}

type PostEdit struct {
	Title   string `json:"title"`
	Content string `json:"content"`
}
