// Code generated by ent, DO NOT EDIT.

package user

import (
	"time"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"github.com/go-board/ent/predicate"
)

// ID filters vertices based on their ID field.
func ID(id string) predicate.User {
	return predicate.User(sql.FieldEQ(FieldID, id))
}

// IDEQ applies the EQ predicate on the ID field.
func IDEQ(id string) predicate.User {
	return predicate.User(sql.FieldEQ(FieldID, id))
}

// IDNEQ applies the NEQ predicate on the ID field.
func IDNEQ(id string) predicate.User {
	return predicate.User(sql.FieldNEQ(FieldID, id))
}

// IDIn applies the In predicate on the ID field.
func IDIn(ids ...string) predicate.User {
	return predicate.User(sql.FieldIn(FieldID, ids...))
}

// IDNotIn applies the NotIn predicate on the ID field.
func IDNotIn(ids ...string) predicate.User {
	return predicate.User(sql.FieldNotIn(FieldID, ids...))
}

// IDGT applies the GT predicate on the ID field.
func IDGT(id string) predicate.User {
	return predicate.User(sql.FieldGT(FieldID, id))
}

// IDGTE applies the GTE predicate on the ID field.
func IDGTE(id string) predicate.User {
	return predicate.User(sql.FieldGTE(FieldID, id))
}

// IDLT applies the LT predicate on the ID field.
func IDLT(id string) predicate.User {
	return predicate.User(sql.FieldLT(FieldID, id))
}

// IDLTE applies the LTE predicate on the ID field.
func IDLTE(id string) predicate.User {
	return predicate.User(sql.FieldLTE(FieldID, id))
}

// IDEqualFold applies the EqualFold predicate on the ID field.
func IDEqualFold(id string) predicate.User {
	return predicate.User(sql.FieldEqualFold(FieldID, id))
}

// IDContainsFold applies the ContainsFold predicate on the ID field.
func IDContainsFold(id string) predicate.User {
	return predicate.User(sql.FieldContainsFold(FieldID, id))
}

// Password applies equality check predicate on the "password" field. It's identical to PasswordEQ.
func Password(v []byte) predicate.User {
	return predicate.User(sql.FieldEQ(FieldPassword, v))
}

// Grade applies equality check predicate on the "grade" field. It's identical to GradeEQ.
func Grade(v string) predicate.User {
	return predicate.User(sql.FieldEQ(FieldGrade, v))
}

// Description applies equality check predicate on the "description" field. It's identical to DescriptionEQ.
func Description(v string) predicate.User {
	return predicate.User(sql.FieldEQ(FieldDescription, v))
}

// CreatedDate applies equality check predicate on the "created_date" field. It's identical to CreatedDateEQ.
func CreatedDate(v time.Time) predicate.User {
	return predicate.User(sql.FieldEQ(FieldCreatedDate, v))
}

// PasswordEQ applies the EQ predicate on the "password" field.
func PasswordEQ(v []byte) predicate.User {
	return predicate.User(sql.FieldEQ(FieldPassword, v))
}

// PasswordNEQ applies the NEQ predicate on the "password" field.
func PasswordNEQ(v []byte) predicate.User {
	return predicate.User(sql.FieldNEQ(FieldPassword, v))
}

// PasswordIn applies the In predicate on the "password" field.
func PasswordIn(vs ...[]byte) predicate.User {
	return predicate.User(sql.FieldIn(FieldPassword, vs...))
}

// PasswordNotIn applies the NotIn predicate on the "password" field.
func PasswordNotIn(vs ...[]byte) predicate.User {
	return predicate.User(sql.FieldNotIn(FieldPassword, vs...))
}

// PasswordGT applies the GT predicate on the "password" field.
func PasswordGT(v []byte) predicate.User {
	return predicate.User(sql.FieldGT(FieldPassword, v))
}

// PasswordGTE applies the GTE predicate on the "password" field.
func PasswordGTE(v []byte) predicate.User {
	return predicate.User(sql.FieldGTE(FieldPassword, v))
}

// PasswordLT applies the LT predicate on the "password" field.
func PasswordLT(v []byte) predicate.User {
	return predicate.User(sql.FieldLT(FieldPassword, v))
}

// PasswordLTE applies the LTE predicate on the "password" field.
func PasswordLTE(v []byte) predicate.User {
	return predicate.User(sql.FieldLTE(FieldPassword, v))
}

