package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	"github.com/google/uuid"
)

// ApplicationCategory holds the schema definition for the ApplicationCategory entity.
type ApplicationCategory struct {
	ent.Schema
}

// Fields of the ApplicationCategory.
func (ApplicationCategory) Fields() []ent.Field {
	return []ent.Field{
		field.UUID("id", uuid.UUID{}).Default(uuid.New).Unique(),
		field.String("name"),
		field.Int("display_order"),
	}
}

// Edges of the ApplicationCategory.
func (ApplicationCategory) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("application", Application.Type).StorageKey(edge.Column("category_id")),
	}
}
