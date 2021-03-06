// Code generated by entc, DO NOT EDIT.

package channel

import (
	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"github.com/lemon-mint/notification-bot/ent/predicate"
)

// ID filters vertices based on their ID field.
func ID(id int) predicate.Channel {
	return predicate.Channel(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldID), id))
	})
}

// IDEQ applies the EQ predicate on the ID field.
func IDEQ(id int) predicate.Channel {
	return predicate.Channel(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldID), id))
	})
}

// IDNEQ applies the NEQ predicate on the ID field.
func IDNEQ(id int) predicate.Channel {
	return predicate.Channel(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldID), id))
	})
}

// IDIn applies the In predicate on the ID field.
func IDIn(ids ...int) predicate.Channel {
	return predicate.Channel(func(s *sql.Selector) {
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
func IDNotIn(ids ...int) predicate.Channel {
	return predicate.Channel(func(s *sql.Selector) {
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
func IDGT(id int) predicate.Channel {
	return predicate.Channel(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldID), id))
	})
}

// IDGTE applies the GTE predicate on the ID field.
func IDGTE(id int) predicate.Channel {
	return predicate.Channel(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldID), id))
	})
}

// IDLT applies the LT predicate on the ID field.
func IDLT(id int) predicate.Channel {
	return predicate.Channel(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldID), id))
	})
}

// IDLTE applies the LTE predicate on the ID field.
func IDLTE(id int) predicate.Channel {
	return predicate.Channel(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldID), id))
	})
}

// ChannelID applies equality check predicate on the "ChannelID" field. It's identical to ChannelIDEQ.
func ChannelID(v string) predicate.Channel {
	return predicate.Channel(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldChannelID), v))
	})
}

// Name applies equality check predicate on the "Name" field. It's identical to NameEQ.
func Name(v string) predicate.Channel {
	return predicate.Channel(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldName), v))
	})
}

// ChannelIDEQ applies the EQ predicate on the "ChannelID" field.
func ChannelIDEQ(v string) predicate.Channel {
	return predicate.Channel(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldChannelID), v))
	})
}

// ChannelIDNEQ applies the NEQ predicate on the "ChannelID" field.
func ChannelIDNEQ(v string) predicate.Channel {
	return predicate.Channel(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldChannelID), v))
	})
}

// ChannelIDIn applies the In predicate on the "ChannelID" field.
func ChannelIDIn(vs ...string) predicate.Channel {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Channel(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.In(s.C(FieldChannelID), v...))
	})
}

// ChannelIDNotIn applies the NotIn predicate on the "ChannelID" field.
func ChannelIDNotIn(vs ...string) predicate.Channel {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Channel(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.NotIn(s.C(FieldChannelID), v...))
	})
}

// ChannelIDGT applies the GT predicate on the "ChannelID" field.
func ChannelIDGT(v string) predicate.Channel {
	return predicate.Channel(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldChannelID), v))
	})
}

// ChannelIDGTE applies the GTE predicate on the "ChannelID" field.
func ChannelIDGTE(v string) predicate.Channel {
	return predicate.Channel(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldChannelID), v))
	})
}

// ChannelIDLT applies the LT predicate on the "ChannelID" field.
func ChannelIDLT(v string) predicate.Channel {
	return predicate.Channel(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldChannelID), v))
	})
}

// ChannelIDLTE applies the LTE predicate on the "ChannelID" field.
func ChannelIDLTE(v string) predicate.Channel {
	return predicate.Channel(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldChannelID), v))
	})
}

// ChannelIDContains applies the Contains predicate on the "ChannelID" field.
func ChannelIDContains(v string) predicate.Channel {
	return predicate.Channel(func(s *sql.Selector) {
		s.Where(sql.Contains(s.C(FieldChannelID), v))
	})
}

// ChannelIDHasPrefix applies the HasPrefix predicate on the "ChannelID" field.
func ChannelIDHasPrefix(v string) predicate.Channel {
	return predicate.Channel(func(s *sql.Selector) {
		s.Where(sql.HasPrefix(s.C(FieldChannelID), v))
	})
}

// ChannelIDHasSuffix applies the HasSuffix predicate on the "ChannelID" field.
func ChannelIDHasSuffix(v string) predicate.Channel {
	return predicate.Channel(func(s *sql.Selector) {
		s.Where(sql.HasSuffix(s.C(FieldChannelID), v))
	})
}

