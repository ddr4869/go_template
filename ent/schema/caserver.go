package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
)

// CaServer holds the schema definition for the CaServer entity.
type CaServer struct {
	ent.Schema
}

// Fields of the CaServer.
func (CaServer) Fields() []ent.Field {
	return []ent.Field{
		field.String("name").Unique().NotEmpty().Comment("CA 서버 이름"),
		// field.String("user").Unique().NotEmpty().Comment("CA 등록 유저"),
	}
}

// Edges of the CaServer.
func (CaServer) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("user", User.Type).
			Ref("caserver").
			Unique().
			// We add the "Required" method to the builder
			// to make this edge required on entity creation.
			// i.e. Card cannot be created without its owner.
			Required(),
	}
}