// GradeEQ applies the EQ predicate on the "grade" field.
func GradeEQ(v string) predicate.User {
	return predicate.User(sql.FieldEQ(FieldGrade, v))
}

// GradeNEQ applies the NEQ predicate on the "grade" field.
func GradeNEQ(v string) predicate.User {
	return predicate.User(sql.FieldNEQ(FieldGrade, v))
}

// GradeIn applies the In predicate on the "grade" field.
func GradeIn(vs ...string) predicate.User {
	return predicate.User(sql.FieldIn(FieldGrade, vs...))
}

// GradeNotIn applies the NotIn predicate on the "grade" field.
func GradeNotIn(vs ...string) predicate.User {
	return predicate.User(sql.FieldNotIn(FieldGrade, vs...))
}

// GradeGT applies the GT predicate on the "grade" field.
func GradeGT(v string) predicate.User {
	return predicate.User(sql.FieldGT(FieldGrade, v))
}

// GradeGTE applies the GTE predicate on the "grade" field.
func GradeGTE(v string) predicate.User {
	return predicate.User(sql.FieldGTE(FieldGrade, v))
}

// GradeLT applies the LT predicate on the "grade" field.
func GradeLT(v string) predicate.User {
	return predicate.User(sql.FieldLT(FieldGrade, v))
}

// GradeLTE applies the LTE predicate on the "grade" field.
func GradeLTE(v string) predicate.User {
	return predicate.User(sql.FieldLTE(FieldGrade, v))
}

// GradeContains applies the Contains predicate on the "grade" field.
func GradeContains(v string) predicate.User {
	return predicate.User(sql.FieldContains(FieldGrade, v))
}

// GradeHasPrefix applies the HasPrefix predicate on the "grade" field.
func GradeHasPrefix(v string) predicate.User {
	return predicate.User(sql.FieldHasPrefix(FieldGrade, v))
}

// GradeHasSuffix applies the HasSuffix predicate on the "grade" field.
func GradeHasSuffix(v string) predicate.User {
	return predicate.User(sql.FieldHasSuffix(FieldGrade, v))
}

// GradeEqualFold applies the EqualFold predicate on the "grade" field.
func GradeEqualFold(v string) predicate.User {
	return predicate.User(sql.FieldEqualFold(FieldGrade, v))
}

// GradeContainsFold applies the ContainsFold predicate on the "grade" field.
func GradeContainsFold(v string) predicate.User {
	return predicate.User(sql.FieldContainsFold(FieldGrade, v))
}

// DescriptionEQ applies the EQ predicate on the "description" field.
func DescriptionEQ(v string) predicate.User {
	return predicate.User(sql.FieldEQ(FieldDescription, v))
}

// DescriptionNEQ applies the NEQ predicate on the "description" field.
func DescriptionNEQ(v string) predicate.User {
	return predicate.User(sql.FieldNEQ(FieldDescription, v))
}

// DescriptionIn applies the In predicate on the "description" field.
func DescriptionIn(vs ...string) predicate.User {
	return predicate.User(sql.FieldIn(FieldDescription, vs...))
}

// DescriptionNotIn applies the NotIn predicate on the "description" field.
func DescriptionNotIn(vs ...string) predicate.User {
	return predicate.User(sql.FieldNotIn(FieldDescription, vs...))
}

// DescriptionGT applies the GT predicate on the "description" field.
func DescriptionGT(v string) predicate.User {
	return predicate.User(sql.FieldGT(FieldDescription, v))
}

// DescriptionGTE applies the GTE predicate on the "description" field.
func DescriptionGTE(v string) predicate.User {
	return predicate.User(sql.FieldGTE(FieldDescription, v))
}

// DescriptionLT applies the LT predicate on the "description" field.
func DescriptionLT(v string) predicate.User {
	return predicate.User(sql.FieldLT(FieldDescription, v))
}

// DescriptionLTE applies the LTE predicate on the "description" field.
func DescriptionLTE(v string) predicate.User {
	return predicate.User(sql.FieldLTE(FieldDescription, v))
}

// DescriptionContains applies the Contains predicate on the "description" field.
func DescriptionContains(v string) predicate.User {
	return predicate.User(sql.FieldContains(FieldDescription, v))
}

// DescriptionHasPrefix applies the HasPrefix predicate on the "description" field.
func DescriptionHasPrefix(v string) predicate.User {
	return predicate.User(sql.FieldHasPrefix(FieldDescription, v))
}

