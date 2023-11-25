package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	"github.com/google/uuid"
)

// Partner holds the schema definition for the Partner entity.
type Partner struct {
	ent.Schema
}

// Fields of the Partner.
func (Partner) Fields() []ent.Field {
	return []ent.Field{
		field.UUID("id", uuid.UUID{}).Default(uuid.New).Unique(),
		field.Int("wham_site_id"),
		field.String("wham_title"),
		field.String("wham_description"),
		field.String("keycloak_organisation"),
		field.String("wham_partner_url"),
		field.Time("wham_created"),
		field.Time("wham_updated"),
	}
}

// Edges of the Partner.
func (Partner) Edges() []ent.Edge {
	return []ent.Edge{
		// edge.To("partner", Partner.Type),
		edge.From("application", Application.Type).Ref("partner"),
	}
}
