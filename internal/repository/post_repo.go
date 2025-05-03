package repository

import (
	"gorm.io/gorm"
	"rest-project/internal/models"
)

type PostRepository interface {
	Create(post models.Post) (*models.Post, error)
	GetAll() ([]models.Post, error)
	GetById(id int) (*models.Post, error)
	Update(id int, post models.Post) error
	Delete(id int) error
}

type PostRepositoryImpl struct {
	db *gorm.DB
}

func NewPostRepository(db *gorm.DB) PostRepository {
	return &PostRepositoryImpl{db: db}
}

func (r *PostRepositoryImpl) Create(post models.Post) (*models.Post, error) {
	if err := r.db.Create(&post).Error; err != nil {
		return nil, err
	}
	return &post, nil
}

func (r *PostRepositoryImpl) GetAll() ([]models.Post, error) {
	var posts []models.Post
	err := r.db.Find(&posts).Error
	return posts, err
}

func (r *PostRepositoryImpl) GetById(id int) (*models.Post, error) {
	var post models.Post
	err := r.db.First(&post, id).Error
	return &post, err
}

func (r *PostRepositoryImpl) Update(id int, post models.Post) error {
	return r.db.Model(&models.Post{}).Where("id = ?", id).Updates(post).Error
}

func (r *PostRepositoryImpl) Delete(id int) error {
	return r.db.Delete(&models.Post{}, id).Error
}
