package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	"github.com/google/uuid"
)

// Community holds the schema definition for the Community entity.
type Community struct {
	ent.Schema
}

// Fields of the Community.
func (Community) Fields() []ent.Field {
	return []ent.Field{
		field.UUID("id", uuid.UUID{}).Default(uuid.New).Unique(),
		field.Int("wham_site_id"),
		field.String("wham_title"),
		field.String("whan_description"),
		field.String("wham_community_url"),
		field.Time("wham_created"),
		field.Time("wham_updated"),
		field.Time("featured_from").Optional(),
		field.Time("featured_to").Optional(),
	}
}

// Edges of the Community.
func (Community) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("adgroup", ADGroup.Type).Ref("community"),
		edge.From("community_category", CommunityCategory.Type).Ref("community"),
	}

}
