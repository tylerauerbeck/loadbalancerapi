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

package generated

import (
	"context"
	"errors"
	"fmt"
	"time"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"go.infratographer.com/load-balancer-api/internal/ent/generated/loadbalancer"
	"go.infratographer.com/load-balancer-api/internal/ent/generated/port"
	"go.infratographer.com/load-balancer-api/internal/ent/generated/predicate"
	"go.infratographer.com/x/gidx"
)

// LoadBalancerUpdate is the builder for updating LoadBalancer entities.
type LoadBalancerUpdate struct {
	config
	hooks    []Hook
	mutation *LoadBalancerMutation
}

// Where appends a list predicates to the LoadBalancerUpdate builder.
func (lbu *LoadBalancerUpdate) Where(ps ...predicate.LoadBalancer) *LoadBalancerUpdate {
	lbu.mutation.Where(ps...)
	return lbu
}

// SetUpdatedBy sets the "updated_by" field.
func (lbu *LoadBalancerUpdate) SetUpdatedBy(s string) *LoadBalancerUpdate {
	lbu.mutation.SetUpdatedBy(s)
	return lbu
}

// SetNillableUpdatedBy sets the "updated_by" field if the given value is not nil.
func (lbu *LoadBalancerUpdate) SetNillableUpdatedBy(s *string) *LoadBalancerUpdate {
	if s != nil {
		lbu.SetUpdatedBy(*s)
	}
	return lbu
}

// ClearUpdatedBy clears the value of the "updated_by" field.
func (lbu *LoadBalancerUpdate) ClearUpdatedBy() *LoadBalancerUpdate {
	lbu.mutation.ClearUpdatedBy()
	return lbu
}

// SetDeletedAt sets the "deleted_at" field.
func (lbu *LoadBalancerUpdate) SetDeletedAt(t time.Time) *LoadBalancerUpdate {
	lbu.mutation.SetDeletedAt(t)
	return lbu
}

// SetNillableDeletedAt sets the "deleted_at" field if the given value is not nil.
func (lbu *LoadBalancerUpdate) SetNillableDeletedAt(t *time.Time) *LoadBalancerUpdate {
	if t != nil {
		lbu.SetDeletedAt(*t)
	}
	return lbu
}

// ClearDeletedAt clears the value of the "deleted_at" field.
func (lbu *LoadBalancerUpdate) ClearDeletedAt() *LoadBalancerUpdate {
	lbu.mutation.ClearDeletedAt()
	return lbu
}

// SetDeletedBy sets the "deleted_by" field.
func (lbu *LoadBalancerUpdate) SetDeletedBy(s string) *LoadBalancerUpdate {
	lbu.mutation.SetDeletedBy(s)
	return lbu
}

// SetNillableDeletedBy sets the "deleted_by" field if the given value is not nil.
func (lbu *LoadBalancerUpdate) SetNillableDeletedBy(s *string) *LoadBalancerUpdate {
	if s != nil {
		lbu.SetDeletedBy(*s)
	}
	return lbu
}

// ClearDeletedBy clears the value of the "deleted_by" field.
func (lbu *LoadBalancerUpdate) ClearDeletedBy() *LoadBalancerUpdate {
	lbu.mutation.ClearDeletedBy()
	return lbu
}

// SetName sets the "name" field.
func (lbu *LoadBalancerUpdate) SetName(s string) *LoadBalancerUpdate {
	lbu.mutation.SetName(s)
	return lbu
}

// AddPortIDs adds the "ports" edge to the Port entity by IDs.
func (lbu *LoadBalancerUpdate) AddPortIDs(ids ...gidx.PrefixedID) *LoadBalancerUpdate {
	lbu.mutation.AddPortIDs(ids...)
	return lbu
}

// AddPorts adds the "ports" edges to the Port entity.
func (lbu *LoadBalancerUpdate) AddPorts(p ...*Port) *LoadBalancerUpdate {
	ids := make([]gidx.PrefixedID, len(p))
	for i := range p {
		ids[i] = p[i].ID
	}
	return lbu.AddPortIDs(ids...)
}

