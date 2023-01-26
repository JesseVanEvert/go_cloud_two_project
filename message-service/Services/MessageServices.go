package Services

import (
	"MesseageMicroService/restApi/Domain"
)

type MessageService interface {
	GetAllMessages() ([]Domain.Message, error)
	FindMessageById(id string) (*Domain.Message, *Domain.AppError)
	FindMessageByLecturerId(lecturerId string) (*Domain.Message, *Domain.AppError)
}

// "constructor" like function
// whereby we pass in the repo (interface) as a dependency
type DefaultMessageService struct {
	repo Domain.MessageRepository
}

// receiver function -attaches it as a method to a class
func (s DefaultMessageService) GetAllMessages() ([]Domain.Message, error) {

	//Once again we talk to the interface
	return s.repo.FindAll()
}
func (s DefaultMessageService) FindMessageById(id string) (*Domain.Message, *Domain.AppError) {

	//Once again we talk to the interface
	return s.repo.FindById(id)
}

func (s DefaultMessageService) FindMessageByLecturerId(lecturerId string) (*Domain.Message, *Domain.AppError) {

	//Once again we talk to the interface
	return s.repo.FindMessageByLecturerId(lecturerId)
}

// Helper function to instantiate customer service
func NewMessageService(repo Domain.MessageRepository) DefaultMessageService {
	return DefaultMessageService{repo}
}
