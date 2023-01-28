package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
)

// Lecturer holds the schema definition for the Lecturer entity.
type Lecturer struct {
	ent.Schema
}

// Fields of the Lecturer.
func (Lecturer) Fields() []ent.Field {
	return []ent.Field{
		field.Int("id").Unique(),
		field.Int("person_id").Unique(),
		field.Time("deleted_at"),
	}
}

// Edges of the Lecturer.
func (Lecturer) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("class_lecturers", ClassLecturer.Type),
		edge.From("person", Person.Type),
	}
}
