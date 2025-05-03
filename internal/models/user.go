package models

// User модель для пользователей
type User struct {
	ID        uint   `gorm:"primaryKey"`
	Username  string `gorm:"unique"`
	Password  string
	CreatedAt uint
	UpdatedAt uint
}
