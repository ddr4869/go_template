// Code generated by ent, DO NOT EDIT.

package board

import (
	"time"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
)

const (
	// Label holds the string label denoting the board type in the database.
	Label = "board"
	// FieldID holds the string denoting the id field in the database.
	FieldID = "id"
	// FieldTitle holds the string denoting the title field in the database.
	FieldTitle = "title"
	// FieldContent holds the string denoting the content field in the database.
	FieldContent = "content"
	// FieldWriter holds the string denoting the writer field in the database.
	FieldWriter = "writer"
	// FieldView holds the string denoting the view field in the database.
	FieldView = "view"
	// FieldRecommend holds the string denoting the recommend field in the database.
	FieldRecommend = "recommend"
	// FieldHot holds the string denoting the hot field in the database.
	FieldHot = "hot"
	// FieldCreatedDate holds the string denoting the created_date field in the database.
	FieldCreatedDate = "created_date"
	// EdgeUser holds the string denoting the user edge name in mutations.
	EdgeUser = "user"
	// UserFieldID holds the string denoting the ID field of the User.
	UserFieldID = "name"
	// Table holds the table name of the board in the database.
	Table = "boards"
	// UserTable is the table that holds the user relation/edge.
	UserTable = "boards"
	// UserInverseTable is the table name for the User entity.
	// It exists in this package in order to avoid circular dependency with the "user" package.
	UserInverseTable = "users"
	// UserColumn is the table column denoting the user relation/edge.
	UserColumn = "user_boards"
)

// Columns holds all SQL columns for board fields.
var Columns = []string{
	FieldID,
	FieldTitle,
	FieldContent,
	FieldWriter,
	FieldView,
	FieldRecommend,
	FieldHot,
	FieldCreatedDate,
}

// ForeignKeys holds the SQL foreign-keys that are owned by the "boards"
// table and are not defined as standalone fields in the schema.
var ForeignKeys = []string{
	"user_boards",
}

// ValidColumn reports if the column name is valid (part of the table columns).
func ValidColumn(column string) bool {
	for i := range Columns {
		if column == Columns[i] {
			return true
		}
	}
	for i := range ForeignKeys {
		if column == ForeignKeys[i] {
			return true
		}
	}
	return false
}

var (
	// TitleValidator is a validator for the "title" field. It is called by the builders before save.
	TitleValidator func(string) error
	// WriterValidator is a validator for the "writer" field. It is called by the builders before save.
	WriterValidator func(string) error
	// DefaultView holds the default value on creation for the "view" field.
	DefaultView uint
	// DefaultRecommend holds the default value on creation for the "recommend" field.
	DefaultRecommend uint
	// DefaultHot holds the default value on creation for the "hot" field.
	DefaultHot bool
	// DefaultCreatedDate holds the default value on creation for the "created_date" field.
	DefaultCreatedDate func() time.Time
)

// OrderOption defines the ordering options for the Board queries.
type OrderOption func(*sql.Selector)

// ByID orders the results by the id field.
func ByID(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldID, opts...).ToFunc()
}

// ByTitle orders the results by the title field.
func ByTitle(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldTitle, opts...).ToFunc()
}

// ByContent orders the results by the content field.
func ByContent(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldContent, opts...).ToFunc()
}

// ByWriter orders the results by the writer field.
func ByWriter(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldWriter, opts...).ToFunc()
}

// ByView orders the results by the view field.
func ByView(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldView, opts...).ToFunc()
}

// ByRecommend orders the results by the recommend field.
func ByRecommend(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldRecommend, opts...).ToFunc()
}

// ByHot orders the results by the hot field.
func ByHot(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldHot, opts...).ToFunc()
}

// ByCreatedDate orders the results by the created_date field.
func ByCreatedDate(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldCreatedDate, opts...).ToFunc()
}

// ByUserField orders the results by user field.
func ByUserField(field string, opts ...sql.OrderTermOption) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborTerms(s, newUserStep(), sql.OrderByField(field, opts...))
	}
}
func newUserStep() *sqlgraph.Step {
	return sqlgraph.NewStep(
		sqlgraph.From(Table, FieldID),
		sqlgraph.To(UserInverseTable, UserFieldID),
		sqlgraph.Edge(sqlgraph.M2O, true, UserTable, UserColumn),
	)
}
