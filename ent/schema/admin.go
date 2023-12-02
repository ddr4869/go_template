package schema

import (
	"time"

	"entgo.io/ent"
	"entgo.io/ent/schema/field"
)

// Admin holds the schema definition for the Admin entity.
type Admin struct {
	ent.Schema
}

// Fields of the Admin.
func (Admin) Fields() []ent.Field {
	return []ent.Field{
		field.String("name").NotEmpty().Unique(),
		field.Bytes("password").MaxLen(255).NotEmpty(),
		field.Time("created_date").Default(time.Now()),
	}
}

// Edges of the Admin.
func (Admin) Edges() []ent.Edge {
	return nil
}
