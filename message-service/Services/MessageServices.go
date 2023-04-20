package services

import (
	"messages/domain"
	"messages/repositories"
)

type MessageService interface {
	GetAllMessages() ([]*domain.Message, error)
	FindMessageById(id string) (*domain.Message, error)
	FindMessageByLecturerEmail(lecturerEmail string) ([]*domain.Message, error)
}

type DefaultMessageService struct {
	repo repositories.MessageRepository
}

func (s DefaultMessageService) GetAllMessages() ([]*domain.Message, error) {

	return s.repo.FindAll()
}
func (s DefaultMessageService) FindMessageById(id string) (*domain.Message, error) {

	return s.repo.FindById(id)
}

func (s DefaultMessageService) FindMessageByLecturerEmail(lecturerEmail string) ([]*domain.Message, error) {

	return s.repo.FindMessageByLecturerEmail(lecturerEmail)
}

func NewMessageService(repo repositories.MessageRepository) DefaultMessageService {
	return DefaultMessageService{repo}
}
