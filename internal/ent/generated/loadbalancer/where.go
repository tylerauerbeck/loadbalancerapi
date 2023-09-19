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

package loadbalancer

import (
	"time"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"go.infratographer.com/load-balancer-api/internal/ent/generated/predicate"
	"go.infratographer.com/x/gidx"
)

// ID filters vertices based on their ID field.
func ID(id gidx.PrefixedID) predicate.LoadBalancer {
	return predicate.LoadBalancer(sql.FieldEQ(FieldID, id))
}

// IDEQ applies the EQ predicate on the ID field.
func IDEQ(id gidx.PrefixedID) predicate.LoadBalancer {
	return predicate.LoadBalancer(sql.FieldEQ(FieldID, id))
}

// IDNEQ applies the NEQ predicate on the ID field.
func IDNEQ(id gidx.PrefixedID) predicate.LoadBalancer {
	return predicate.LoadBalancer(sql.FieldNEQ(FieldID, id))
}

// IDIn applies the In predicate on the ID field.
func IDIn(ids ...gidx.PrefixedID) predicate.LoadBalancer {
	return predicate.LoadBalancer(sql.FieldIn(FieldID, ids...))
}

// IDNotIn applies the NotIn predicate on the ID field.
func IDNotIn(ids ...gidx.PrefixedID) predicate.LoadBalancer {
	return predicate.LoadBalancer(sql.FieldNotIn(FieldID, ids...))
}

// IDGT applies the GT predicate on the ID field.
func IDGT(id gidx.PrefixedID) predicate.LoadBalancer {
	return predicate.LoadBalancer(sql.FieldGT(FieldID, id))
}

// IDGTE applies the GTE predicate on the ID field.
func IDGTE(id gidx.PrefixedID) predicate.LoadBalancer {
	return predicate.LoadBalancer(sql.FieldGTE(FieldID, id))
}

// IDLT applies the LT predicate on the ID field.
func IDLT(id gidx.PrefixedID) predicate.LoadBalancer {
	return predicate.LoadBalancer(sql.FieldLT(FieldID, id))
}

// IDLTE applies the LTE predicate on the ID field.
func IDLTE(id gidx.PrefixedID) predicate.LoadBalancer {
	return predicate.LoadBalancer(sql.FieldLTE(FieldID, id))
}

// CreatedAt applies equality check predicate on the "created_at" field. It's identical to CreatedAtEQ.
func CreatedAt(v time.Time) predicate.LoadBalancer {
	return predicate.LoadBalancer(sql.FieldEQ(FieldCreatedAt, v))
}

// UpdatedAt applies equality check predicate on the "updated_at" field. It's identical to UpdatedAtEQ.
func UpdatedAt(v time.Time) predicate.LoadBalancer {
	return predicate.LoadBalancer(sql.FieldEQ(FieldUpdatedAt, v))
}

// DeletedAt applies equality check predicate on the "deleted_at" field. It's identical to DeletedAtEQ.
func DeletedAt(v time.Time) predicate.LoadBalancer {
	return predicate.LoadBalancer(sql.FieldEQ(FieldDeletedAt, v))
}

// Name applies equality check predicate on the "name" field. It's identical to NameEQ.
func Name(v string) predicate.LoadBalancer {
	return predicate.LoadBalancer(sql.FieldEQ(FieldName, v))
}

// OwnerID applies equality check predicate on the "owner_id" field. It's identical to OwnerIDEQ.
func OwnerID(v gidx.PrefixedID) predicate.LoadBalancer {
	return predicate.LoadBalancer(sql.FieldEQ(FieldOwnerID, v))
}

// LocationID applies equality check predicate on the "location_id" field. It's identical to LocationIDEQ.
func LocationID(v gidx.PrefixedID) predicate.LoadBalancer {
	return predicate.LoadBalancer(sql.FieldEQ(FieldLocationID, v))
}

// ProviderID applies equality check predicate on the "provider_id" field. It's identical to ProviderIDEQ.
func ProviderID(v gidx.PrefixedID) predicate.LoadBalancer {
	return predicate.LoadBalancer(sql.FieldEQ(FieldProviderID, v))
}

// CreatedAtEQ applies the EQ predicate on the "created_at" field.
func CreatedAtEQ(v time.Time) predicate.LoadBalancer {
	return predicate.LoadBalancer(sql.FieldEQ(FieldCreatedAt, v))
}

