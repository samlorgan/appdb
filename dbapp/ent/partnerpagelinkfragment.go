// Code generated by ent, DO NOT EDIT.

package ent

import (
	"dbapp/ent/partnerpagelinkfragment"
	"fmt"
	"strings"

	"entgo.io/ent"
	"entgo.io/ent/dialect/sql"
	"github.com/google/uuid"
)

// PartnerPageLinkFragment is the model entity for the PartnerPageLinkFragment schema.
type PartnerPageLinkFragment struct {
	config `json:"-"`
	// ID of the ent.
	ID uuid.UUID `json:"id,omitempty"`
	// LinkText holds the value of the "link_text" field.
	LinkText string `json:"link_text,omitempty"`
	// WhamPartnerURLSuffix holds the value of the "wham_partner_url_suffix" field.
	WhamPartnerURLSuffix string `json:"wham_partner_url_suffix,omitempty"`
	// DisplayOrder holds the value of the "display_order" field.
	DisplayOrder int `json:"display_order,omitempty"`
	selectValues sql.SelectValues
}

// scanValues returns the types for scanning values from sql.Rows.
func (*PartnerPageLinkFragment) scanValues(columns []string) ([]any, error) {
	values := make([]any, len(columns))
	for i := range columns {
		switch columns[i] {
		case partnerpagelinkfragment.FieldDisplayOrder:
			values[i] = new(sql.NullInt64)
		case partnerpagelinkfragment.FieldLinkText, partnerpagelinkfragment.FieldWhamPartnerURLSuffix:
			values[i] = new(sql.NullString)
		case partnerpagelinkfragment.FieldID:
			values[i] = new(uuid.UUID)
		default:
			values[i] = new(sql.UnknownType)
		}
	}
	return values, nil
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the PartnerPageLinkFragment fields.
func (pplf *PartnerPageLinkFragment) assignValues(columns []string, values []any) error {
	if m, n := len(values), len(columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	for i := range columns {
		switch columns[i] {
		case partnerpagelinkfragment.FieldID:
			if value, ok := values[i].(*uuid.UUID); !ok {
				return fmt.Errorf("unexpected type %T for field id", values[i])
			} else if value != nil {
				pplf.ID = *value
			}
		case partnerpagelinkfragment.FieldLinkText:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field link_text", values[i])
			} else if value.Valid {
				pplf.LinkText = value.String
			}
		case partnerpagelinkfragment.FieldWhamPartnerURLSuffix:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field wham_partner_url_suffix", values[i])
			} else if value.Valid {
				pplf.WhamPartnerURLSuffix = value.String
			}
		case partnerpagelinkfragment.FieldDisplayOrder:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field display_order", values[i])
			} else if value.Valid {
				pplf.DisplayOrder = int(value.Int64)
			}
		default:
			pplf.selectValues.Set(columns[i], values[i])
		}
	}
	return nil
}

// Value returns the ent.Value that was dynamically selected and assigned to the PartnerPageLinkFragment.
// This includes values selected through modifiers, order, etc.
func (pplf *PartnerPageLinkFragment) Value(name string) (ent.Value, error) {
	return pplf.selectValues.Get(name)
}

// Update returns a builder for updating this PartnerPageLinkFragment.
// Note that you need to call PartnerPageLinkFragment.Unwrap() before calling this method if this PartnerPageLinkFragment
// was returned from a transaction, and the transaction was committed or rolled back.
func (pplf *PartnerPageLinkFragment) Update() *PartnerPageLinkFragmentUpdateOne {
	return NewPartnerPageLinkFragmentClient(pplf.config).UpdateOne(pplf)
}

// Unwrap unwraps the PartnerPageLinkFragment entity that was returned from a transaction after it was closed,
// so that all future queries will be executed through the driver which created the transaction.
func (pplf *PartnerPageLinkFragment) Unwrap() *PartnerPageLinkFragment {
	_tx, ok := pplf.config.driver.(*txDriver)
	if !ok {
		panic("ent: PartnerPageLinkFragment is not a transactional entity")
	}
	pplf.config.driver = _tx.drv
	return pplf
}

// String implements the fmt.Stringer.
func (pplf *PartnerPageLinkFragment) String() string {
	var builder strings.Builder
	builder.WriteString("PartnerPageLinkFragment(")
	builder.WriteString(fmt.Sprintf("id=%v, ", pplf.ID))
	builder.WriteString("link_text=")
	builder.WriteString(pplf.LinkText)
	builder.WriteString(", ")
	builder.WriteString("wham_partner_url_suffix=")
	builder.WriteString(pplf.WhamPartnerURLSuffix)
	builder.WriteString(", ")
	builder.WriteString("display_order=")
	builder.WriteString(fmt.Sprintf("%v", pplf.DisplayOrder))
	builder.WriteByte(')')
	return builder.String()
}

// PartnerPageLinkFragments is a parsable slice of PartnerPageLinkFragment.
type PartnerPageLinkFragments []*PartnerPageLinkFragment
