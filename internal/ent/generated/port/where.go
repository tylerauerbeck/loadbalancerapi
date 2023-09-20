// Copyright 2023 The Infratographer Authors
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
// 	http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.
//
// Code generated by entc, DO NOT EDIT.

package port

import (
	"time"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"go.infratographer.com/load-balancer-api/internal/ent/generated/predicate"
	"go.infratographer.com/x/gidx"
)

// ID filters vertices based on their ID field.
func ID(id gidx.PrefixedID) predicate.Port {
	return predicate.Port(sql.FieldEQ(FieldID, id))
}

// IDEQ applies the EQ predicate on the ID field.
func IDEQ(id gidx.PrefixedID) predicate.Port {
	return predicate.Port(sql.FieldEQ(FieldID, id))
}

// IDNEQ applies the NEQ predicate on the ID field.
func IDNEQ(id gidx.PrefixedID) predicate.Port {
	return predicate.Port(sql.FieldNEQ(FieldID, id))
}

// IDIn applies the In predicate on the ID field.
func IDIn(ids ...gidx.PrefixedID) predicate.Port {
	return predicate.Port(sql.FieldIn(FieldID, ids...))
}

// IDNotIn applies the NotIn predicate on the ID field.
func IDNotIn(ids ...gidx.PrefixedID) predicate.Port {
	return predicate.Port(sql.FieldNotIn(FieldID, ids...))
}

// IDGT applies the GT predicate on the ID field.
func IDGT(id gidx.PrefixedID) predicate.Port {
	return predicate.Port(sql.FieldGT(FieldID, id))
}

// IDGTE applies the GTE predicate on the ID field.
func IDGTE(id gidx.PrefixedID) predicate.Port {
	return predicate.Port(sql.FieldGTE(FieldID, id))
}

// IDLT applies the LT predicate on the ID field.
func IDLT(id gidx.PrefixedID) predicate.Port {
	return predicate.Port(sql.FieldLT(FieldID, id))
}

// IDLTE applies the LTE predicate on the ID field.
func IDLTE(id gidx.PrefixedID) predicate.Port {
	return predicate.Port(sql.FieldLTE(FieldID, id))
}

// CreatedAt applies equality check predicate on the "created_at" field. It's identical to CreatedAtEQ.
func CreatedAt(v time.Time) predicate.Port {
	return predicate.Port(sql.FieldEQ(FieldCreatedAt, v))
}

// UpdatedAt applies equality check predicate on the "updated_at" field. It's identical to UpdatedAtEQ.
func UpdatedAt(v time.Time) predicate.Port {
	return predicate.Port(sql.FieldEQ(FieldUpdatedAt, v))
}

// DeletedAt applies equality check predicate on the "deleted_at" field. It's identical to DeletedAtEQ.
func DeletedAt(v time.Time) predicate.Port {
	return predicate.Port(sql.FieldEQ(FieldDeletedAt, v))
}

// Number applies equality check predicate on the "number" field. It's identical to NumberEQ.
func Number(v int) predicate.Port {
	return predicate.Port(sql.FieldEQ(FieldNumber, v))
}

// Name applies equality check predicate on the "name" field. It's identical to NameEQ.
func Name(v string) predicate.Port {
	return predicate.Port(sql.FieldEQ(FieldName, v))
}

// LoadBalancerID applies equality check predicate on the "load_balancer_id" field. It's identical to LoadBalancerIDEQ.
func LoadBalancerID(v gidx.PrefixedID) predicate.Port {
	return predicate.Port(sql.FieldEQ(FieldLoadBalancerID, v))
}

// CreatedAtEQ applies the EQ predicate on the "created_at" field.
func CreatedAtEQ(v time.Time) predicate.Port {
	return predicate.Port(sql.FieldEQ(FieldCreatedAt, v))
}

// CreatedAtNEQ applies the NEQ predicate on the "created_at" field.
func CreatedAtNEQ(v time.Time) predicate.Port {
	return predicate.Port(sql.FieldNEQ(FieldCreatedAt, v))
}

// CreatedAtIn applies the In predicate on the "created_at" field.
func CreatedAtIn(vs ...time.Time) predicate.Port {
	return predicate.Port(sql.FieldIn(FieldCreatedAt, vs...))
}

// CreatedAtNotIn applies the NotIn predicate on the "created_at" field.
func CreatedAtNotIn(vs ...time.Time) predicate.Port {
	return predicate.Port(sql.FieldNotIn(FieldCreatedAt, vs...))
}

// CreatedAtGT applies the GT predicate on the "created_at" field.
func CreatedAtGT(v time.Time) predicate.Port {
	return predicate.Port(sql.FieldGT(FieldCreatedAt, v))
}

