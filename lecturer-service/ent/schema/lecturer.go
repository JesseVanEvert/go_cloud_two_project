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
		field.String("first_name"),
		field.String("last_name"),
		field.String("email"),
		field.String("deleted_at").Optional(),
	}
}

// Edges of the Lecturer.
func (Lecturer) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("classes", Class.Type).Ref("lecturers"),
	}
}
