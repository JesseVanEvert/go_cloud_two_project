// Code generated by ent, DO NOT EDIT.

package classlecturer

import (
	"lecturer/ent/predicate"
	"time"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
)

// ID filters vertices based on their ID field.
func ID(id int) predicate.ClassLecturer {
	return predicate.ClassLecturer(sql.FieldEQ(FieldID, id))
}

// IDEQ applies the EQ predicate on the ID field.
func IDEQ(id int) predicate.ClassLecturer {
	return predicate.ClassLecturer(sql.FieldEQ(FieldID, id))
}

// IDNEQ applies the NEQ predicate on the ID field.
func IDNEQ(id int) predicate.ClassLecturer {
	return predicate.ClassLecturer(sql.FieldNEQ(FieldID, id))
}

// IDIn applies the In predicate on the ID field.
func IDIn(ids ...int) predicate.ClassLecturer {
	return predicate.ClassLecturer(sql.FieldIn(FieldID, ids...))
}

// IDNotIn applies the NotIn predicate on the ID field.
func IDNotIn(ids ...int) predicate.ClassLecturer {
	return predicate.ClassLecturer(sql.FieldNotIn(FieldID, ids...))
}

// IDGT applies the GT predicate on the ID field.
func IDGT(id int) predicate.ClassLecturer {
	return predicate.ClassLecturer(sql.FieldGT(FieldID, id))
}

// IDGTE applies the GTE predicate on the ID field.
func IDGTE(id int) predicate.ClassLecturer {
	return predicate.ClassLecturer(sql.FieldGTE(FieldID, id))
}

// IDLT applies the LT predicate on the ID field.
func IDLT(id int) predicate.ClassLecturer {
	return predicate.ClassLecturer(sql.FieldLT(FieldID, id))
}

// IDLTE applies the LTE predicate on the ID field.
func IDLTE(id int) predicate.ClassLecturer {
	return predicate.ClassLecturer(sql.FieldLTE(FieldID, id))
}

// DeletedAt applies equality check predicate on the "deleted_at" field. It's identical to DeletedAtEQ.
func DeletedAt(v time.Time) predicate.ClassLecturer {
	return predicate.ClassLecturer(sql.FieldEQ(FieldDeletedAt, v))
}

// DeletedAtEQ applies the EQ predicate on the "deleted_at" field.
func DeletedAtEQ(v time.Time) predicate.ClassLecturer {
	return predicate.ClassLecturer(sql.FieldEQ(FieldDeletedAt, v))
}

// DeletedAtNEQ applies the NEQ predicate on the "deleted_at" field.
func DeletedAtNEQ(v time.Time) predicate.ClassLecturer {
	return predicate.ClassLecturer(sql.FieldNEQ(FieldDeletedAt, v))
}

// DeletedAtIn applies the In predicate on the "deleted_at" field.
func DeletedAtIn(vs ...time.Time) predicate.ClassLecturer {
	return predicate.ClassLecturer(sql.FieldIn(FieldDeletedAt, vs...))
}

// DeletedAtNotIn applies the NotIn predicate on the "deleted_at" field.
func DeletedAtNotIn(vs ...time.Time) predicate.ClassLecturer {
	return predicate.ClassLecturer(sql.FieldNotIn(FieldDeletedAt, vs...))
}

// DeletedAtGT applies the GT predicate on the "deleted_at" field.
func DeletedAtGT(v time.Time) predicate.ClassLecturer {
	return predicate.ClassLecturer(sql.FieldGT(FieldDeletedAt, v))
}

// DeletedAtGTE applies the GTE predicate on the "deleted_at" field.
func DeletedAtGTE(v time.Time) predicate.ClassLecturer {
	return predicate.ClassLecturer(sql.FieldGTE(FieldDeletedAt, v))
}

// DeletedAtLT applies the LT predicate on the "deleted_at" field.
func DeletedAtLT(v time.Time) predicate.ClassLecturer {
	return predicate.ClassLecturer(sql.FieldLT(FieldDeletedAt, v))
}

// DeletedAtLTE applies the LTE predicate on the "deleted_at" field.
func DeletedAtLTE(v time.Time) predicate.ClassLecturer {
	return predicate.ClassLecturer(sql.FieldLTE(FieldDeletedAt, v))
}

// HasClass applies the HasEdge predicate on the "class" edge.
func HasClass() predicate.ClassLecturer {
	return predicate.ClassLecturer(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, ClassTable, ClassColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasClassWith applies the HasEdge predicate on the "class" edge with a given conditions (other predicates).
func HasClassWith(preds ...predicate.Class) predicate.ClassLecturer {
	return predicate.ClassLecturer(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.To(ClassInverseTable, FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, ClassTable, ClassColumn),
		)
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// HasLecturer applies the HasEdge predicate on the "lecturer" edge.
func HasLecturer() predicate.ClassLecturer {
	return predicate.ClassLecturer(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, LecturerTable, LecturerColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasLecturerWith applies the HasEdge predicate on the "lecturer" edge with a given conditions (other predicates).
func HasLecturerWith(preds ...predicate.Lecturer) predicate.ClassLecturer {
	return predicate.ClassLecturer(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.To(LecturerInverseTable, FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, LecturerTable, LecturerColumn),
		)
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// And groups predicates with the AND operator between them.
func And(predicates ...predicate.ClassLecturer) predicate.ClassLecturer {
	return predicate.ClassLecturer(func(s *sql.Selector) {
		s1 := s.Clone().SetP(nil)
		for _, p := range predicates {
			p(s1)
		}
		s.Where(s1.P())
	})
}

// Or groups predicates with the OR operator between them.
func Or(predicates ...predicate.ClassLecturer) predicate.ClassLecturer {
	return predicate.ClassLecturer(func(s *sql.Selector) {
		s1 := s.Clone().SetP(nil)
		for i, p := range predicates {
			if i > 0 {
				s1.Or()
			}
			p(s1)
		}
		s.Where(s1.P())
	})
}

// Not applies the not operator on the given predicate.
func Not(p predicate.ClassLecturer) predicate.ClassLecturer {
	return predicate.ClassLecturer(func(s *sql.Selector) {
		p(s.Not())
	})
}