// Code generated by ent, DO NOT EDIT.

package ent

import (
	"dbapp/ent/applicationcategory"
	"fmt"
	"strings"

	"entgo.io/ent"
	"entgo.io/ent/dialect/sql"
	"github.com/google/uuid"
)

// ApplicationCategory is the model entity for the ApplicationCategory schema.
type ApplicationCategory struct {
	config `json:"-"`
	// ID of the ent.
	ID uuid.UUID `json:"id,omitempty"`
	// Name holds the value of the "name" field.
	Name string `json:"name,omitempty"`
	// DisplayOrder holds the value of the "display_order" field.
	DisplayOrder int `json:"display_order,omitempty"`
	// Edges holds the relations/edges for other nodes in the graph.
	// The values are being populated by the ApplicationCategoryQuery when eager-loading is set.
	Edges        ApplicationCategoryEdges `json:"edges"`
	selectValues sql.SelectValues
}

// ApplicationCategoryEdges holds the relations/edges for other nodes in the graph.
type ApplicationCategoryEdges struct {
	// Application holds the value of the application edge.
	Application []*Application `json:"application,omitempty"`
	// loadedTypes holds the information for reporting if a
	// type was loaded (or requested) in eager-loading or not.
	loadedTypes [1]bool
}

// ApplicationOrErr returns the Application value or an error if the edge
// was not loaded in eager-loading.
func (e ApplicationCategoryEdges) ApplicationOrErr() ([]*Application, error) {
	if e.loadedTypes[0] {
		return e.Application, nil
	}
	return nil, &NotLoadedError{edge: "application"}
}

// scanValues returns the types for scanning values from sql.Rows.
func (*ApplicationCategory) scanValues(columns []string) ([]any, error) {
	values := make([]any, len(columns))
	for i := range columns {
		switch columns[i] {
		case applicationcategory.FieldDisplayOrder:
			values[i] = new(sql.NullInt64)
		case applicationcategory.FieldName:
			values[i] = new(sql.NullString)
		case applicationcategory.FieldID:
			values[i] = new(uuid.UUID)
		default:
			values[i] = new(sql.UnknownType)
		}
	}
	return values, nil
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the ApplicationCategory fields.
func (ac *ApplicationCategory) assignValues(columns []string, values []any) error {
	if m, n := len(values), len(columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	for i := range columns {
		switch columns[i] {
		case applicationcategory.FieldID:
			if value, ok := values[i].(*uuid.UUID); !ok {
				return fmt.Errorf("unexpected type %T for field id", values[i])
			} else if value != nil {
				ac.ID = *value
			}
		case applicationcategory.FieldName:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field name", values[i])
			} else if value.Valid {
				ac.Name = value.String
			}
		case applicationcategory.FieldDisplayOrder:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field display_order", values[i])
			} else if value.Valid {
				ac.DisplayOrder = int(value.Int64)
			}
		default:
			ac.selectValues.Set(columns[i], values[i])
		}
	}
	return nil
}

// Value returns the ent.Value that was dynamically selected and assigned to the ApplicationCategory.
// This includes values selected through modifiers, order, etc.
func (ac *ApplicationCategory) Value(name string) (ent.Value, error) {
	return ac.selectValues.Get(name)
}

// QueryApplication queries the "application" edge of the ApplicationCategory entity.
func (ac *ApplicationCategory) QueryApplication() *ApplicationQuery {
	return NewApplicationCategoryClient(ac.config).QueryApplication(ac)
}

// Update returns a builder for updating this ApplicationCategory.
// Note that you need to call ApplicationCategory.Unwrap() before calling this method if this ApplicationCategory
// was returned from a transaction, and the transaction was committed or rolled back.
func (ac *ApplicationCategory) Update() *ApplicationCategoryUpdateOne {
	return NewApplicationCategoryClient(ac.config).UpdateOne(ac)
}

// Unwrap unwraps the ApplicationCategory entity that was returned from a transaction after it was closed,
// so that all future queries will be executed through the driver which created the transaction.
func (ac *ApplicationCategory) Unwrap() *ApplicationCategory {
	_tx, ok := ac.config.driver.(*txDriver)
	if !ok {
		panic("ent: ApplicationCategory is not a transactional entity")
	}
	ac.config.driver = _tx.drv
	return ac
}

// String implements the fmt.Stringer.
func (ac *ApplicationCategory) String() string {
	var builder strings.Builder
	builder.WriteString("ApplicationCategory(")
	builder.WriteString(fmt.Sprintf("id=%v, ", ac.ID))
	builder.WriteString("name=")
	builder.WriteString(ac.Name)
	builder.WriteString(", ")
	builder.WriteString("display_order=")
	builder.WriteString(fmt.Sprintf("%v", ac.DisplayOrder))
	builder.WriteByte(')')
	return builder.String()
}

// ApplicationCategories is a parsable slice of ApplicationCategory.
type ApplicationCategories []*ApplicationCategory