// CreatedAtNEQ applies the NEQ predicate on the "created_at" field.
func CreatedAtNEQ(v time.Time) predicate.LoadBalancer {
	return predicate.LoadBalancer(sql.FieldNEQ(FieldCreatedAt, v))
}

// CreatedAtIn applies the In predicate on the "created_at" field.
func CreatedAtIn(vs ...time.Time) predicate.LoadBalancer {
	return predicate.LoadBalancer(sql.FieldIn(FieldCreatedAt, vs...))
}

// CreatedAtNotIn applies the NotIn predicate on the "created_at" field.
func CreatedAtNotIn(vs ...time.Time) predicate.LoadBalancer {
	return predicate.LoadBalancer(sql.FieldNotIn(FieldCreatedAt, vs...))
}

// CreatedAtGT applies the GT predicate on the "created_at" field.
func CreatedAtGT(v time.Time) predicate.LoadBalancer {
	return predicate.LoadBalancer(sql.FieldGT(FieldCreatedAt, v))
}

// CreatedAtGTE applies the GTE predicate on the "created_at" field.
func CreatedAtGTE(v time.Time) predicate.LoadBalancer {
	return predicate.LoadBalancer(sql.FieldGTE(FieldCreatedAt, v))
}

// CreatedAtLT applies the LT predicate on the "created_at" field.
func CreatedAtLT(v time.Time) predicate.LoadBalancer {
	return predicate.LoadBalancer(sql.FieldLT(FieldCreatedAt, v))
}

// CreatedAtLTE applies the LTE predicate on the "created_at" field.
func CreatedAtLTE(v time.Time) predicate.LoadBalancer {
	return predicate.LoadBalancer(sql.FieldLTE(FieldCreatedAt, v))
}

// UpdatedAtEQ applies the EQ predicate on the "updated_at" field.
func UpdatedAtEQ(v time.Time) predicate.LoadBalancer {
	return predicate.LoadBalancer(sql.FieldEQ(FieldUpdatedAt, v))
}

// UpdatedAtNEQ applies the NEQ predicate on the "updated_at" field.
func UpdatedAtNEQ(v time.Time) predicate.LoadBalancer {
	return predicate.LoadBalancer(sql.FieldNEQ(FieldUpdatedAt, v))
}

// UpdatedAtIn applies the In predicate on the "updated_at" field.
func UpdatedAtIn(vs ...time.Time) predicate.LoadBalancer {
	return predicate.LoadBalancer(sql.FieldIn(FieldUpdatedAt, vs...))
}

// UpdatedAtNotIn applies the NotIn predicate on the "updated_at" field.
func UpdatedAtNotIn(vs ...time.Time) predicate.LoadBalancer {
	return predicate.LoadBalancer(sql.FieldNotIn(FieldUpdatedAt, vs...))
}

// UpdatedAtGT applies the GT predicate on the "updated_at" field.
func UpdatedAtGT(v time.Time) predicate.LoadBalancer {
	return predicate.LoadBalancer(sql.FieldGT(FieldUpdatedAt, v))
}

// UpdatedAtGTE applies the GTE predicate on the "updated_at" field.
func UpdatedAtGTE(v time.Time) predicate.LoadBalancer {
	return predicate.LoadBalancer(sql.FieldGTE(FieldUpdatedAt, v))
}

// UpdatedAtLT applies the LT predicate on the "updated_at" field.
func UpdatedAtLT(v time.Time) predicate.LoadBalancer {
	return predicate.LoadBalancer(sql.FieldLT(FieldUpdatedAt, v))
}

// UpdatedAtLTE applies the LTE predicate on the "updated_at" field.
func UpdatedAtLTE(v time.Time) predicate.LoadBalancer {
	return predicate.LoadBalancer(sql.FieldLTE(FieldUpdatedAt, v))
}

// DeletedAtEQ applies the EQ predicate on the "deleted_at" field.
func DeletedAtEQ(v time.Time) predicate.LoadBalancer {
	return predicate.LoadBalancer(sql.FieldEQ(FieldDeletedAt, v))
}

// DeletedAtNEQ applies the NEQ predicate on the "deleted_at" field.
func DeletedAtNEQ(v time.Time) predicate.LoadBalancer {
	return predicate.LoadBalancer(sql.FieldNEQ(FieldDeletedAt, v))
}

