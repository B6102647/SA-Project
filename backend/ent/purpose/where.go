// Code generated by entc, DO NOT EDIT.

package purpose

import (
	"github.com/B6102647/app/ent/predicate"
	"github.com/facebookincubator/ent/dialect/sql"
	"github.com/facebookincubator/ent/dialect/sql/sqlgraph"
)

// ID filters vertices based on their identifier.
func ID(id int) predicate.Purpose {
	return predicate.Purpose(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldID), id))
	})
}

// IDEQ applies the EQ predicate on the ID field.
func IDEQ(id int) predicate.Purpose {
	return predicate.Purpose(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldID), id))
	})
}

// IDNEQ applies the NEQ predicate on the ID field.
func IDNEQ(id int) predicate.Purpose {
	return predicate.Purpose(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldID), id))
	})
}

// IDIn applies the In predicate on the ID field.
func IDIn(ids ...int) predicate.Purpose {
	return predicate.Purpose(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(ids) == 0 {
			s.Where(sql.False())
			return
		}
		v := make([]interface{}, len(ids))
		for i := range v {
			v[i] = ids[i]
		}
		s.Where(sql.In(s.C(FieldID), v...))
	})
}

// IDNotIn applies the NotIn predicate on the ID field.
func IDNotIn(ids ...int) predicate.Purpose {
	return predicate.Purpose(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(ids) == 0 {
			s.Where(sql.False())
			return
		}
		v := make([]interface{}, len(ids))
		for i := range v {
			v[i] = ids[i]
		}
		s.Where(sql.NotIn(s.C(FieldID), v...))
	})
}

// IDGT applies the GT predicate on the ID field.
func IDGT(id int) predicate.Purpose {
	return predicate.Purpose(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldID), id))
	})
}

// IDGTE applies the GTE predicate on the ID field.
func IDGTE(id int) predicate.Purpose {
	return predicate.Purpose(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldID), id))
	})
}

// IDLT applies the LT predicate on the ID field.
func IDLT(id int) predicate.Purpose {
	return predicate.Purpose(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldID), id))
	})
}

// IDLTE applies the LTE predicate on the ID field.
func IDLTE(id int) predicate.Purpose {
	return predicate.Purpose(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldID), id))
	})
}

// PURPOSENAME applies equality check predicate on the "PURPOSE_NAME" field. It's identical to PURPOSENAMEEQ.
func PURPOSENAME(v string) predicate.Purpose {
	return predicate.Purpose(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldPURPOSENAME), v))
	})
}

// PURPOSENAMEEQ applies the EQ predicate on the "PURPOSE_NAME" field.
func PURPOSENAMEEQ(v string) predicate.Purpose {
	return predicate.Purpose(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldPURPOSENAME), v))
	})
}

// PURPOSENAMENEQ applies the NEQ predicate on the "PURPOSE_NAME" field.
func PURPOSENAMENEQ(v string) predicate.Purpose {
	return predicate.Purpose(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldPURPOSENAME), v))
	})
}

// PURPOSENAMEIn applies the In predicate on the "PURPOSE_NAME" field.
func PURPOSENAMEIn(vs ...string) predicate.Purpose {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Purpose(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.In(s.C(FieldPURPOSENAME), v...))
	})
}

// PURPOSENAMENotIn applies the NotIn predicate on the "PURPOSE_NAME" field.
func PURPOSENAMENotIn(vs ...string) predicate.Purpose {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Purpose(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.NotIn(s.C(FieldPURPOSENAME), v...))
	})
}

// PURPOSENAMEGT applies the GT predicate on the "PURPOSE_NAME" field.
func PURPOSENAMEGT(v string) predicate.Purpose {
	return predicate.Purpose(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldPURPOSENAME), v))
	})
}

// PURPOSENAMEGTE applies the GTE predicate on the "PURPOSE_NAME" field.
func PURPOSENAMEGTE(v string) predicate.Purpose {
	return predicate.Purpose(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldPURPOSENAME), v))
	})
}

// PURPOSENAMELT applies the LT predicate on the "PURPOSE_NAME" field.
func PURPOSENAMELT(v string) predicate.Purpose {
	return predicate.Purpose(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldPURPOSENAME), v))
	})
}

// PURPOSENAMELTE applies the LTE predicate on the "PURPOSE_NAME" field.
func PURPOSENAMELTE(v string) predicate.Purpose {
	return predicate.Purpose(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldPURPOSENAME), v))
	})
}

// PURPOSENAMEContains applies the Contains predicate on the "PURPOSE_NAME" field.
func PURPOSENAMEContains(v string) predicate.Purpose {
	return predicate.Purpose(func(s *sql.Selector) {
		s.Where(sql.Contains(s.C(FieldPURPOSENAME), v))
	})
}

// PURPOSENAMEHasPrefix applies the HasPrefix predicate on the "PURPOSE_NAME" field.
func PURPOSENAMEHasPrefix(v string) predicate.Purpose {
	return predicate.Purpose(func(s *sql.Selector) {
		s.Where(sql.HasPrefix(s.C(FieldPURPOSENAME), v))
	})
}

// PURPOSENAMEHasSuffix applies the HasSuffix predicate on the "PURPOSE_NAME" field.
func PURPOSENAMEHasSuffix(v string) predicate.Purpose {
	return predicate.Purpose(func(s *sql.Selector) {
		s.Where(sql.HasSuffix(s.C(FieldPURPOSENAME), v))
	})
}

// PURPOSENAMEEqualFold applies the EqualFold predicate on the "PURPOSE_NAME" field.
func PURPOSENAMEEqualFold(v string) predicate.Purpose {
	return predicate.Purpose(func(s *sql.Selector) {
		s.Where(sql.EqualFold(s.C(FieldPURPOSENAME), v))
	})
}

// PURPOSENAMEContainsFold applies the ContainsFold predicate on the "PURPOSE_NAME" field.
func PURPOSENAMEContainsFold(v string) predicate.Purpose {
	return predicate.Purpose(func(s *sql.Selector) {
		s.Where(sql.ContainsFold(s.C(FieldPURPOSENAME), v))
	})
}

// HasBooklist applies the HasEdge predicate on the "Booklist" edge.
func HasBooklist() predicate.Purpose {
	return predicate.Purpose(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.To(BooklistTable, FieldID),
			sqlgraph.Edge(sqlgraph.O2M, false, BooklistTable, BooklistColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasBooklistWith applies the HasEdge predicate on the "Booklist" edge with a given conditions (other predicates).
func HasBooklistWith(preds ...predicate.BookBorrow) predicate.Purpose {
	return predicate.Purpose(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.To(BooklistInverseTable, FieldID),
			sqlgraph.Edge(sqlgraph.O2M, false, BooklistTable, BooklistColumn),
		)
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// And groups list of predicates with the AND operator between them.
func And(predicates ...predicate.Purpose) predicate.Purpose {
	return predicate.Purpose(func(s *sql.Selector) {
		s1 := s.Clone().SetP(nil)
		for _, p := range predicates {
			p(s1)
		}
		s.Where(s1.P())
	})
}

// Or groups list of predicates with the OR operator between them.
func Or(predicates ...predicate.Purpose) predicate.Purpose {
	return predicate.Purpose(func(s *sql.Selector) {
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
func Not(p predicate.Purpose) predicate.Purpose {
	return predicate.Purpose(func(s *sql.Selector) {
		p(s.Not())
	})
}