// ChannelIDEqualFold applies the EqualFold predicate on the "ChannelID" field.
func ChannelIDEqualFold(v string) predicate.Channel {
	return predicate.Channel(func(s *sql.Selector) {
		s.Where(sql.EqualFold(s.C(FieldChannelID), v))
	})
}

// ChannelIDContainsFold applies the ContainsFold predicate on the "ChannelID" field.
func ChannelIDContainsFold(v string) predicate.Channel {
	return predicate.Channel(func(s *sql.Selector) {
		s.Where(sql.ContainsFold(s.C(FieldChannelID), v))
	})
}

// NameEQ applies the EQ predicate on the "Name" field.
func NameEQ(v string) predicate.Channel {
	return predicate.Channel(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldName), v))
	})
}

// NameNEQ applies the NEQ predicate on the "Name" field.
func NameNEQ(v string) predicate.Channel {
	return predicate.Channel(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldName), v))
	})
}

// NameIn applies the In predicate on the "Name" field.
func NameIn(vs ...string) predicate.Channel {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Channel(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.In(s.C(FieldName), v...))
	})
}

// NameNotIn applies the NotIn predicate on the "Name" field.
func NameNotIn(vs ...string) predicate.Channel {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Channel(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.NotIn(s.C(FieldName), v...))
	})
}

// NameGT applies the GT predicate on the "Name" field.
func NameGT(v string) predicate.Channel {
	return predicate.Channel(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldName), v))
	})
}

// NameGTE applies the GTE predicate on the "Name" field.
func NameGTE(v string) predicate.Channel {
	return predicate.Channel(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldName), v))
	})
}

// NameLT applies the LT predicate on the "Name" field.
func NameLT(v string) predicate.Channel {
	return predicate.Channel(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldName), v))
	})
}

// NameLTE applies the LTE predicate on the "Name" field.
func NameLTE(v string) predicate.Channel {
	return predicate.Channel(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldName), v))
	})
}

// NameContains applies the Contains predicate on the "Name" field.
func NameContains(v string) predicate.Channel {
	return predicate.Channel(func(s *sql.Selector) {
		s.Where(sql.Contains(s.C(FieldName), v))
	})
}

// NameHasPrefix applies the HasPrefix predicate on the "Name" field.
func NameHasPrefix(v string) predicate.Channel {
	return predicate.Channel(func(s *sql.Selector) {
		s.Where(sql.HasPrefix(s.C(FieldName), v))
	})
}

// NameHasSuffix applies the HasSuffix predicate on the "Name" field.
func NameHasSuffix(v string) predicate.Channel {
	return predicate.Channel(func(s *sql.Selector) {
		s.Where(sql.HasSuffix(s.C(FieldName), v))
	})
}

// NameEqualFold applies the EqualFold predicate on the "Name" field.
func NameEqualFold(v string) predicate.Channel {
	return predicate.Channel(func(s *sql.Selector) {
		s.Where(sql.EqualFold(s.C(FieldName), v))
	})
}

// NameContainsFold applies the ContainsFold predicate on the "Name" field.
func NameContainsFold(v string) predicate.Channel {
	return predicate.Channel(func(s *sql.Selector) {
		s.Where(sql.ContainsFold(s.C(FieldName), v))
	})
}

// HasSubscribers applies the HasEdge predicate on the "Subscribers" edge.
func HasSubscribers() predicate.Channel {
	return predicate.Channel(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.To(SubscribersTable, FieldID),
			sqlgraph.Edge(sqlgraph.M2M, false, SubscribersTable, SubscribersPrimaryKey...),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasSubscribersWith applies the HasEdge predicate on the "Subscribers" edge with a given conditions (other predicates).
func HasSubscribersWith(preds ...predicate.User) predicate.Channel {
	return predicate.Channel(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.To(SubscribersInverseTable, FieldID),
			sqlgraph.Edge(sqlgraph.M2M, false, SubscribersTable, SubscribersPrimaryKey...),
		)
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// And groups predicates with the AND operator between them.
func And(predicates ...predicate.Channel) predicate.Channel {
	return predicate.Channel(func(s *sql.Selector) {
		s1 := s.Clone().SetP(nil)
		for _, p := range predicates {
			p(s1)
		}
		s.Where(s1.P())
	})
}

// Or groups predicates with the OR operator between them.
func Or(predicates ...predicate.Channel) predicate.Channel {
	return predicate.Channel(func(s *sql.Selector) {
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
func Not(p predicate.Channel) predicate.Channel {
	return predicate.Channel(func(s *sql.Selector) {
		p(s.Not())
	})
}
