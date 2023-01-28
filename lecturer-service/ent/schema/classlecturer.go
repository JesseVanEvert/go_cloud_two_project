package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/field"
)

// ClassLecturer holds the schema definition for the ClassLecturer entity.
type ClassLecturer struct {
	ent.Schema
}

// Fields of the ClassLecturer.
func (ClassLecturer) Fields() []ent.Field {
	return []ent.Field{
		field.Int("class_id").Unique(),
		field.Int("lecturer_id").Unique(),
		field.Time("deleted_at"),
	}
}

// Edges of the ClassLecturer.
func (ClassLecturer) Edges() []ent.Edge {
	return nil
}
