package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	"github.com/google/uuid"
)

// CommunityCategory holds the schema definition for the CommunityCategory entity.
type CommunityCategory struct {
	ent.Schema
}

// Fields of the CommunityCategory.
func (CommunityCategory) Fields() []ent.Field {
	return []ent.Field{
		field.UUID("id", uuid.UUID{}).Default(uuid.New).Unique(),
		field.String("name"),
		field.Int("display_order"),
	}
}

// Edges of the CommunityCategory.
func (CommunityCategory) Edges() []ent.Edge {
	return []ent.Edge{
		// edge.To("community_category_mapping", CommunityCategoryMapping.Type),
		edge.To("community", Community.Type).StorageKey(edge.Table("community_category_mapping")),
	}
}
