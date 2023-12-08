package schema

import (
	"time"

	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
)

// User holds the schema definition for the User entity.
type User struct {
	ent.Schema
}

// Fields of the User.
func (User) Fields() []ent.Field {
	return []ent.Field{
		field.String("id").StorageKey("name").NotEmpty().Unique(),
		field.Bytes("password").MaxLen(255).NotEmpty(),
		// field.String("email").NotEmpty().Unique(),
		// field.String("tel").NotEmpty().Unique(),
		// field.String("postcode").NotEmpty(),
		// field.String("address").NotEmpty(),
		// field.String("detail_address").NotEmpty(),
		field.String("grade").NotEmpty().Default("BRONZE"),
		// field.String("extraAddress").NotEmpty().Unique(),
		// field.Bool("delete").Default(false),
		field.String("description").Default("none"),
		field.Time("created_date").Default(time.Now()),
	}
}

// Edges of the User.
func (User) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("caserver", CaServer.Type),
		edge.To("boards", Board.Type),
		edge.To("payment", Payment.Type),
	}
}
