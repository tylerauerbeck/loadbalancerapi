// Copyright 2023 The Infratographer Authors
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
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
	"fmt"
	"math"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"go.infratographer.com/load-balancer-api/internal/ent/generated/loadbalancer"
	"go.infratographer.com/load-balancer-api/internal/ent/generated/loadbalancerannotation"
	"go.infratographer.com/load-balancer-api/internal/ent/generated/predicate"
	"go.infratographer.com/x/gidx"
)

// LoadBalancerAnnotationQuery is the builder for querying LoadBalancerAnnotation entities.
type LoadBalancerAnnotationQuery struct {
	config
	ctx              *QueryContext
	order            []loadbalancerannotation.OrderOption
	inters           []Interceptor
	predicates       []predicate.LoadBalancerAnnotation
	withLoadBalancer *LoadBalancerQuery
	modifiers        []func(*sql.Selector)
	loadTotal        []func(context.Context, []*LoadBalancerAnnotation) error
	// intermediate query (i.e. traversal path).
	sql  *sql.Selector
	path func(context.Context) (*sql.Selector, error)
}

// Where adds a new predicate for the LoadBalancerAnnotationQuery builder.
func (lbaq *LoadBalancerAnnotationQuery) Where(ps ...predicate.LoadBalancerAnnotation) *LoadBalancerAnnotationQuery {
	lbaq.predicates = append(lbaq.predicates, ps...)
	return lbaq
}

// Limit the number of records to be returned by this query.
func (lbaq *LoadBalancerAnnotationQuery) Limit(limit int) *LoadBalancerAnnotationQuery {
	lbaq.ctx.Limit = &limit
	return lbaq
}

// Offset to start from.
func (lbaq *LoadBalancerAnnotationQuery) Offset(offset int) *LoadBalancerAnnotationQuery {
	lbaq.ctx.Offset = &offset
	return lbaq
}

// Unique configures the query builder to filter duplicate records on query.
// By default, unique is set to true, and can be disabled using this method.
func (lbaq *LoadBalancerAnnotationQuery) Unique(unique bool) *LoadBalancerAnnotationQuery {
	lbaq.ctx.Unique = &unique
	return lbaq
}

// Order specifies how the records should be ordered.
func (lbaq *LoadBalancerAnnotationQuery) Order(o ...loadbalancerannotation.OrderOption) *LoadBalancerAnnotationQuery {
	lbaq.order = append(lbaq.order, o...)
	return lbaq
}

// QueryLoadBalancer chains the current query on the "load_balancer" edge.
func (lbaq *LoadBalancerAnnotationQuery) QueryLoadBalancer() *LoadBalancerQuery {
	query := (&LoadBalancerClient{config: lbaq.config}).Query()
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := lbaq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		selector := lbaq.sqlQuery(ctx)
		if err := selector.Err(); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(loadbalancerannotation.Table, loadbalancerannotation.FieldID, selector),
			sqlgraph.To(loadbalancer.Table, loadbalancer.FieldID),
			sqlgraph.Edge(sqlgraph.M2O, false, loadbalancerannotation.LoadBalancerTable, loadbalancerannotation.LoadBalancerColumn),
		)
		fromU = sqlgraph.SetNeighbors(lbaq.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// First returns the first LoadBalancerAnnotation entity from the query.
// Returns a *NotFoundError when no LoadBalancerAnnotation was found.
func (lbaq *LoadBalancerAnnotationQuery) First(ctx context.Context) (*LoadBalancerAnnotation, error) {
	nodes, err := lbaq.Limit(1).All(setContextOp(ctx, lbaq.ctx, "First"))
	if err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nil, &NotFoundError{loadbalancerannotation.Label}
	}
	return nodes[0], nil
}