// CreatedAtGTE applies the GTE predicate on the "created_at" field.
func CreatedAtGTE(v time.Time) predicate.Port {
	return predicate.Port(sql.FieldGTE(FieldCreatedAt, v))
}

// CreatedAtLT applies the LT predicate on the "created_at" field.
func CreatedAtLT(v time.Time) predicate.Port {
	return predicate.Port(sql.FieldLT(FieldCreatedAt, v))
}

// CreatedAtLTE applies the LTE predicate on the "created_at" field.
func CreatedAtLTE(v time.Time) predicate.Port {
	return predicate.Port(sql.FieldLTE(FieldCreatedAt, v))
}

// UpdatedAtEQ applies the EQ predicate on the "updated_at" field.
func UpdatedAtEQ(v time.Time) predicate.Port {
	return predicate.Port(sql.FieldEQ(FieldUpdatedAt, v))
}

// UpdatedAtNEQ applies the NEQ predicate on the "updated_at" field.
func UpdatedAtNEQ(v time.Time) predicate.Port {
	return predicate.Port(sql.FieldNEQ(FieldUpdatedAt, v))
}

// UpdatedAtIn applies the In predicate on the "updated_at" field.
func UpdatedAtIn(vs ...time.Time) predicate.Port {
	return predicate.Port(sql.FieldIn(FieldUpdatedAt, vs...))
}

// UpdatedAtNotIn applies the NotIn predicate on the "updated_at" field.
func UpdatedAtNotIn(vs ...time.Time) predicate.Port {
	return predicate.Port(sql.FieldNotIn(FieldUpdatedAt, vs...))
}

// UpdatedAtGT applies the GT predicate on the "updated_at" field.
func UpdatedAtGT(v time.Time) predicate.Port {
	return predicate.Port(sql.FieldGT(FieldUpdatedAt, v))
}

// UpdatedAtGTE applies the GTE predicate on the "updated_at" field.
func UpdatedAtGTE(v time.Time) predicate.Port {
	return predicate.Port(sql.FieldGTE(FieldUpdatedAt, v))
}

// UpdatedAtLT applies the LT predicate on the "updated_at" field.
func UpdatedAtLT(v time.Time) predicate.Port {
	return predicate.Port(sql.FieldLT(FieldUpdatedAt, v))
}

// UpdatedAtLTE applies the LTE predicate on the "updated_at" field.
func UpdatedAtLTE(v time.Time) predicate.Port {
	return predicate.Port(sql.FieldLTE(FieldUpdatedAt, v))
}

// DeletedAtEQ applies the EQ predicate on the "deleted_at" field.
func DeletedAtEQ(v time.Time) predicate.Port {
	return predicate.Port(sql.FieldEQ(FieldDeletedAt, v))
}

// DeletedAtNEQ applies the NEQ predicate on the "deleted_at" field.
func DeletedAtNEQ(v time.Time) predicate.Port {
	return predicate.Port(sql.FieldNEQ(FieldDeletedAt, v))
}

// DeletedAtIn applies the In predicate on the "deleted_at" field.
func DeletedAtIn(vs ...time.Time) predicate.Port {
	return predicate.Port(sql.FieldIn(FieldDeletedAt, vs...))
}

// DeletedAtNotIn applies the NotIn predicate on the "deleted_at" field.
func DeletedAtNotIn(vs ...time.Time) predicate.Port {
	return predicate.Port(sql.FieldNotIn(FieldDeletedAt, vs...))
}

// DeletedAtGT applies the GT predicate on the "deleted_at" field.
func DeletedAtGT(v time.Time) predicate.Port {
	return predicate.Port(sql.FieldGT(FieldDeletedAt, v))
}

// DeletedAtGTE applies the GTE predicate on the "deleted_at" field.
func DeletedAtGTE(v time.Time) predicate.Port {
	return predicate.Port(sql.FieldGTE(FieldDeletedAt, v))
}

// DeletedAtLT applies the LT predicate on the "deleted_at" field.
func DeletedAtLT(v time.Time) predicate.Port {
	return predicate.Port(sql.FieldLT(FieldDeletedAt, v))
}

// DeletedAtLTE applies the LTE predicate on the "deleted_at" field.
func DeletedAtLTE(v time.Time) predicate.Port {
	return predicate.Port(sql.FieldLTE(FieldDeletedAt, v))
}

// DeletedAtIsNil applies the IsNil predicate on the "deleted_at" field.
func DeletedAtIsNil() predicate.Port {
	return predicate.Port(sql.FieldIsNull(FieldDeletedAt))
}

