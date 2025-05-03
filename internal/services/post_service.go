package services

import (
	"rest-project/internal/models"
	"rest-project/internal/repository"
)

type PostService struct {
	repo repository.PostRepository
}

// Конструктор для создания нового PostService
func NewPostService(repo repository.PostRepository) *PostService {
	return &PostService{repo: repo}
}

// Создание нового поста
func (s *PostService) CreatePost(post models.Post) (*models.Post, error) {
	return s.repo.Create(post)
}

// Получение всех постов
func (s *PostService) GetAllPosts() ([]models.Post, error) {
	return s.repo.GetAll()
}

// Получение поста по ID
func (s *PostService) GetPostByID(id int) (*models.Post, error) {
	return s.repo.GetById(id)
}

// Обновление поста
func (s *PostService) UpdatePost(id int, post models.Post) error {
	return s.repo.Update(id, post)
}

// Удаление поста
func (s *PostService) DeletePost(id int) error {
	return s.repo.Delete(id)
}
