package Services

import (
	"fmt"
	"lecturer/ent"
	"lecturer/repositories"
)

type LecturerService interface {
	GetAllLecturers() ([]*ent.Lecturer, error)
	GetLecturerByID(id int) (*ent.Lecturer, error)
	CreateLecturer(lecturer *ent.Lecturer) (*ent.Lecturer, error)
	AddLecturerToClass(lecturerID, classID int) (string, error)
	GetAllClasses() ([]*ent.Class, error)
	//NewLecturerService(repo repositories.LecturerRepository) DefaultLecturerService
}	

type DefaultLecturerService struct {
	repo repositories.LecturerRepository
}

func (dl DefaultLecturerService) GetLecturerByID(id int) (*ent.Lecturer, error){
	lecturer, err := dl.repo.GetLecturerByID(id)
	if err != nil {
		return nil, fmt.Errorf("getting lecturer by id: %w", err)
	}
	return lecturer, nil
}

func (dl DefaultLecturerService) GetAllLecturers() ([]*ent.Lecturer, error){
	lecturers, err := dl.repo.GetAllLecturers()
	if err != nil {
		return lecturers, fmt.Errorf("getting all lecturers: %w", err)
	}
	return lecturers, nil
}

func (dl DefaultLecturerService) CreateLecturer(lecturer *ent.Lecturer) (*ent.Lecturer, error){
	lecturer, err := dl.repo.CreateLecturer(lecturer)
	if err != nil {
		return nil, fmt.Errorf("creating lecturer: %w", err)
	}
	return lecturer, nil
}

func (dl DefaultLecturerService) GetAllClasses() ([]*ent.Class, error){
	classes, err := dl.repo.GetAllClasses()
	if err != nil {
		return nil, fmt.Errorf("getting all classes: %w", err)
	}
	return classes, nil
}

func (dl DefaultLecturerService) AddLecturerToClass(lecturerID, classID int) (string, error){
	message, err := dl.repo.AddLecturerToClass(lecturerID, classID)
	if err != nil {
		return "", fmt.Errorf("adding lecturer to class: %w", err)
	}
	return message, nil
}

func NewLecturerService(repo repositories.LecturerRepository) DefaultLecturerService {
	return DefaultLecturerService{repo}
}