// DeletedAtIn applies the In predicate on the "deleted_at" field.
func DeletedAtIn(vs ...time.Time) predicate.LoadBalancer {
	return predicate.LoadBalancer(sql.FieldIn(FieldDeletedAt, vs...))
}

// DeletedAtNotIn applies the NotIn predicate on the "deleted_at" field.
func DeletedAtNotIn(vs ...time.Time) predicate.LoadBalancer {
	return predicate.LoadBalancer(sql.FieldNotIn(FieldDeletedAt, vs...))
}

// DeletedAtGT applies the GT predicate on the "deleted_at" field.
func DeletedAtGT(v time.Time) predicate.LoadBalancer {
	return predicate.LoadBalancer(sql.FieldGT(FieldDeletedAt, v))
}

// DeletedAtGTE applies the GTE predicate on the "deleted_at" field.
func DeletedAtGTE(v time.Time) predicate.LoadBalancer {
	return predicate.LoadBalancer(sql.FieldGTE(FieldDeletedAt, v))
}

// DeletedAtLT applies the LT predicate on the "deleted_at" field.
func DeletedAtLT(v time.Time) predicate.LoadBalancer {
	return predicate.LoadBalancer(sql.FieldLT(FieldDeletedAt, v))
}

// DeletedAtLTE applies the LTE predicate on the "deleted_at" field.
func DeletedAtLTE(v time.Time) predicate.LoadBalancer {
	return predicate.LoadBalancer(sql.FieldLTE(FieldDeletedAt, v))
}

// DeletedAtIsNil applies the IsNil predicate on the "deleted_at" field.
func DeletedAtIsNil() predicate.LoadBalancer {
	return predicate.LoadBalancer(sql.FieldIsNull(FieldDeletedAt))
}

// DeletedAtNotNil applies the NotNil predicate on the "deleted_at" field.
func DeletedAtNotNil() predicate.LoadBalancer {
	return predicate.LoadBalancer(sql.FieldNotNull(FieldDeletedAt))
}

// NameEQ applies the EQ predicate on the "name" field.
func NameEQ(v string) predicate.LoadBalancer {
	return predicate.LoadBalancer(sql.FieldEQ(FieldName, v))
}

// NameNEQ applies the NEQ predicate on the "name" field.
func NameNEQ(v string) predicate.LoadBalancer {
	return predicate.LoadBalancer(sql.FieldNEQ(FieldName, v))
}

// NameIn applies the In predicate on the "name" field.
func NameIn(vs ...string) predicate.LoadBalancer {
	return predicate.LoadBalancer(sql.FieldIn(FieldName, vs...))
}

// NameNotIn applies the NotIn predicate on the "name" field.
func NameNotIn(vs ...string) predicate.LoadBalancer {
	return predicate.LoadBalancer(sql.FieldNotIn(FieldName, vs...))
}

// NameGT applies the GT predicate on the "name" field.
func NameGT(v string) predicate.LoadBalancer {
	return predicate.LoadBalancer(sql.FieldGT(FieldName, v))
}

// NameGTE applies the GTE predicate on the "name" field.
func NameGTE(v string) predicate.LoadBalancer {
	return predicate.LoadBalancer(sql.FieldGTE(FieldName, v))
}

// NameLT applies the LT predicate on the "name" field.
func NameLT(v string) predicate.LoadBalancer {
	return predicate.LoadBalancer(sql.FieldLT(FieldName, v))
}

// NameLTE applies the LTE predicate on the "name" field.
func NameLTE(v string) predicate.LoadBalancer {
	return predicate.LoadBalancer(sql.FieldLTE(FieldName, v))
}

// NameContains applies the Contains predicate on the "name" field.
func NameContains(v string) predicate.LoadBalancer {
	return predicate.LoadBalancer(sql.FieldContains(FieldName, v))
}

// NameHasPrefix applies the HasPrefix predicate on the "name" field.
func NameHasPrefix(v string) predicate.LoadBalancer {
	return predicate.LoadBalancer(sql.FieldHasPrefix(FieldName, v))
}

// NameHasSuffix applies the HasSuffix predicate on the "name" field.
func NameHasSuffix(v string) predicate.LoadBalancer {
	return predicate.LoadBalancer(sql.FieldHasSuffix(FieldName, v))
}

