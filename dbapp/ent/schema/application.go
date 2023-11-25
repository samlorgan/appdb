package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	"github.com/google/uuid"
)

// Application holds the schema definition for the Application entity.
type Application struct {
	ent.Schema
}

// Fields of the Application.
func (Application) Fields() []ent.Field {
	return []ent.Field{
		field.UUID("id", uuid.UUID{}).Default(uuid.New).Unique(),
		field.String("name"),
		field.String("description"),
		field.String("alt_text"),
		field.String("uri"),
		field.String("icon_uri"),
		field.Bool("is_favourite"),
		field.Time("valid_from"),
		field.Time("valid_to"),
	}
}

// Edges of the Application.
func (Application) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("application_category", ApplicationCategory.Type).Ref("application").Unique().Required(),
		edge.To("partner", Partner.Type).StorageKey(edge.Table("application_partner_access")),
	}
}