// Mutation returns the LoadBalancerMutation object of the builder.
func (lbu *LoadBalancerUpdate) Mutation() *LoadBalancerMutation {
	return lbu.mutation
}

// ClearPorts clears all "ports" edges to the Port entity.
func (lbu *LoadBalancerUpdate) ClearPorts() *LoadBalancerUpdate {
	lbu.mutation.ClearPorts()
	return lbu
}

// RemovePortIDs removes the "ports" edge to Port entities by IDs.
func (lbu *LoadBalancerUpdate) RemovePortIDs(ids ...gidx.PrefixedID) *LoadBalancerUpdate {
	lbu.mutation.RemovePortIDs(ids...)
	return lbu
}

// RemovePorts removes "ports" edges to Port entities.
func (lbu *LoadBalancerUpdate) RemovePorts(p ...*Port) *LoadBalancerUpdate {
	ids := make([]gidx.PrefixedID, len(p))
	for i := range p {
		ids[i] = p[i].ID
	}
	return lbu.RemovePortIDs(ids...)
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (lbu *LoadBalancerUpdate) Save(ctx context.Context) (int, error) {
	if err := lbu.defaults(); err != nil {
		return 0, err
	}
	return withHooks(ctx, lbu.sqlSave, lbu.mutation, lbu.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (lbu *LoadBalancerUpdate) SaveX(ctx context.Context) int {
	affected, err := lbu.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (lbu *LoadBalancerUpdate) Exec(ctx context.Context) error {
	_, err := lbu.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (lbu *LoadBalancerUpdate) ExecX(ctx context.Context) {
	if err := lbu.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (lbu *LoadBalancerUpdate) defaults() error {
	if _, ok := lbu.mutation.UpdatedAt(); !ok {
		if loadbalancer.UpdateDefaultUpdatedAt == nil {
			return fmt.Errorf("generated: uninitialized loadbalancer.UpdateDefaultUpdatedAt (forgotten import generated/runtime?)")
		}
		v := loadbalancer.UpdateDefaultUpdatedAt()
		lbu.mutation.SetUpdatedAt(v)
	}
	return nil
}

// check runs all checks and user-defined validators on the builder.
func (lbu *LoadBalancerUpdate) check() error {
	if v, ok := lbu.mutation.Name(); ok {
		if err := loadbalancer.NameValidator(v); err != nil {
			return &ValidationError{Name: "name", err: fmt.Errorf(`generated: validator failed for field "LoadBalancer.name": %w`, err)}
		}
	}
	if _, ok := lbu.mutation.ProviderID(); lbu.mutation.ProviderCleared() && !ok {
		return errors.New(`generated: clearing a required unique edge "LoadBalancer.provider"`)
	}
	return nil
}

func (lbu *LoadBalancerUpdate) sqlSave(ctx context.Context) (n int, err error) {
	if err := lbu.check(); err != nil {
		return n, err
	}
	_spec := sqlgraph.NewUpdateSpec(loadbalancer.Table, loadbalancer.Columns, sqlgraph.NewFieldSpec(loadbalancer.FieldID, field.TypeString))
	if ps := lbu.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := lbu.mutation.UpdatedAt(); ok {
		_spec.SetField(loadbalancer.FieldUpdatedAt, field.TypeTime, value)
	}
	if lbu.mutation.CreatedByCleared() {
		_spec.ClearField(loadbalancer.FieldCreatedBy, field.TypeString)
	}
	if value, ok := lbu.mutation.UpdatedBy(); ok {
		_spec.SetField(loadbalancer.FieldUpdatedBy, field.TypeString, value)
	}
	if lbu.mutation.UpdatedByCleared() {
		_spec.ClearField(loadbalancer.FieldUpdatedBy, field.TypeString)
	}
	if value, ok := lbu.mutation.DeletedAt(); ok {
		_spec.SetField(loadbalancer.FieldDeletedAt, field.TypeTime, value)
	}
	if lbu.mutation.DeletedAtCleared() {
		_spec.ClearField(loadbalancer.FieldDeletedAt, field.TypeTime)
	}
	if value, ok := lbu.mutation.DeletedBy(); ok {
		_spec.SetField(loadbalancer.FieldDeletedBy, field.TypeString, value)
	}
	if lbu.mutation.DeletedByCleared() {
		_spec.ClearField(loadbalancer.FieldDeletedBy, field.TypeString)
	}
	if value, ok := lbu.mutation.Name(); ok {
		_spec.SetField(loadbalancer.FieldName, field.TypeString, value)
	}
	if lbu.mutation.PortsCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: true,
			Table:   loadbalancer.PortsTable,
			Columns: []string{loadbalancer.PortsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(port.FieldID, field.TypeString),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := lbu.mutation.RemovedPortsIDs(); len(nodes) > 0 && !lbu.mutation.PortsCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: true,
			Table:   loadbalancer.PortsTable,
			Columns: []string{loadbalancer.PortsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(port.FieldID, field.TypeString),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := lbu.mutation.PortsIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: true,
			Table:   loadbalancer.PortsTable,
			Columns: []string{loadbalancer.PortsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(port.FieldID, field.TypeString),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if n, err = sqlgraph.UpdateNodes(ctx, lbu.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{loadbalancer.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return 0, err
	}
	lbu.mutation.done = true
	return n, nil
}

// LoadBalancerUpdateOne is the builder for updating a single LoadBalancer entity.
type LoadBalancerUpdateOne struct {
	config
	fields   []string
	hooks    []Hook
	mutation *LoadBalancerMutation
}

// SetUpdatedBy sets the "updated_by" field.
func (lbuo *LoadBalancerUpdateOne) SetUpdatedBy(s string) *LoadBalancerUpdateOne {
	lbuo.mutation.SetUpdatedBy(s)
	return lbuo
}

// SetNillableUpdatedBy sets the "updated_by" field if the given value is not nil.
func (lbuo *LoadBalancerUpdateOne) SetNillableUpdatedBy(s *string) *LoadBalancerUpdateOne {
	if s != nil {
		lbuo.SetUpdatedBy(*s)
	}
	return lbuo
}

// ClearUpdatedBy clears the value of the "updated_by" field.
func (lbuo *LoadBalancerUpdateOne) ClearUpdatedBy() *LoadBalancerUpdateOne {
	lbuo.mutation.ClearUpdatedBy()
	return lbuo
}

// SetDeletedAt sets the "deleted_at" field.
func (lbuo *LoadBalancerUpdateOne) SetDeletedAt(t time.Time) *LoadBalancerUpdateOne {
	lbuo.mutation.SetDeletedAt(t)
	return lbuo
}

// SetNillableDeletedAt sets the "deleted_at" field if the given value is not nil.
func (lbuo *LoadBalancerUpdateOne) SetNillableDeletedAt(t *time.Time) *LoadBalancerUpdateOne {
	if t != nil {
		lbuo.SetDeletedAt(*t)
	}
	return lbuo
}

// ClearDeletedAt clears the value of the "deleted_at" field.
func (lbuo *LoadBalancerUpdateOne) ClearDeletedAt() *LoadBalancerUpdateOne {
	lbuo.mutation.ClearDeletedAt()
	return lbuo
}

// SetDeletedBy sets the "deleted_by" field.
func (lbuo *LoadBalancerUpdateOne) SetDeletedBy(s string) *LoadBalancerUpdateOne {
	lbuo.mutation.SetDeletedBy(s)
	return lbuo
}

// SetNillableDeletedBy sets the "deleted_by" field if the given value is not nil.
func (lbuo *LoadBalancerUpdateOne) SetNillableDeletedBy(s *string) *LoadBalancerUpdateOne {
	if s != nil {
		lbuo.SetDeletedBy(*s)
	}
	return lbuo
}

// ClearDeletedBy clears the value of the "deleted_by" field.
func (lbuo *LoadBalancerUpdateOne) ClearDeletedBy() *LoadBalancerUpdateOne {
	lbuo.mutation.ClearDeletedBy()
	return lbuo
}

// SetName sets the "name" field.
func (lbuo *LoadBalancerUpdateOne) SetName(s string) *LoadBalancerUpdateOne {
	lbuo.mutation.SetName(s)
	return lbuo
}

// AddPortIDs adds the "ports" edge to the Port entity by IDs.
func (lbuo *LoadBalancerUpdateOne) AddPortIDs(ids ...gidx.PrefixedID) *LoadBalancerUpdateOne {
	lbuo.mutation.AddPortIDs(ids...)
	return lbuo
}

// AddPorts adds the "ports" edges to the Port entity.
func (lbuo *LoadBalancerUpdateOne) AddPorts(p ...*Port) *LoadBalancerUpdateOne {
	ids := make([]gidx.PrefixedID, len(p))
	for i := range p {
		ids[i] = p[i].ID
	}
	return lbuo.AddPortIDs(ids...)
}

// Mutation returns the LoadBalancerMutation object of the builder.
func (lbuo *LoadBalancerUpdateOne) Mutation() *LoadBalancerMutation {
	return lbuo.mutation
}

// ClearPorts clears all "ports" edges to the Port entity.
func (lbuo *LoadBalancerUpdateOne) ClearPorts() *LoadBalancerUpdateOne {
	lbuo.mutation.ClearPorts()
	return lbuo
}

// RemovePortIDs removes the "ports" edge to Port entities by IDs.
func (lbuo *LoadBalancerUpdateOne) RemovePortIDs(ids ...gidx.PrefixedID) *LoadBalancerUpdateOne {
	lbuo.mutation.RemovePortIDs(ids...)
	return lbuo
}

// RemovePorts removes "ports" edges to Port entities.
func (lbuo *LoadBalancerUpdateOne) RemovePorts(p ...*Port) *LoadBalancerUpdateOne {
	ids := make([]gidx.PrefixedID, len(p))
	for i := range p {
		ids[i] = p[i].ID
	}
	return lbuo.RemovePortIDs(ids...)
}

// Where appends a list predicates to the LoadBalancerUpdate builder.
func (lbuo *LoadBalancerUpdateOne) Where(ps ...predicate.LoadBalancer) *LoadBalancerUpdateOne {
	lbuo.mutation.Where(ps...)
	return lbuo
}

// Select allows selecting one or more fields (columns) of the returned entity.
// The default is selecting all fields defined in the entity schema.
func (lbuo *LoadBalancerUpdateOne) Select(field string, fields ...string) *LoadBalancerUpdateOne {
	lbuo.fields = append([]string{field}, fields...)
	return lbuo
}

// Save executes the query and returns the updated LoadBalancer entity.
func (lbuo *LoadBalancerUpdateOne) Save(ctx context.Context) (*LoadBalancer, error) {
	if err := lbuo.defaults(); err != nil {
		return nil, err
	}
	return withHooks(ctx, lbuo.sqlSave, lbuo.mutation, lbuo.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (lbuo *LoadBalancerUpdateOne) SaveX(ctx context.Context) *LoadBalancer {
	node, err := lbuo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (lbuo *LoadBalancerUpdateOne) Exec(ctx context.Context) error {
	_, err := lbuo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (lbuo *LoadBalancerUpdateOne) ExecX(ctx context.Context) {
	if err := lbuo.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (lbuo *LoadBalancerUpdateOne) defaults() error {
	if _, ok := lbuo.mutation.UpdatedAt(); !ok {
		if loadbalancer.UpdateDefaultUpdatedAt == nil {
			return fmt.Errorf("generated: uninitialized loadbalancer.UpdateDefaultUpdatedAt (forgotten import generated/runtime?)")
		}
		v := loadbalancer.UpdateDefaultUpdatedAt()
		lbuo.mutation.SetUpdatedAt(v)
	}
	return nil
}

// check runs all checks and user-defined validators on the builder.
func (lbuo *LoadBalancerUpdateOne) check() error {
	if v, ok := lbuo.mutation.Name(); ok {
		if err := loadbalancer.NameValidator(v); err != nil {
			return &ValidationError{Name: "name", err: fmt.Errorf(`generated: validator failed for field "LoadBalancer.name": %w`, err)}
		}
	}
	if _, ok := lbuo.mutation.ProviderID(); lbuo.mutation.ProviderCleared() && !ok {
		return errors.New(`generated: clearing a required unique edge "LoadBalancer.provider"`)
	}
	return nil
}

func (lbuo *LoadBalancerUpdateOne) sqlSave(ctx context.Context) (_node *LoadBalancer, err error) {
	if err := lbuo.check(); err != nil {
		return _node, err
	}
	_spec := sqlgraph.NewUpdateSpec(loadbalancer.Table, loadbalancer.Columns, sqlgraph.NewFieldSpec(loadbalancer.FieldID, field.TypeString))
	id, ok := lbuo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "id", err: errors.New(`generated: missing "LoadBalancer.id" for update`)}
	}
	_spec.Node.ID.Value = id
	if fields := lbuo.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, loadbalancer.FieldID)
		for _, f := range fields {
			if !loadbalancer.ValidColumn(f) {
				return nil, &ValidationError{Name: f, err: fmt.Errorf("generated: invalid field %q for query", f)}
			}
			if f != loadbalancer.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, f)
			}
		}
	}
	if ps := lbuo.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := lbuo.mutation.UpdatedAt(); ok {
		_spec.SetField(loadbalancer.FieldUpdatedAt, field.TypeTime, value)
	}
	if lbuo.mutation.CreatedByCleared() {
		_spec.ClearField(loadbalancer.FieldCreatedBy, field.TypeString)
	}
	if value, ok := lbuo.mutation.UpdatedBy(); ok {
		_spec.SetField(loadbalancer.FieldUpdatedBy, field.TypeString, value)
	}
	if lbuo.mutation.UpdatedByCleared() {
		_spec.ClearField(loadbalancer.FieldUpdatedBy, field.TypeString)
	}
	if value, ok := lbuo.mutation.DeletedAt(); ok {
		_spec.SetField(loadbalancer.FieldDeletedAt, field.TypeTime, value)
	}
	if lbuo.mutation.DeletedAtCleared() {
		_spec.ClearField(loadbalancer.FieldDeletedAt, field.TypeTime)
	}
	if value, ok := lbuo.mutation.DeletedBy(); ok {
		_spec.SetField(loadbalancer.FieldDeletedBy, field.TypeString, value)
	}
	if lbuo.mutation.DeletedByCleared() {
		_spec.ClearField(loadbalancer.FieldDeletedBy, field.TypeString)
	}
	if value, ok := lbuo.mutation.Name(); ok {
		_spec.SetField(loadbalancer.FieldName, field.TypeString, value)
	}
	if lbuo.mutation.PortsCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: true,
			Table:   loadbalancer.PortsTable,
			Columns: []string{loadbalancer.PortsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(port.FieldID, field.TypeString),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := lbuo.mutation.RemovedPortsIDs(); len(nodes) > 0 && !lbuo.mutation.PortsCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: true,
			Table:   loadbalancer.PortsTable,
			Columns: []string{loadbalancer.PortsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(port.FieldID, field.TypeString),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := lbuo.mutation.PortsIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: true,
			Table:   loadbalancer.PortsTable,
			Columns: []string{loadbalancer.PortsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(port.FieldID, field.TypeString),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	_node = &LoadBalancer{config: lbuo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, lbuo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{loadbalancer.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	lbuo.mutation.done = true
	return _node, nil
}
