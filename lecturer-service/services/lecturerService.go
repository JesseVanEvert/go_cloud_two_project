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
}

type DefaultLecturerService struct {
	repo repositories.LecturerRepository
	cl  *ent.Client
}

func GetAllLecturers(dl DefaultLecturerService) ([]*ent.Lecturer, error){
	lecturers, err := dl.repo.GetAllLecturers(dl.cl)
	if err != nil {
		return nil, fmt.Errorf("getting all lecturers: %w", err)
	}
	return lecturers, nil
}