// DeletedAtNotNil applies the NotNil predicate on the "deleted_at" field.
func DeletedAtNotNil() predicate.Port {
	return predicate.Port(sql.FieldNotNull(FieldDeletedAt))
}

// NumberEQ applies the EQ predicate on the "number" field.
func NumberEQ(v int) predicate.Port {
	return predicate.Port(sql.FieldEQ(FieldNumber, v))
}

// NumberNEQ applies the NEQ predicate on the "number" field.
func NumberNEQ(v int) predicate.Port {
	return predicate.Port(sql.FieldNEQ(FieldNumber, v))
}

// NumberIn applies the In predicate on the "number" field.
func NumberIn(vs ...int) predicate.Port {
	return predicate.Port(sql.FieldIn(FieldNumber, vs...))
}

// NumberNotIn applies the NotIn predicate on the "number" field.
func NumberNotIn(vs ...int) predicate.Port {
	return predicate.Port(sql.FieldNotIn(FieldNumber, vs...))
}

// NumberGT applies the GT predicate on the "number" field.
func NumberGT(v int) predicate.Port {
	return predicate.Port(sql.FieldGT(FieldNumber, v))
}

// NumberGTE applies the GTE predicate on the "number" field.
func NumberGTE(v int) predicate.Port {
	return predicate.Port(sql.FieldGTE(FieldNumber, v))
}

// NumberLT applies the LT predicate on the "number" field.
func NumberLT(v int) predicate.Port {
	return predicate.Port(sql.FieldLT(FieldNumber, v))
}

// NumberLTE applies the LTE predicate on the "number" field.
func NumberLTE(v int) predicate.Port {
	return predicate.Port(sql.FieldLTE(FieldNumber, v))
}

// NameEQ applies the EQ predicate on the "name" field.
func NameEQ(v string) predicate.Port {
	return predicate.Port(sql.FieldEQ(FieldName, v))
}

// NameNEQ applies the NEQ predicate on the "name" field.
func NameNEQ(v string) predicate.Port {
	return predicate.Port(sql.FieldNEQ(FieldName, v))
}

// NameIn applies the In predicate on the "name" field.
func NameIn(vs ...string) predicate.Port {
	return predicate.Port(sql.FieldIn(FieldName, vs...))
}

// NameNotIn applies the NotIn predicate on the "name" field.
func NameNotIn(vs ...string) predicate.Port {
	return predicate.Port(sql.FieldNotIn(FieldName, vs...))
}

// NameGT applies the GT predicate on the "name" field.
func NameGT(v string) predicate.Port {
	return predicate.Port(sql.FieldGT(FieldName, v))
}

// NameGTE applies the GTE predicate on the "name" field.
func NameGTE(v string) predicate.Port {
	return predicate.Port(sql.FieldGTE(FieldName, v))
}

// NameLT applies the LT predicate on the "name" field.
func NameLT(v string) predicate.Port {
	return predicate.Port(sql.FieldLT(FieldName, v))
}

// NameLTE applies the LTE predicate on the "name" field.
func NameLTE(v string) predicate.Port {
	return predicate.Port(sql.FieldLTE(FieldName, v))
}

// NameContains applies the Contains predicate on the "name" field.
func NameContains(v string) predicate.Port {
	return predicate.Port(sql.FieldContains(FieldName, v))
}

// NameHasPrefix applies the HasPrefix predicate on the "name" field.
func NameHasPrefix(v string) predicate.Port {
	return predicate.Port(sql.FieldHasPrefix(FieldName, v))
}

// NameHasSuffix applies the HasSuffix predicate on the "name" field.
func NameHasSuffix(v string) predicate.Port {
	return predicate.Port(sql.FieldHasSuffix(FieldName, v))
}

// NameEqualFold applies the EqualFold predicate on the "name" field.
func NameEqualFold(v string) predicate.Port {
	return predicate.Port(sql.FieldEqualFold(FieldName, v))
}

// NameContainsFold applies the ContainsFold predicate on the "name" field.
func NameContainsFold(v string) predicate.Port {
	return predicate.Port(sql.FieldContainsFold(FieldName, v))
}

// LoadBalancerIDEQ applies the EQ predicate on the "load_balancer_id" field.
func LoadBalancerIDEQ(v gidx.PrefixedID) predicate.Port {
	return predicate.Port(sql.FieldEQ(FieldLoadBalancerID, v))
}

// LoadBalancerIDNEQ applies the NEQ predicate on the "load_balancer_id" field.
func LoadBalancerIDNEQ(v gidx.PrefixedID) predicate.Port {
	return predicate.Port(sql.FieldNEQ(FieldLoadBalancerID, v))
}