// NameEqualFold applies the EqualFold predicate on the "name" field.
func NameEqualFold(v string) predicate.LoadBalancer {
	return predicate.LoadBalancer(sql.FieldEqualFold(FieldName, v))
}

// NameContainsFold applies the ContainsFold predicate on the "name" field.
func NameContainsFold(v string) predicate.LoadBalancer {
	return predicate.LoadBalancer(sql.FieldContainsFold(FieldName, v))
}

// OwnerIDEQ applies the EQ predicate on the "owner_id" field.
func OwnerIDEQ(v gidx.PrefixedID) predicate.LoadBalancer {
	return predicate.LoadBalancer(sql.FieldEQ(FieldOwnerID, v))
}

// OwnerIDNEQ applies the NEQ predicate on the "owner_id" field.
func OwnerIDNEQ(v gidx.PrefixedID) predicate.LoadBalancer {
	return predicate.LoadBalancer(sql.FieldNEQ(FieldOwnerID, v))
}

// OwnerIDIn applies the In predicate on the "owner_id" field.
func OwnerIDIn(vs ...gidx.PrefixedID) predicate.LoadBalancer {
	return predicate.LoadBalancer(sql.FieldIn(FieldOwnerID, vs...))
}

// OwnerIDNotIn applies the NotIn predicate on the "owner_id" field.
func OwnerIDNotIn(vs ...gidx.PrefixedID) predicate.LoadBalancer {
	return predicate.LoadBalancer(sql.FieldNotIn(FieldOwnerID, vs...))
}

// OwnerIDGT applies the GT predicate on the "owner_id" field.
func OwnerIDGT(v gidx.PrefixedID) predicate.LoadBalancer {
	return predicate.LoadBalancer(sql.FieldGT(FieldOwnerID, v))
}

// OwnerIDGTE applies the GTE predicate on the "owner_id" field.
func OwnerIDGTE(v gidx.PrefixedID) predicate.LoadBalancer {
	return predicate.LoadBalancer(sql.FieldGTE(FieldOwnerID, v))
}

// OwnerIDLT applies the LT predicate on the "owner_id" field.
func OwnerIDLT(v gidx.PrefixedID) predicate.LoadBalancer {
	return predicate.LoadBalancer(sql.FieldLT(FieldOwnerID, v))
}

// OwnerIDLTE applies the LTE predicate on the "owner_id" field.
func OwnerIDLTE(v gidx.PrefixedID) predicate.LoadBalancer {
	return predicate.LoadBalancer(sql.FieldLTE(FieldOwnerID, v))
}

// OwnerIDContains applies the Contains predicate on the "owner_id" field.
func OwnerIDContains(v gidx.PrefixedID) predicate.LoadBalancer {
	vc := string(v)
	return predicate.LoadBalancer(sql.FieldContains(FieldOwnerID, vc))
}

// OwnerIDHasPrefix applies the HasPrefix predicate on the "owner_id" field.
func OwnerIDHasPrefix(v gidx.PrefixedID) predicate.LoadBalancer {
	vc := string(v)
	return predicate.LoadBalancer(sql.FieldHasPrefix(FieldOwnerID, vc))
}

// OwnerIDHasSuffix applies the HasSuffix predicate on the "owner_id" field.
func OwnerIDHasSuffix(v gidx.PrefixedID) predicate.LoadBalancer {
	vc := string(v)
	return predicate.LoadBalancer(sql.FieldHasSuffix(FieldOwnerID, vc))
}

// OwnerIDEqualFold applies the EqualFold predicate on the "owner_id" field.
func OwnerIDEqualFold(v gidx.PrefixedID) predicate.LoadBalancer {
	vc := string(v)
	return predicate.LoadBalancer(sql.FieldEqualFold(FieldOwnerID, vc))
}

// OwnerIDContainsFold applies the ContainsFold predicate on the "owner_id" field.
func OwnerIDContainsFold(v gidx.PrefixedID) predicate.LoadBalancer {
	vc := string(v)
	return predicate.LoadBalancer(sql.FieldContainsFold(FieldOwnerID, vc))
}

// LocationIDEQ applies the EQ predicate on the "location_id" field.
func LocationIDEQ(v gidx.PrefixedID) predicate.LoadBalancer {
	return predicate.LoadBalancer(sql.FieldEQ(FieldLocationID, v))
}

