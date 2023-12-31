// Code generated by ent, DO NOT EDIT.

package adgroup

import (
	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"github.com/google/uuid"
)

const (
	// Label holds the string label denoting the adgroup type in the database.
	Label = "ad_group"
	// FieldID holds the string denoting the id field in the database.
	FieldID = "id"
	// FieldName holds the string denoting the name field in the database.
	FieldName = "name"
	// EdgeCommunity holds the string denoting the community edge name in mutations.
	EdgeCommunity = "community"
	// Table holds the table name of the adgroup in the database.
	Table = "ad_groups"
	// CommunityTable is the table that holds the community relation/edge. The primary key declared below.
	CommunityTable = "adgroup_community_mapping"
	// CommunityInverseTable is the table name for the Community entity.
	// It exists in this package in order to avoid circular dependency with the "community" package.
	CommunityInverseTable = "communities"
)

// Columns holds all SQL columns for adgroup fields.
var Columns = []string{
	FieldID,
	FieldName,
}

var (
	// CommunityPrimaryKey and CommunityColumn2 are the table columns denoting the
	// primary key for the community relation (M2M).
	CommunityPrimaryKey = []string{"ad_group_id", "community_id"}
)

// ValidColumn reports if the column name is valid (part of the table columns).
func ValidColumn(column string) bool {
	for i := range Columns {
		if column == Columns[i] {
			return true
		}
	}
	return false
}

var (
	// DefaultID holds the default value on creation for the "id" field.
	DefaultID func() uuid.UUID
)

// OrderOption defines the ordering options for the ADGroup queries.
type OrderOption func(*sql.Selector)

// ByID orders the results by the id field.
func ByID(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldID, opts...).ToFunc()
}

// ByName orders the results by the name field.
func ByName(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldName, opts...).ToFunc()
}

// ByCommunityCount orders the results by community count.
func ByCommunityCount(opts ...sql.OrderTermOption) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborsCount(s, newCommunityStep(), opts...)
	}
}

// ByCommunity orders the results by community terms.
func ByCommunity(term sql.OrderTerm, terms ...sql.OrderTerm) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborTerms(s, newCommunityStep(), append([]sql.OrderTerm{term}, terms...)...)
	}
}
func newCommunityStep() *sqlgraph.Step {
	return sqlgraph.NewStep(
		sqlgraph.From(Table, FieldID),
		sqlgraph.To(CommunityInverseTable, FieldID),
		sqlgraph.Edge(sqlgraph.M2M, false, CommunityTable, CommunityPrimaryKey...),
	)
}
