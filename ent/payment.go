// Code generated by ent, DO NOT EDIT.

package ent

import (
	"fmt"
	"strings"
	"time"

	"entgo.io/ent"
	"entgo.io/ent/dialect/sql"
	"github.com/go-board/ent/payment"
	"github.com/go-board/ent/user"
)

// Payment is the model entity for the Payment schema.
type Payment struct {
	config `json:"-"`
	// ID of the ent.
	ID int `json:"id,omitempty"`
	// UserName holds the value of the "user_name" field.
	UserName string `json:"user_name,omitempty"`
	// Grade holds the value of the "grade" field.
	Grade string `json:"grade,omitempty"`
	// MovieID holds the value of the "movie_id" field.
	MovieID string `json:"movie_id,omitempty"`
	// CreatedDate holds the value of the "created_date" field.
	CreatedDate time.Time `json:"created_date,omitempty"`
	// Edges holds the relations/edges for other nodes in the graph.
	// The values are being populated by the PaymentQuery when eager-loading is set.
	Edges        PaymentEdges `json:"edges"`
	user_payment *string
	selectValues sql.SelectValues
}

// PaymentEdges holds the relations/edges for other nodes in the graph.
type PaymentEdges struct {
	// User holds the value of the user edge.
	User *User `json:"user,omitempty"`
	// loadedTypes holds the information for reporting if a
	// type was loaded (or requested) in eager-loading or not.
	loadedTypes [1]bool
}

// UserOrErr returns the User value or an error if the edge
// was not loaded in eager-loading, or loaded but was not found.
func (e PaymentEdges) UserOrErr() (*User, error) {
	if e.loadedTypes[0] {
		if e.User == nil {
			// Edge was loaded but was not found.
			return nil, &NotFoundError{label: user.Label}
		}
		return e.User, nil
	}
	return nil, &NotLoadedError{edge: "user"}
}

// scanValues returns the types for scanning values from sql.Rows.
func (*Payment) scanValues(columns []string) ([]any, error) {
	values := make([]any, len(columns))
	for i := range columns {
		switch columns[i] {
		case payment.FieldID:
			values[i] = new(sql.NullInt64)
		case payment.FieldUserName, payment.FieldGrade, payment.FieldMovieID:
			values[i] = new(sql.NullString)
		case payment.FieldCreatedDate:
			values[i] = new(sql.NullTime)
		case payment.ForeignKeys[0]: // user_payment
			values[i] = new(sql.NullString)
		default:
			values[i] = new(sql.UnknownType)
		}
	}
	return values, nil
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the Payment fields.
func (pa *Payment) assignValues(columns []string, values []any) error {
	if m, n := len(values), len(columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	for i := range columns {
		switch columns[i] {
		case payment.FieldID:
			value, ok := values[i].(*sql.NullInt64)
			if !ok {
				return fmt.Errorf("unexpected type %T for field id", value)
			}
			pa.ID = int(value.Int64)
		case payment.FieldUserName:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field user_name", values[i])
			} else if value.Valid {
				pa.UserName = value.String
			}
		case payment.FieldGrade:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field grade", values[i])
			} else if value.Valid {
				pa.Grade = value.String
			}
		case payment.FieldMovieID:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field movie_id", values[i])
			} else if value.Valid {
				pa.MovieID = value.String
			}
		case payment.FieldCreatedDate:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field created_date", values[i])
			} else if value.Valid {
				pa.CreatedDate = value.Time
			}
		case payment.ForeignKeys[0]:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field user_payment", values[i])
			} else if value.Valid {
				pa.user_payment = new(string)
				*pa.user_payment = value.String
			}
		default:
			pa.selectValues.Set(columns[i], values[i])
		}
	}
	return nil
}

// Value returns the ent.Value that was dynamically selected and assigned to the Payment.
// This includes values selected through modifiers, order, etc.
func (pa *Payment) Value(name string) (ent.Value, error) {
	return pa.selectValues.Get(name)
}

// QueryUser queries the "user" edge of the Payment entity.
func (pa *Payment) QueryUser() *UserQuery {
	return NewPaymentClient(pa.config).QueryUser(pa)
}

// Update returns a builder for updating this Payment.
// Note that you need to call Payment.Unwrap() before calling this method if this Payment
// was returned from a transaction, and the transaction was committed or rolled back.
func (pa *Payment) Update() *PaymentUpdateOne {
	return NewPaymentClient(pa.config).UpdateOne(pa)
}

// Unwrap unwraps the Payment entity that was returned from a transaction after it was closed,
// so that all future queries will be executed through the driver which created the transaction.
func (pa *Payment) Unwrap() *Payment {
	_tx, ok := pa.config.driver.(*txDriver)
	if !ok {
		panic("ent: Payment is not a transactional entity")
	}
	pa.config.driver = _tx.drv
	return pa
}

// String implements the fmt.Stringer.
func (pa *Payment) String() string {
	var builder strings.Builder
	builder.WriteString("Payment(")
	builder.WriteString(fmt.Sprintf("id=%v, ", pa.ID))
	builder.WriteString("user_name=")
	builder.WriteString(pa.UserName)
	builder.WriteString(", ")
	builder.WriteString("grade=")
	builder.WriteString(pa.Grade)
	builder.WriteString(", ")
	builder.WriteString("movie_id=")
	builder.WriteString(pa.MovieID)
	builder.WriteString(", ")
	builder.WriteString("created_date=")
	builder.WriteString(pa.CreatedDate.Format(time.ANSIC))
	builder.WriteByte(')')
	return builder.String()
}

// Payments is a parsable slice of Payment.
type Payments []*Payment