// DescriptionHasSuffix applies the HasSuffix predicate on the "description" field.
func DescriptionHasSuffix(v string) predicate.User {
	return predicate.User(sql.FieldHasSuffix(FieldDescription, v))
}

// DescriptionEqualFold applies the EqualFold predicate on the "description" field.
func DescriptionEqualFold(v string) predicate.User {
	return predicate.User(sql.FieldEqualFold(FieldDescription, v))
}

// DescriptionContainsFold applies the ContainsFold predicate on the "description" field.
func DescriptionContainsFold(v string) predicate.User {
	return predicate.User(sql.FieldContainsFold(FieldDescription, v))
}

// CreatedDateEQ applies the EQ predicate on the "created_date" field.
func CreatedDateEQ(v time.Time) predicate.User {
	return predicate.User(sql.FieldEQ(FieldCreatedDate, v))
}

// CreatedDateNEQ applies the NEQ predicate on the "created_date" field.
func CreatedDateNEQ(v time.Time) predicate.User {
	return predicate.User(sql.FieldNEQ(FieldCreatedDate, v))
}

// CreatedDateIn applies the In predicate on the "created_date" field.
func CreatedDateIn(vs ...time.Time) predicate.User {
	return predicate.User(sql.FieldIn(FieldCreatedDate, vs...))
}

// CreatedDateNotIn applies the NotIn predicate on the "created_date" field.
func CreatedDateNotIn(vs ...time.Time) predicate.User {
	return predicate.User(sql.FieldNotIn(FieldCreatedDate, vs...))
}

// CreatedDateGT applies the GT predicate on the "created_date" field.
func CreatedDateGT(v time.Time) predicate.User {
	return predicate.User(sql.FieldGT(FieldCreatedDate, v))
}

// CreatedDateGTE applies the GTE predicate on the "created_date" field.
func CreatedDateGTE(v time.Time) predicate.User {
	return predicate.User(sql.FieldGTE(FieldCreatedDate, v))
}

// CreatedDateLT applies the LT predicate on the "created_date" field.
func CreatedDateLT(v time.Time) predicate.User {
	return predicate.User(sql.FieldLT(FieldCreatedDate, v))
}

// CreatedDateLTE applies the LTE predicate on the "created_date" field.
func CreatedDateLTE(v time.Time) predicate.User {
	return predicate.User(sql.FieldLTE(FieldCreatedDate, v))
}

// HasCaserver applies the HasEdge predicate on the "caserver" edge.
func HasCaserver() predicate.User {
	return predicate.User(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.Edge(sqlgraph.O2M, false, CaserverTable, CaserverColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasCaserverWith applies the HasEdge predicate on the "caserver" edge with a given conditions (other predicates).
func HasCaserverWith(preds ...predicate.CaServer) predicate.User {
	return predicate.User(func(s *sql.Selector) {
		step := newCaserverStep()
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// HasBoards applies the HasEdge predicate on the "boards" edge.
func HasBoards() predicate.User {
	return predicate.User(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.Edge(sqlgraph.O2M, false, BoardsTable, BoardsColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasBoardsWith applies the HasEdge predicate on the "boards" edge with a given conditions (other predicates).
func HasBoardsWith(preds ...predicate.Board) predicate.User {
	return predicate.User(func(s *sql.Selector) {
		step := newBoardsStep()
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// HasPayment applies the HasEdge predicate on the "payment" edge.
func HasPayment() predicate.User {
	return predicate.User(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.Edge(sqlgraph.O2M, false, PaymentTable, PaymentColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasPaymentWith applies the HasEdge predicate on the "payment" edge with a given conditions (other predicates).
func HasPaymentWith(preds ...predicate.Payment) predicate.User {
	return predicate.User(func(s *sql.Selector) {
		step := newPaymentStep()
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// And groups predicates with the AND operator between them.
func And(predicates ...predicate.User) predicate.User {
	return predicate.User(sql.AndPredicates(predicates...))
}

// Or groups predicates with the OR operator between them.
func Or(predicates ...predicate.User) predicate.User {
	return predicate.User(sql.OrPredicates(predicates...))
}

// Not applies the not operator on the given predicate.
func Not(p predicate.User) predicate.User {
	return predicate.User(sql.NotPredicates(p))
}