// LoadBalancerIDIn applies the In predicate on the "load_balancer_id" field.
func LoadBalancerIDIn(vs ...gidx.PrefixedID) predicate.Port {
	return predicate.Port(sql.FieldIn(FieldLoadBalancerID, vs...))
}

// LoadBalancerIDNotIn applies the NotIn predicate on the "load_balancer_id" field.
func LoadBalancerIDNotIn(vs ...gidx.PrefixedID) predicate.Port {
	return predicate.Port(sql.FieldNotIn(FieldLoadBalancerID, vs...))
}

// LoadBalancerIDGT applies the GT predicate on the "load_balancer_id" field.
func LoadBalancerIDGT(v gidx.PrefixedID) predicate.Port {
	return predicate.Port(sql.FieldGT(FieldLoadBalancerID, v))
}

// LoadBalancerIDGTE applies the GTE predicate on the "load_balancer_id" field.
func LoadBalancerIDGTE(v gidx.PrefixedID) predicate.Port {
	return predicate.Port(sql.FieldGTE(FieldLoadBalancerID, v))
}

// LoadBalancerIDLT applies the LT predicate on the "load_balancer_id" field.
func LoadBalancerIDLT(v gidx.PrefixedID) predicate.Port {
	return predicate.Port(sql.FieldLT(FieldLoadBalancerID, v))
}

// LoadBalancerIDLTE applies the LTE predicate on the "load_balancer_id" field.
func LoadBalancerIDLTE(v gidx.PrefixedID) predicate.Port {
	return predicate.Port(sql.FieldLTE(FieldLoadBalancerID, v))
}

// LoadBalancerIDContains applies the Contains predicate on the "load_balancer_id" field.
func LoadBalancerIDContains(v gidx.PrefixedID) predicate.Port {
	vc := string(v)
	return predicate.Port(sql.FieldContains(FieldLoadBalancerID, vc))
}

// LoadBalancerIDHasPrefix applies the HasPrefix predicate on the "load_balancer_id" field.
func LoadBalancerIDHasPrefix(v gidx.PrefixedID) predicate.Port {
	vc := string(v)
	return predicate.Port(sql.FieldHasPrefix(FieldLoadBalancerID, vc))
}

// LoadBalancerIDHasSuffix applies the HasSuffix predicate on the "load_balancer_id" field.
func LoadBalancerIDHasSuffix(v gidx.PrefixedID) predicate.Port {
	vc := string(v)
	return predicate.Port(sql.FieldHasSuffix(FieldLoadBalancerID, vc))
}

// LoadBalancerIDEqualFold applies the EqualFold predicate on the "load_balancer_id" field.
func LoadBalancerIDEqualFold(v gidx.PrefixedID) predicate.Port {
	vc := string(v)
	return predicate.Port(sql.FieldEqualFold(FieldLoadBalancerID, vc))
}

// LoadBalancerIDContainsFold applies the ContainsFold predicate on the "load_balancer_id" field.
func LoadBalancerIDContainsFold(v gidx.PrefixedID) predicate.Port {
	vc := string(v)
	return predicate.Port(sql.FieldContainsFold(FieldLoadBalancerID, vc))
}

// HasPools applies the HasEdge predicate on the "pools" edge.
func HasPools() predicate.Port {
	return predicate.Port(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.Edge(sqlgraph.M2M, true, PoolsTable, PoolsPrimaryKey...),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasPoolsWith applies the HasEdge predicate on the "pools" edge with a given conditions (other predicates).
func HasPoolsWith(preds ...predicate.Pool) predicate.Port {
	return predicate.Port(func(s *sql.Selector) {
		step := newPoolsStep()
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// HasLoadBalancer applies the HasEdge predicate on the "load_balancer" edge.
func HasLoadBalancer() predicate.Port {
	return predicate.Port(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.Edge(sqlgraph.M2O, false, LoadBalancerTable, LoadBalancerColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasLoadBalancerWith applies the HasEdge predicate on the "load_balancer" edge with a given conditions (other predicates).
func HasLoadBalancerWith(preds ...predicate.LoadBalancer) predicate.Port {
	return predicate.Port(func(s *sql.Selector) {
		step := newLoadBalancerStep()
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// And groups predicates with the AND operator between them.
func And(predicates ...predicate.Port) predicate.Port {
	return predicate.Port(func(s *sql.Selector) {
		s1 := s.Clone().SetP(nil)
		for _, p := range predicates {
			p(s1)
		}
		s.Where(s1.P())
	})
}

// Or groups predicates with the OR operator between them.
func Or(predicates ...predicate.Port) predicate.Port {
	return predicate.Port(func(s *sql.Selector) {
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
func Not(p predicate.Port) predicate.Port {
	return predicate.Port(func(s *sql.Selector) {
		p(s.Not())
	})
}
