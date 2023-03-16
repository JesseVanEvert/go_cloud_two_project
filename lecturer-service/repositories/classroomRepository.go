package repositories

import (
	"context"
	"lecturer/ent"
	"lecturer/models"
)

type ClassRoomRepository interface {
   CreateClassRoom(classroom models.ClassRoom) (*ent.Class, error)
   DeleteClassRoom(id int) (string, error)
   UpdateClassRoom(classroom models.ClassRoom) (*ent.Class, error)
}

type DefaultClassRoomRepository struct {
	ctx context.Context
	client *ent.Client
}

func (dl DefaultClassRoomRepository) CreateClassRoom(classroom models.ClassRoom) (*ent.Class, error) {
	 return dl.client.Class.Create().SetName(classroom.Classname).Save(dl.ctx)
}

func (dl DefaultClassRoomRepository) DeleteClassRoom(id int) (string, error) {
	dl.client.Class.DeleteOneID(id).Exec(dl.ctx)
	return "Classroom deleted", nil
}

func (dl DefaultClassRoomRepository) UpdateClassRoom(classroom models.ClassRoom) (*ent.Class, error) {
	return dl.client.Class.UpdateOneID(classroom.ID).SetName(classroom.Classname).Save(dl.ctx)
}

func NewClassRoomRepository(ctx context.Context, client *ent.Client) ClassRoomRepository {
	return DefaultClassRoomRepository{ctx: ctx, client: client}
}
	
