package Services

import (
	"MesseageMicroService/restApi/Domain"
)

type MessageService interface {
	GetAllMessages() ([]Domain.Message, error)
	FindMessageById(id string) (*Domain.Message, *Domain.AppError)
	FindMessageByLecturerEmail(lecturerEmail string) ([]Domain.Message, *Domain.AppError)
}
type DefaultMessageService struct {
	repo Domain.MessageRepository
}

func (s DefaultMessageService) GetAllMessages() ([]Domain.Message, error) {

	return s.repo.FindAll()
}
func (s DefaultMessageService) FindMessageById(id string) (*Domain.Message, *Domain.AppError) {

	return s.repo.FindById(id)
}

func (s DefaultMessageService) FindMessageByLecturerEmail(lecturerEmail string) ([]Domain.Message, *Domain.AppError) {

	return s.repo.FindMessageByLecturerEmail(lecturerEmail)
}

func NewMessageService(repo Domain.MessageRepository) DefaultMessageService {
	return DefaultMessageService{repo}
}
