package schema

import (
	"time"

	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
)

// Board holds the schema definition for the Board entity.
type Board struct {
	ent.Schema
}

// Fields of the Board.
func (Board) Fields() []ent.Field {
	return []ent.Field{
		field.Uint("id").Unique(),
		field.String("title").NotEmpty(),
		field.Text("content"),
		field.String("writer").NotEmpty(),
		field.Uint("view").Default(0),
		field.Uint("recommend").Default(0),
		field.Bool("hot").Default(false),
		field.Time("created_date").Default(time.Now),
	}
}

// Edges of the Board.
func (Board) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("user", User.Type).
			Ref("boards").
			Unique().
			Required(),
	}
}
