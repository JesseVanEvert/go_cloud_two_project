package repositories

import (
	"context"
	"fmt"
	"lecturer/ent"
	"lecturer/ent/lecturer"
)

type LecturerRepository interface {
	GetAllLecturers(client *ent.Client) ([]*ent.Lecturer, error)
	GetLecturerByID(client *ent.Client, id int) (*ent.Lecturer, error)
	CreateLecturer(client *ent.Client, lecturer *ent.Lecturer) (*ent.Lecturer, error)
	AddLecturerToClass(client *ent.Client, lecturerID, classID int) (string, error)
}
func GetAllLecturers(ctx context.Context, client *ent.Client) ([]*ent.Lecturer, error) {
	return client.Lecturer.Query().All(ctx)
}

func GetLecturerByID(ctx context.Context, client *ent.Client, id int) (*ent.Lecturer, error) {
	return client.Lecturer.Query().Where(lecturer.ID(id)).Only(ctx)
}

func createLecturer(ctx context.Context, client *ent.Client, lecturer *ent.Lecturer) (*ent.Lecturer, error) {
	return client.Lecturer.Create().SetEmail(lecturer.Email).SetFirstName(lecturer.FirstName).SetLastName(lecturer.LastName).Save(ctx)
}

func addLecturerToClass(ctx context.Context, client *ent.Client, lecturerID, classID int) (string, error) {
	_, err := client.Class.UpdateOneID(classID).AddClassLecturerIDs(lecturerID).Save(ctx)
	if err != nil {
		return "Adding lecturer failed", fmt.Errorf("adding lecturer to class: %w", err)
	}
	return "Adding lecturer succeeded", nil
}