package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
)

// ClassLecturer holds the schema definition for the ClassLecturer entity.
type ClassLecturer struct {
	ent.Schema
}

// Fields of the ClassLecturer.
func (ClassLecturer) Fields() []ent.Field {
	return []ent.Field{
		field.Time("deleted_at"),
	}
}

// Edges of the ClassLecturer.
func (ClassLecturer) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("class", Class.Type).Ref("class_lecturers").Unique(),
		edge.From("lecturer", Lecturer.Type).Ref("class_lecturers").Unique(),
	}
}