// LocationIDNEQ applies the NEQ predicate on the "location_id" field.
func LocationIDNEQ(v gidx.PrefixedID) predicate.LoadBalancer {
	return predicate.LoadBalancer(sql.FieldNEQ(FieldLocationID, v))
}

// LocationIDIn applies the In predicate on the "location_id" field.
func LocationIDIn(vs ...gidx.PrefixedID) predicate.LoadBalancer {
	return predicate.LoadBalancer(sql.FieldIn(FieldLocationID, vs...))
}

// LocationIDNotIn applies the NotIn predicate on the "location_id" field.
func LocationIDNotIn(vs ...gidx.PrefixedID) predicate.LoadBalancer {
	return predicate.LoadBalancer(sql.FieldNotIn(FieldLocationID, vs...))
}

// LocationIDGT applies the GT predicate on the "location_id" field.
func LocationIDGT(v gidx.PrefixedID) predicate.LoadBalancer {
	return predicate.LoadBalancer(sql.FieldGT(FieldLocationID, v))
}

// LocationIDGTE applies the GTE predicate on the "location_id" field.
func LocationIDGTE(v gidx.PrefixedID) predicate.LoadBalancer {
	return predicate.LoadBalancer(sql.FieldGTE(FieldLocationID, v))
}

// LocationIDLT applies the LT predicate on the "location_id" field.
func LocationIDLT(v gidx.PrefixedID) predicate.LoadBalancer {
	return predicate.LoadBalancer(sql.FieldLT(FieldLocationID, v))
}

// LocationIDLTE applies the LTE predicate on the "location_id" field.
func LocationIDLTE(v gidx.PrefixedID) predicate.LoadBalancer {
	return predicate.LoadBalancer(sql.FieldLTE(FieldLocationID, v))
}

// LocationIDContains applies the Contains predicate on the "location_id" field.
func LocationIDContains(v gidx.PrefixedID) predicate.LoadBalancer {
	vc := string(v)
	return predicate.LoadBalancer(sql.FieldContains(FieldLocationID, vc))
}

// LocationIDHasPrefix applies the HasPrefix predicate on the "location_id" field.
func LocationIDHasPrefix(v gidx.PrefixedID) predicate.LoadBalancer {
	vc := string(v)
	return predicate.LoadBalancer(sql.FieldHasPrefix(FieldLocationID, vc))
}

// LocationIDHasSuffix applies the HasSuffix predicate on the "location_id" field.
func LocationIDHasSuffix(v gidx.PrefixedID) predicate.LoadBalancer {
	vc := string(v)
	return predicate.LoadBalancer(sql.FieldHasSuffix(FieldLocationID, vc))
}

// LocationIDEqualFold applies the EqualFold predicate on the "location_id" field.
func LocationIDEqualFold(v gidx.PrefixedID) predicate.LoadBalancer {
	vc := string(v)
	return predicate.LoadBalancer(sql.FieldEqualFold(FieldLocationID, vc))
}

// LocationIDContainsFold applies the ContainsFold predicate on the "location_id" field.
func LocationIDContainsFold(v gidx.PrefixedID) predicate.LoadBalancer {
	vc := string(v)
	return predicate.LoadBalancer(sql.FieldContainsFold(FieldLocationID, vc))
}

// ProviderIDEQ applies the EQ predicate on the "provider_id" field.
func ProviderIDEQ(v gidx.PrefixedID) predicate.LoadBalancer {
	return predicate.LoadBalancer(sql.FieldEQ(FieldProviderID, v))
}

// ProviderIDNEQ applies the NEQ predicate on the "provider_id" field.
func ProviderIDNEQ(v gidx.PrefixedID) predicate.LoadBalancer {
	return predicate.LoadBalancer(sql.FieldNEQ(FieldProviderID, v))
}

// ProviderIDIn applies the In predicate on the "provider_id" field.
func ProviderIDIn(vs ...gidx.PrefixedID) predicate.LoadBalancer {
	return predicate.LoadBalancer(sql.FieldIn(FieldProviderID, vs...))
}

// ProviderIDNotIn applies the NotIn predicate on the "provider_id" field.
func ProviderIDNotIn(vs ...gidx.PrefixedID) predicate.LoadBalancer {
	return predicate.LoadBalancer(sql.FieldNotIn(FieldProviderID, vs...))
}

