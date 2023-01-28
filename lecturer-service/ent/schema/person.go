package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/field"
)

// Person holds the schema definition for the Person entity.
type Person struct {
	ent.Schema
}

// Fields of the Person.
func (Person) Fields() []ent.Field {
	return []ent.Field{
		field.Int("id").Unique(),
		field.String("first_name"),
		field.String("last_name"),
		field.String("email"),
		field.String("deleted_at"),
	}
}

// Edges of the Person.
func (Person) Edges() []ent.Edge {
	return nil
}
