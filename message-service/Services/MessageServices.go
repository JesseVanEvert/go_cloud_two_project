package services

import (
	"messages/domain"
)

type MessageService interface {
	GetAllMessages() ([]domain.Message, error)
	FindMessageById(id string) (*domain.Message, *domain.AppError)
	FindMessageByLecturerEmail(lecturerEmail string) ([]domain.Message, *domain.AppError)
}

type DefaultMessageService struct {
	repo domain.MessageRepository
}

func (s DefaultMessageService) GetAllMessages() ([]domain.Message, error) {

	return s.repo.FindAll()
}
func (s DefaultMessageService) FindMessageById(id string) (*domain.Message, *domain.AppError) {

	return s.repo.FindById(id)
}

func (s DefaultMessageService) FindMessageByLecturerEmail(lecturerEmail string) ([]domain.Message, *domain.AppError) {

	return s.repo.FindMessageByLecturerEmail(lecturerEmail)
}

func NewMessageService(repo domain.MessageRepository) DefaultMessageService {
	return DefaultMessageService{repo}
}