// ProviderIDGT applies the GT predicate on the "provider_id" field.
func ProviderIDGT(v gidx.PrefixedID) predicate.LoadBalancer {
	return predicate.LoadBalancer(sql.FieldGT(FieldProviderID, v))
}

// ProviderIDGTE applies the GTE predicate on the "provider_id" field.
func ProviderIDGTE(v gidx.PrefixedID) predicate.LoadBalancer {
	return predicate.LoadBalancer(sql.FieldGTE(FieldProviderID, v))
}

// ProviderIDLT applies the LT predicate on the "provider_id" field.
func ProviderIDLT(v gidx.PrefixedID) predicate.LoadBalancer {
	return predicate.LoadBalancer(sql.FieldLT(FieldProviderID, v))
}

// ProviderIDLTE applies the LTE predicate on the "provider_id" field.
func ProviderIDLTE(v gidx.PrefixedID) predicate.LoadBalancer {
	return predicate.LoadBalancer(sql.FieldLTE(FieldProviderID, v))
}

// ProviderIDContains applies the Contains predicate on the "provider_id" field.
func ProviderIDContains(v gidx.PrefixedID) predicate.LoadBalancer {
	vc := string(v)
	return predicate.LoadBalancer(sql.FieldContains(FieldProviderID, vc))
}

// ProviderIDHasPrefix applies the HasPrefix predicate on the "provider_id" field.
func ProviderIDHasPrefix(v gidx.PrefixedID) predicate.LoadBalancer {
	vc := string(v)
	return predicate.LoadBalancer(sql.FieldHasPrefix(FieldProviderID, vc))
}

// ProviderIDHasSuffix applies the HasSuffix predicate on the "provider_id" field.
func ProviderIDHasSuffix(v gidx.PrefixedID) predicate.LoadBalancer {
	vc := string(v)
	return predicate.LoadBalancer(sql.FieldHasSuffix(FieldProviderID, vc))
}

// ProviderIDEqualFold applies the EqualFold predicate on the "provider_id" field.
func ProviderIDEqualFold(v gidx.PrefixedID) predicate.LoadBalancer {
	vc := string(v)
	return predicate.LoadBalancer(sql.FieldEqualFold(FieldProviderID, vc))
}

// ProviderIDContainsFold applies the ContainsFold predicate on the "provider_id" field.
func ProviderIDContainsFold(v gidx.PrefixedID) predicate.LoadBalancer {
	vc := string(v)
	return predicate.LoadBalancer(sql.FieldContainsFold(FieldProviderID, vc))
}

// HasPorts applies the HasEdge predicate on the "ports" edge.
func HasPorts() predicate.LoadBalancer {
	return predicate.LoadBalancer(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.Edge(sqlgraph.O2M, true, PortsTable, PortsColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasPortsWith applies the HasEdge predicate on the "ports" edge with a given conditions (other predicates).
func HasPortsWith(preds ...predicate.Port) predicate.LoadBalancer {
	return predicate.LoadBalancer(func(s *sql.Selector) {
		step := newPortsStep()
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// HasProvider applies the HasEdge predicate on the "provider" edge.
func HasProvider() predicate.LoadBalancer {
	return predicate.LoadBalancer(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.Edge(sqlgraph.M2O, false, ProviderTable, ProviderColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasProviderWith applies the HasEdge predicate on the "provider" edge with a given conditions (other predicates).
func HasProviderWith(preds ...predicate.Provider) predicate.LoadBalancer {
	return predicate.LoadBalancer(func(s *sql.Selector) {
		step := newProviderStep()
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// And groups predicates with the AND operator between them.
func And(predicates ...predicate.LoadBalancer) predicate.LoadBalancer {
	return predicate.LoadBalancer(func(s *sql.Selector) {
		s1 := s.Clone().SetP(nil)
		for _, p := range predicates {
			p(s1)
		}
		s.Where(s1.P())
	})
}

// Or groups predicates with the OR operator between them.
func Or(predicates ...predicate.LoadBalancer) predicate.LoadBalancer {
	return predicate.LoadBalancer(func(s *sql.Selector) {
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
func Not(p predicate.LoadBalancer) predicate.LoadBalancer {
	return predicate.LoadBalancer(func(s *sql.Selector) {
		p(s.Not())
	})
}
