package repositories

import (
	"context"
	"fmt"
	"lecturer/ent"
	"lecturer/ent/class"
	"lecturer/ent/lecturer"
	"lecturer/models"
)

type LecturerRepository interface {
	GetAllLecturers() ([]*ent.Lecturer, error)
	GetLecturerByID(id int) (*ent.Lecturer, error)
	CreateLecturer(lecturer models.LecturerPayload) (*ent.Lecturer, error)
	AddLecturerToClass(lecturerID, classID int) (string, error)
	GetAllClasses() ([]*ent.Class, error)
}

type LecturerRepositoryDefault struct {
	ctx context.Context
	client *ent.Client
}

func (dl LecturerRepositoryDefault) GetAllLecturers() ([]*ent.Lecturer, error) {
	lecturers, err := dl.client.Lecturer.Query().All(dl.ctx)

	if err != nil {
		return nil, fmt.Errorf("getting all lecturers: %w", err)
	}

	return lecturers, nil
}

func (dl LecturerRepositoryDefault) GetLecturerByID(id int) (*ent.Lecturer, error) {
	return dl.client.Lecturer.Query().Where(lecturer.ID(id)).Only(dl.ctx)
}

func (dl LecturerRepositoryDefault) CreateLecturer(lecturer models.LecturerPayload) (*ent.Lecturer, error) {
	return dl.client.Lecturer.Create().SetEmail(lecturer.Email).SetFirstName(lecturer.FirstName).SetLastName(lecturer.LastName).Save(dl.ctx)
}

func (dl LecturerRepositoryDefault) GetAllClasses() ([]*ent.Class, error) {
	return dl.client.Class.Query().All(dl.ctx)
}

func (dl LecturerRepositoryDefault) AddLecturerToClass(lecturerID, classID int) (string, error) {
	class, error := dl.client.Class.Query().Where(class.ID(classID)).Only(dl.ctx)
	lecturer, error := dl.client.Lecturer.Query().Where(lecturer.ID(lecturerID)).Only(dl.ctx)
	lect, error := lecturer.Update().AddClasses(class).Save(dl.ctx)

	if error != nil {
		fmt.Print(lect)
		return "Adding lecturer failed", fmt.Errorf("adding lecturer to class: %w", error)
	}

	return "Adding lecturer succeeded", nil
}

func NewLecturerRepository(ctx context.Context, client *ent.Client) LecturerRepositoryDefault {
	return LecturerRepositoryDefault{ctx, client}
}