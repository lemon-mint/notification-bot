package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
)

// Channel holds the schema definition for the Channel entity.
type Channel struct {
	ent.Schema
}

// Fields of the Channel.
func (Channel) Fields() []ent.Field {
	return []ent.Field{
		field.String("ChannelID").Unique(),
		field.String("Name"),
	}
}

// Edges of the Channel.
func (Channel) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("Subscribers", User.Type),
	}
}
