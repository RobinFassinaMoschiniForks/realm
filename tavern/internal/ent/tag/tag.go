// Code generated by ent, DO NOT EDIT.

package tag

import (
	"fmt"
	"io"
	"strconv"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
)

const (
	// Label holds the string label denoting the tag type in the database.
	Label = "tag"
	// FieldID holds the string denoting the id field in the database.
	FieldID = "id"
	// FieldName holds the string denoting the name field in the database.
	FieldName = "name"
	// FieldKind holds the string denoting the kind field in the database.
	FieldKind = "kind"
	// EdgeHosts holds the string denoting the hosts edge name in mutations.
	EdgeHosts = "hosts"
	// Table holds the table name of the tag in the database.
	Table = "tags"
	// HostsTable is the table that holds the hosts relation/edge. The primary key declared below.
	HostsTable = "host_tags"
	// HostsInverseTable is the table name for the Host entity.
	// It exists in this package in order to avoid circular dependency with the "host" package.
	HostsInverseTable = "hosts"
)

// Columns holds all SQL columns for tag fields.
var Columns = []string{
	FieldID,
	FieldName,
	FieldKind,
}

var (
	// HostsPrimaryKey and HostsColumn2 are the table columns denoting the
	// primary key for the hosts relation (M2M).
	HostsPrimaryKey = []string{"host_id", "tag_id"}
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
	// NameValidator is a validator for the "name" field. It is called by the builders before save.
	NameValidator func(string) error
)

// Kind defines the type for the "kind" enum field.
type Kind string

// Kind values.
const (
	KindGroup   Kind = "group"
	KindService Kind = "service"
)

func (k Kind) String() string {
	return string(k)
}

// KindValidator is a validator for the "kind" field enum values. It is called by the builders before save.
func KindValidator(k Kind) error {
	switch k {
	case KindGroup, KindService:
		return nil
	default:
		return fmt.Errorf("tag: invalid enum value for kind field: %q", k)
	}
}

// OrderOption defines the ordering options for the Tag queries.
type OrderOption func(*sql.Selector)

// ByID orders the results by the id field.
func ByID(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldID, opts...).ToFunc()
}

// ByName orders the results by the name field.
func ByName(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldName, opts...).ToFunc()
}

// ByKind orders the results by the kind field.
func ByKind(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldKind, opts...).ToFunc()
}

// ByHostsCount orders the results by hosts count.
func ByHostsCount(opts ...sql.OrderTermOption) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborsCount(s, newHostsStep(), opts...)
	}
}

// ByHosts orders the results by hosts terms.
func ByHosts(term sql.OrderTerm, terms ...sql.OrderTerm) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborTerms(s, newHostsStep(), append([]sql.OrderTerm{term}, terms...)...)
	}
}
func newHostsStep() *sqlgraph.Step {
	return sqlgraph.NewStep(
		sqlgraph.From(Table, FieldID),
		sqlgraph.To(HostsInverseTable, FieldID),
		sqlgraph.Edge(sqlgraph.M2M, true, HostsTable, HostsPrimaryKey...),
	)
}

// MarshalGQL implements graphql.Marshaler interface.
func (e Kind) MarshalGQL(w io.Writer) {
	io.WriteString(w, strconv.Quote(e.String()))
}

// UnmarshalGQL implements graphql.Unmarshaler interface.
func (e *Kind) UnmarshalGQL(val interface{}) error {
	str, ok := val.(string)
	if !ok {
		return fmt.Errorf("enum %T must be a string", val)
	}
	*e = Kind(str)
	if err := KindValidator(*e); err != nil {
		return fmt.Errorf("%s is not a valid Kind", str)
	}
	return nil
}
