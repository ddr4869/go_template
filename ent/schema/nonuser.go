package schema

import (
	"time"

	"entgo.io/ent"
	"entgo.io/ent/schema/field"
)

// NonUser holds the schema definition for the NonUser entity.
type NonUser struct {
	ent.Schema
}

// Fields of the NonUser.
func (NonUser) Fields() []ent.Field {
	return []ent.Field{
		field.String("name").NotEmpty().Unique(),
		field.Bytes("password").MaxLen(255).NotEmpty(),
		field.String("tel").NotEmpty().Unique(),
		field.String("description").Default("none"),
		field.Time("created_date").Default(time.Now()),
	}
}

// Edges of the NonUser.
func (NonUser) Edges() []ent.Edge {
	return nil
}
