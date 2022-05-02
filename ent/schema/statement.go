package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
)

// Statement holds the schema definition for the Statement entity.
type Statement struct {
	ent.Schema
}

// Fields of the Statement.
func (Statement) Fields() []ent.Field {
	return []ent.Field{}
}

// Edges of the Statement.
func (Statement) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("subjects", Subject.Type),

		edge.From("dsse", Dsse.Type).Ref("statement"),
	}
}
