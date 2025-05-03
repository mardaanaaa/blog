// internal/models/comment.go
package models

// Comment модель для комментариев к постам
type Comment struct {
	ID        int    `json:"id" gorm:"primary_key"`
	PostID    int    `json:"post_id"`
	Author    string `json:"author"`
	Content   string `json:"content"`
	CreatedAt string `json:"created_at"`
}
