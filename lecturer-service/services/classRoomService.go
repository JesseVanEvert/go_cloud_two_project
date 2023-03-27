package Services

import (
	"fmt"
	"lecturer/ent"
	"lecturer/models"
	"lecturer/repositories"
)

type ClassRoomService interface {
	CreateClassRoom(classroom models.ClassRoom) (*ent.Class, error)
	DeleteClassRoom(id int) (string, error)
	UpdateClassRoom(classroom models.ClassRoom) (*ent.Class, error)
	GetAllClasses() ([]*ent.Class)
}	

type DefaultClassRoomService struct {
	repo repositories.ClassRoomRepository
}

func (dl DefaultClassRoomService) CreateClassRoom(classroom models.ClassRoom) (*ent.Class, error) {
	fmt.Print("Classroom: ", classroom)
	classroomResponse, err := dl.repo.CreateClassRoom(classroom)
	if err != nil {
		return nil, fmt.Errorf("creating classroom: %w", err)
	}

	return classroomResponse, nil
}

func (dl DefaultClassRoomService) DeleteClassRoom(id int) (string, error) {
	classroomResponse, error := dl.repo.DeleteClassRoom(id)
	if error != nil {
		return "", fmt.Errorf("deleting classroom: %w", error)
	}

	return classroomResponse, nil
}

func (dl DefaultClassRoomService) UpdateClassRoom(classroom models.ClassRoom) (*ent.Class, error) {
	classroomResponse, error := dl.repo.UpdateClassRoom(classroom)
	if error != nil {
		return nil, fmt.Errorf("updating classroom: %w", error)
	}

	return classroomResponse, nil
}

func (dl DefaultClassRoomService) GetAllClasses() ([]*ent.Class) {
	return dl.repo.GetAllClasses()
}

func NewClassRoomService(repo repositories.ClassRoomRepository) ClassRoomService {
	return DefaultClassRoomService{repo: repo}
}