// FirstX is like First, but panics if an error occurs.
func (lbaq *LoadBalancerAnnotationQuery) FirstX(ctx context.Context) *LoadBalancerAnnotation {
	node, err := lbaq.First(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return node
}

// FirstID returns the first LoadBalancerAnnotation ID from the query.
// Returns a *NotFoundError when no LoadBalancerAnnotation ID was found.
func (lbaq *LoadBalancerAnnotationQuery) FirstID(ctx context.Context) (id gidx.PrefixedID, err error) {
	var ids []gidx.PrefixedID
	if ids, err = lbaq.Limit(1).IDs(setContextOp(ctx, lbaq.ctx, "FirstID")); err != nil {
		return
	}
	if len(ids) == 0 {
		err = &NotFoundError{loadbalancerannotation.Label}
		return
	}
	return ids[0], nil
}

// FirstIDX is like FirstID, but panics if an error occurs.
func (lbaq *LoadBalancerAnnotationQuery) FirstIDX(ctx context.Context) gidx.PrefixedID {
	id, err := lbaq.FirstID(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return id
}

// Only returns a single LoadBalancerAnnotation entity found by the query, ensuring it only returns one.
// Returns a *NotSingularError when more than one LoadBalancerAnnotation entity is found.
// Returns a *NotFoundError when no LoadBalancerAnnotation entities are found.
func (lbaq *LoadBalancerAnnotationQuery) Only(ctx context.Context) (*LoadBalancerAnnotation, error) {
	nodes, err := lbaq.Limit(2).All(setContextOp(ctx, lbaq.ctx, "Only"))
	if err != nil {
		return nil, err
	}
	switch len(nodes) {
	case 1:
		return nodes[0], nil
	case 0:
		return nil, &NotFoundError{loadbalancerannotation.Label}
	default:
		return nil, &NotSingularError{loadbalancerannotation.Label}
	}
}

// OnlyX is like Only, but panics if an error occurs.
func (lbaq *LoadBalancerAnnotationQuery) OnlyX(ctx context.Context) *LoadBalancerAnnotation {
	node, err := lbaq.Only(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// OnlyID is like Only, but returns the only LoadBalancerAnnotation ID in the query.
// Returns a *NotSingularError when more than one LoadBalancerAnnotation ID is found.
// Returns a *NotFoundError when no entities are found.
func (lbaq *LoadBalancerAnnotationQuery) OnlyID(ctx context.Context) (id gidx.PrefixedID, err error) {
	var ids []gidx.PrefixedID
	if ids, err = lbaq.Limit(2).IDs(setContextOp(ctx, lbaq.ctx, "OnlyID")); err != nil {
		return
	}
	switch len(ids) {
	case 1:
		id = ids[0]
	case 0:
		err = &NotFoundError{loadbalancerannotation.Label}
	default:
		err = &NotSingularError{loadbalancerannotation.Label}
	}
	return
}

// OnlyIDX is like OnlyID, but panics if an error occurs.
func (lbaq *LoadBalancerAnnotationQuery) OnlyIDX(ctx context.Context) gidx.PrefixedID {
	id, err := lbaq.OnlyID(ctx)
	if err != nil {
		panic(err)
	}
	return id
}

// All executes the query and returns a list of LoadBalancerAnnotations.
func (lbaq *LoadBalancerAnnotationQuery) All(ctx context.Context) ([]*LoadBalancerAnnotation, error) {
	ctx = setContextOp(ctx, lbaq.ctx, "All")
	if err := lbaq.prepareQuery(ctx); err != nil {
		return nil, err
	}
	qr := querierAll[[]*LoadBalancerAnnotation, *LoadBalancerAnnotationQuery]()
	return withInterceptors[[]*LoadBalancerAnnotation](ctx, lbaq, qr, lbaq.inters)
}

// AllX is like All, but panics if an error occurs.
func (lbaq *LoadBalancerAnnotationQuery) AllX(ctx context.Context) []*LoadBalancerAnnotation {
	nodes, err := lbaq.All(ctx)
	if err != nil {
		panic(err)
	}
	return nodes
}

// IDs executes the query and returns a list of LoadBalancerAnnotation IDs.
func (lbaq *LoadBalancerAnnotationQuery) IDs(ctx context.Context) (ids []gidx.PrefixedID, err error) {
	if lbaq.ctx.Unique == nil && lbaq.path != nil {
		lbaq.Unique(true)
	}
	ctx = setContextOp(ctx, lbaq.ctx, "IDs")
	if err = lbaq.Select(loadbalancerannotation.FieldID).Scan(ctx, &ids); err != nil {
		return nil, err
	}
	return ids, nil
}

// IDsX is like IDs, but panics if an error occurs.
func (lbaq *LoadBalancerAnnotationQuery) IDsX(ctx context.Context) []gidx.PrefixedID {
	ids, err := lbaq.IDs(ctx)
	if err != nil {
		panic(err)
	}
	return ids
}

// Count returns the count of the given query.
func (lbaq *LoadBalancerAnnotationQuery) Count(ctx context.Context) (int, error) {
	ctx = setContextOp(ctx, lbaq.ctx, "Count")
	if err := lbaq.prepareQuery(ctx); err != nil {
		return 0, err
	}
	return withInterceptors[int](ctx, lbaq, querierCount[*LoadBalancerAnnotationQuery](), lbaq.inters)
}

// CountX is like Count, but panics if an error occurs.
func (lbaq *LoadBalancerAnnotationQuery) CountX(ctx context.Context) int {
	count, err := lbaq.Count(ctx)
	if err != nil {
		panic(err)
	}
	return count
}

// Exist returns true if the query has elements in the graph.
func (lbaq *LoadBalancerAnnotationQuery) Exist(ctx context.Context) (bool, error) {
	ctx = setContextOp(ctx, lbaq.ctx, "Exist")
	switch _, err := lbaq.FirstID(ctx); {
	case IsNotFound(err):
		return false, nil
	case err != nil:
		return false, fmt.Errorf("generated: check existence: %w", err)
	default:
		return true, nil
	}
}

// ExistX is like Exist, but panics if an error occurs.
func (lbaq *LoadBalancerAnnotationQuery) ExistX(ctx context.Context) bool {
	exist, err := lbaq.Exist(ctx)
	if err != nil {
		panic(err)
	}
	return exist
}

// Clone returns a duplicate of the LoadBalancerAnnotationQuery builder, including all associated steps. It can be
// used to prepare common query builders and use them differently after the clone is made.
func (lbaq *LoadBalancerAnnotationQuery) Clone() *LoadBalancerAnnotationQuery {
	if lbaq == nil {
		return nil
	}
	return &LoadBalancerAnnotationQuery{
		config:           lbaq.config,
		ctx:              lbaq.ctx.Clone(),
		order:            append([]loadbalancerannotation.OrderOption{}, lbaq.order...),
		inters:           append([]Interceptor{}, lbaq.inters...),
		predicates:       append([]predicate.LoadBalancerAnnotation{}, lbaq.predicates...),
		withLoadBalancer: lbaq.withLoadBalancer.Clone(),
		// clone intermediate query.
		sql:  lbaq.sql.Clone(),
		path: lbaq.path,
	}
}

// WithLoadBalancer tells the query-builder to eager-load the nodes that are connected to
// the "load_balancer" edge. The optional arguments are used to configure the query builder of the edge.
func (lbaq *LoadBalancerAnnotationQuery) WithLoadBalancer(opts ...func(*LoadBalancerQuery)) *LoadBalancerAnnotationQuery {
	query := (&LoadBalancerClient{config: lbaq.config}).Query()
	for _, opt := range opts {
		opt(query)
	}
	lbaq.withLoadBalancer = query
	return lbaq
}

// GroupBy is used to group vertices by one or more fields/columns.
// It is often used with aggregate functions, like: count, max, mean, min, sum.
//
// Example:
//
//	var v []struct {
//		Namespace string `json:"namespace,omitempty"`
//		Count int `json:"count,omitempty"`
//	}
//
//	client.LoadBalancerAnnotation.Query().
//		GroupBy(loadbalancerannotation.FieldNamespace).
//		Aggregate(generated.Count()).
//		Scan(ctx, &v)
func (lbaq *LoadBalancerAnnotationQuery) GroupBy(field string, fields ...string) *LoadBalancerAnnotationGroupBy {
	lbaq.ctx.Fields = append([]string{field}, fields...)
	grbuild := &LoadBalancerAnnotationGroupBy{build: lbaq}
	grbuild.flds = &lbaq.ctx.Fields
	grbuild.label = loadbalancerannotation.Label
	grbuild.scan = grbuild.Scan
	return grbuild
}

// Select allows the selection one or more fields/columns for the given query,
// instead of selecting all fields in the entity.
//
// Example:
//
//	var v []struct {
//		Namespace string `json:"namespace,omitempty"`
//	}
//
//	client.LoadBalancerAnnotation.Query().
//		Select(loadbalancerannotation.FieldNamespace).
//		Scan(ctx, &v)
func (lbaq *LoadBalancerAnnotationQuery) Select(fields ...string) *LoadBalancerAnnotationSelect {
	lbaq.ctx.Fields = append(lbaq.ctx.Fields, fields...)
	sbuild := &LoadBalancerAnnotationSelect{LoadBalancerAnnotationQuery: lbaq}
	sbuild.label = loadbalancerannotation.Label
	sbuild.flds, sbuild.scan = &lbaq.ctx.Fields, sbuild.Scan
	return sbuild
}

// Aggregate returns a LoadBalancerAnnotationSelect configured with the given aggregations.
func (lbaq *LoadBalancerAnnotationQuery) Aggregate(fns ...AggregateFunc) *LoadBalancerAnnotationSelect {
	return lbaq.Select().Aggregate(fns...)
}

func (lbaq *LoadBalancerAnnotationQuery) prepareQuery(ctx context.Context) error {
	for _, inter := range lbaq.inters {
		if inter == nil {
			return fmt.Errorf("generated: uninitialized interceptor (forgotten import generated/runtime?)")
		}
		if trv, ok := inter.(Traverser); ok {
			if err := trv.Traverse(ctx, lbaq); err != nil {
				return err
			}
		}
	}
	for _, f := range lbaq.ctx.Fields {
		if !loadbalancerannotation.ValidColumn(f) {
			return &ValidationError{Name: f, err: fmt.Errorf("generated: invalid field %q for query", f)}
		}
	}
	if lbaq.path != nil {
		prev, err := lbaq.path(ctx)
		if err != nil {
			return err
		}
		lbaq.sql = prev
	}
	return nil
}

func (lbaq *LoadBalancerAnnotationQuery) sqlAll(ctx context.Context, hooks ...queryHook) ([]*LoadBalancerAnnotation, error) {
	var (
		nodes       = []*LoadBalancerAnnotation{}
		_spec       = lbaq.querySpec()
		loadedTypes = [1]bool{
			lbaq.withLoadBalancer != nil,
		}
	)
	_spec.ScanValues = func(columns []string) ([]any, error) {
		return (*LoadBalancerAnnotation).scanValues(nil, columns)
	}
	_spec.Assign = func(columns []string, values []any) error {
		node := &LoadBalancerAnnotation{config: lbaq.config}
		nodes = append(nodes, node)
		node.Edges.loadedTypes = loadedTypes
		return node.assignValues(columns, values)
	}
	if len(lbaq.modifiers) > 0 {
		_spec.Modifiers = lbaq.modifiers
	}
	for i := range hooks {
		hooks[i](ctx, _spec)
	}
	if err := sqlgraph.QueryNodes(ctx, lbaq.driver, _spec); err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nodes, nil
	}
	if query := lbaq.withLoadBalancer; query != nil {
		if err := lbaq.loadLoadBalancer(ctx, query, nodes, nil,
			func(n *LoadBalancerAnnotation, e *LoadBalancer) { n.Edges.LoadBalancer = e }); err != nil {
			return nil, err
		}
	}
	for i := range lbaq.loadTotal {
		if err := lbaq.loadTotal[i](ctx, nodes); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

func (lbaq *LoadBalancerAnnotationQuery) loadLoadBalancer(ctx context.Context, query *LoadBalancerQuery, nodes []*LoadBalancerAnnotation, init func(*LoadBalancerAnnotation), assign func(*LoadBalancerAnnotation, *LoadBalancer)) error {
	ids := make([]gidx.PrefixedID, 0, len(nodes))
	nodeids := make(map[gidx.PrefixedID][]*LoadBalancerAnnotation)
	for i := range nodes {
		fk := nodes[i].LoadBalancerID
		if _, ok := nodeids[fk]; !ok {
			ids = append(ids, fk)
		}
		nodeids[fk] = append(nodeids[fk], nodes[i])
	}
	if len(ids) == 0 {
		return nil
	}
	query.Where(loadbalancer.IDIn(ids...))
	neighbors, err := query.All(ctx)
	if err != nil {
		return err
	}
	for _, n := range neighbors {
		nodes, ok := nodeids[n.ID]
		if !ok {
			return fmt.Errorf(`unexpected foreign-key "load_balancer_id" returned %v`, n.ID)
		}
		for i := range nodes {
			assign(nodes[i], n)
		}
	}
	return nil
}

func (lbaq *LoadBalancerAnnotationQuery) sqlCount(ctx context.Context) (int, error) {
	_spec := lbaq.querySpec()
	if len(lbaq.modifiers) > 0 {
		_spec.Modifiers = lbaq.modifiers
	}
	_spec.Node.Columns = lbaq.ctx.Fields
	if len(lbaq.ctx.Fields) > 0 {
		_spec.Unique = lbaq.ctx.Unique != nil && *lbaq.ctx.Unique
	}
	return sqlgraph.CountNodes(ctx, lbaq.driver, _spec)
}

func (lbaq *LoadBalancerAnnotationQuery) querySpec() *sqlgraph.QuerySpec {
	_spec := sqlgraph.NewQuerySpec(loadbalancerannotation.Table, loadbalancerannotation.Columns, sqlgraph.NewFieldSpec(loadbalancerannotation.FieldID, field.TypeString))
	_spec.From = lbaq.sql
	if unique := lbaq.ctx.Unique; unique != nil {
		_spec.Unique = *unique
	} else if lbaq.path != nil {
		_spec.Unique = true
	}
	if fields := lbaq.ctx.Fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, loadbalancerannotation.FieldID)
		for i := range fields {
			if fields[i] != loadbalancerannotation.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, fields[i])
			}
		}
		if lbaq.withLoadBalancer != nil {
			_spec.Node.AddColumnOnce(loadbalancerannotation.FieldLoadBalancerID)
		}
	}
	if ps := lbaq.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if limit := lbaq.ctx.Limit; limit != nil {
		_spec.Limit = *limit
	}
	if offset := lbaq.ctx.Offset; offset != nil {
		_spec.Offset = *offset
	}
	if ps := lbaq.order; len(ps) > 0 {
		_spec.Order = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	return _spec
}

func (lbaq *LoadBalancerAnnotationQuery) sqlQuery(ctx context.Context) *sql.Selector {
	builder := sql.Dialect(lbaq.driver.Dialect())
	t1 := builder.Table(loadbalancerannotation.Table)
	columns := lbaq.ctx.Fields
	if len(columns) == 0 {
		columns = loadbalancerannotation.Columns
	}
	selector := builder.Select(t1.Columns(columns...)...).From(t1)
	if lbaq.sql != nil {
		selector = lbaq.sql
		selector.Select(selector.Columns(columns...)...)
	}
	if lbaq.ctx.Unique != nil && *lbaq.ctx.Unique {
		selector.Distinct()
	}
	for _, p := range lbaq.predicates {
		p(selector)
	}
	for _, p := range lbaq.order {
		p(selector)
	}
	if offset := lbaq.ctx.Offset; offset != nil {
		// limit is mandatory for offset clause. We start
		// with default value, and override it below if needed.
		selector.Offset(*offset).Limit(math.MaxInt32)
	}
	if limit := lbaq.ctx.Limit; limit != nil {
		selector.Limit(*limit)
	}
	return selector
}

// LoadBalancerAnnotationGroupBy is the group-by builder for LoadBalancerAnnotation entities.
type LoadBalancerAnnotationGroupBy struct {
	selector
	build *LoadBalancerAnnotationQuery
}

// Aggregate adds the given aggregation functions to the group-by query.
func (lbagb *LoadBalancerAnnotationGroupBy) Aggregate(fns ...AggregateFunc) *LoadBalancerAnnotationGroupBy {
	lbagb.fns = append(lbagb.fns, fns...)
	return lbagb
}

// Scan applies the selector query and scans the result into the given value.
func (lbagb *LoadBalancerAnnotationGroupBy) Scan(ctx context.Context, v any) error {
	ctx = setContextOp(ctx, lbagb.build.ctx, "GroupBy")
	if err := lbagb.build.prepareQuery(ctx); err != nil {
		return err
	}
	return scanWithInterceptors[*LoadBalancerAnnotationQuery, *LoadBalancerAnnotationGroupBy](ctx, lbagb.build, lbagb, lbagb.build.inters, v)
}

func (lbagb *LoadBalancerAnnotationGroupBy) sqlScan(ctx context.Context, root *LoadBalancerAnnotationQuery, v any) error {
	selector := root.sqlQuery(ctx).Select()
	aggregation := make([]string, 0, len(lbagb.fns))
	for _, fn := range lbagb.fns {
		aggregation = append(aggregation, fn(selector))
	}
	if len(selector.SelectedColumns()) == 0 {
		columns := make([]string, 0, len(*lbagb.flds)+len(lbagb.fns))
		for _, f := range *lbagb.flds {
			columns = append(columns, selector.C(f))
		}
		columns = append(columns, aggregation...)
		selector.Select(columns...)
	}
	selector.GroupBy(selector.Columns(*lbagb.flds...)...)
	if err := selector.Err(); err != nil {
		return err
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err := lbagb.build.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}

// LoadBalancerAnnotationSelect is the builder for selecting fields of LoadBalancerAnnotation entities.
type LoadBalancerAnnotationSelect struct {
	*LoadBalancerAnnotationQuery
	selector
}

// Aggregate adds the given aggregation functions to the selector query.
func (lbas *LoadBalancerAnnotationSelect) Aggregate(fns ...AggregateFunc) *LoadBalancerAnnotationSelect {
	lbas.fns = append(lbas.fns, fns...)
	return lbas
}

// Scan applies the selector query and scans the result into the given value.
func (lbas *LoadBalancerAnnotationSelect) Scan(ctx context.Context, v any) error {
	ctx = setContextOp(ctx, lbas.ctx, "Select")
	if err := lbas.prepareQuery(ctx); err != nil {
		return err
	}
	return scanWithInterceptors[*LoadBalancerAnnotationQuery, *LoadBalancerAnnotationSelect](ctx, lbas.LoadBalancerAnnotationQuery, lbas, lbas.inters, v)
}

func (lbas *LoadBalancerAnnotationSelect) sqlScan(ctx context.Context, root *LoadBalancerAnnotationQuery, v any) error {
	selector := root.sqlQuery(ctx)
	aggregation := make([]string, 0, len(lbas.fns))
	for _, fn := range lbas.fns {
		aggregation = append(aggregation, fn(selector))
	}
	switch n := len(*lbas.selector.flds); {
	case n == 0 && len(aggregation) > 0:
		selector.Select(aggregation...)
	case n != 0 && len(aggregation) > 0:
		selector.AppendSelect(aggregation...)
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err := lbas.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}
