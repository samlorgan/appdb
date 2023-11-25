package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/field"
	"github.com/google/uuid"
)

// PartnerPageLinkFragment holds the schema definition for the PartnerPageLinkFragment entity.
type PartnerPageLinkFragment struct {
	ent.Schema
}

// Fields of the PartnerPageLinkFragment.
func (PartnerPageLinkFragment) Fields() []ent.Field {
	return []ent.Field{
		field.UUID("id", uuid.UUID{}).Default(uuid.New).Unique(),
		field.String("link_text"),
		field.String("wham_partner_url_suffix").Optional(),
		field.Int("display_order"),
	}
}

// Edges of the PartnerPageLinkFragment.
func (PartnerPageLinkFragment) Edges() []ent.Edge {
	return nil
}
