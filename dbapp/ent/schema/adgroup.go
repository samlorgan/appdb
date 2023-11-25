package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	"github.com/google/uuid"
)

// ADGroup holds the schema definition for the ADGroup entity.
type ADGroup struct {
	ent.Schema
}

// Fields of the ADGroup.
func (ADGroup) Fields() []ent.Field {
	return []ent.Field{
		field.UUID("id", uuid.UUID{}).Default(uuid.New).Unique(),
		field.String("name"),
	}
}

// Edges of the ADGroup.
func (ADGroup) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("community", Community.Type).StorageKey(edge.Table("adgroup_community_mapping")),
	}
}
