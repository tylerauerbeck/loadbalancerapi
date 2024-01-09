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
	"database/sql/driver"
	"fmt"

	"entgo.io/contrib/entgql"
	"entgo.io/ent/dialect/sql"
	"github.com/99designs/gqlgen/graphql"
	"go.infratographer.com/load-balancer-api/internal/ent/generated/loadbalancer"
	"go.infratographer.com/load-balancer-api/internal/ent/generated/origin"
	"go.infratographer.com/load-balancer-api/internal/ent/generated/pool"
	"go.infratographer.com/load-balancer-api/internal/ent/generated/port"
	"go.infratographer.com/load-balancer-api/internal/ent/generated/provider"
	"go.infratographer.com/x/gidx"
)

// CollectFields tells the query-builder to eagerly load connected nodes by resolver context.
func (lb *LoadBalancerQuery) CollectFields(ctx context.Context, satisfies ...string) (*LoadBalancerQuery, error) {
	fc := graphql.GetFieldContext(ctx)
	if fc == nil {
		return lb, nil
	}
	if err := lb.collectField(ctx, graphql.GetOperationContext(ctx), fc.Field, nil, satisfies...); err != nil {
		return nil, err
	}
	return lb, nil
}

func (lb *LoadBalancerQuery) collectField(ctx context.Context, opCtx *graphql.OperationContext, collected graphql.CollectedField, path []string, satisfies ...string) error {
	path = append([]string(nil), path...)
	var (
		unknownSeen    bool
		fieldSeen      = make(map[string]struct{}, len(loadbalancer.Columns))
		selectedFields = []string{loadbalancer.FieldID}
	)
	for _, field := range graphql.CollectFields(opCtx, collected.Selections, satisfies) {
		switch field.Name {
		case "ports":
			var (
				alias = field.Alias
				path  = append(path, alias)
				query = (&PortClient{config: lb.config}).Query()
			)
			args := newLoadBalancerPortPaginateArgs(fieldArgs(ctx, new(LoadBalancerPortWhereInput), path...))
			if err := validateFirstLast(args.first, args.last); err != nil {
				return fmt.Errorf("validate first and last in path %q: %w", path, err)
			}
			pager, err := newLoadBalancerPortPager(args.opts, args.last != nil)
			if err != nil {
				return fmt.Errorf("create new pager in path %q: %w", path, err)
			}
			if query, err = pager.applyFilter(query); err != nil {
				return err
			}
			ignoredEdges := !hasCollectedField(ctx, append(path, edgesField)...)
			if hasCollectedField(ctx, append(path, totalCountField)...) || hasCollectedField(ctx, append(path, pageInfoField)...) {
				hasPagination := args.after != nil || args.first != nil || args.before != nil || args.last != nil
				if hasPagination || ignoredEdges {
					query := query.Clone()
					lb.loadTotal = append(lb.loadTotal, func(ctx context.Context, nodes []*LoadBalancer) error {
						ids := make([]driver.Value, len(nodes))
						for i := range nodes {
							ids[i] = nodes[i].ID
						}
						var v []struct {
							NodeID gidx.PrefixedID `sql:"load_balancer_id"`
							Count  int             `sql:"count"`
						}
						query.Where(func(s *sql.Selector) {
							s.Where(sql.InValues(s.C(loadbalancer.PortsColumn), ids...))
						})
						if err := query.GroupBy(loadbalancer.PortsColumn).Aggregate(Count()).Scan(ctx, &v); err != nil {
							return err
						}
						m := make(map[gidx.PrefixedID]int, len(v))
						for i := range v {
							m[v[i].NodeID] = v[i].Count
						}
						for i := range nodes {
							n := m[nodes[i].ID]
							if nodes[i].Edges.totalCount[0] == nil {
								nodes[i].Edges.totalCount[0] = make(map[string]int)
							}
							nodes[i].Edges.totalCount[0][alias] = n
						}
						return nil
					})
				} else {
					lb.loadTotal = append(lb.loadTotal, func(_ context.Context, nodes []*LoadBalancer) error {
						for i := range nodes {
							n := len(nodes[i].Edges.Ports)
							if nodes[i].Edges.totalCount[0] == nil {
								nodes[i].Edges.totalCount[0] = make(map[string]int)
							}
							nodes[i].Edges.totalCount[0][alias] = n
						}
						return nil
					})
				}
			}
			if ignoredEdges || (args.first != nil && *args.first == 0) || (args.last != nil && *args.last == 0) {
				continue
			}
			if query, err = pager.applyCursors(query, args.after, args.before); err != nil {
				return err
			}
			path = append(path, edgesField, nodeField)
			if field := collectedField(ctx, path...); field != nil {
				if err := query.collectField(ctx, opCtx, *field, path, mayAddCondition(satisfies, "Port")...); err != nil {
					return err
				}
			}
			if limit := paginateLimit(args.first, args.last); limit > 0 {
				modify := limitRows(loadbalancer.PortsColumn, limit, pager.orderExpr(query))
				query.modifiers = append(query.modifiers, modify)
			} else {
				query = pager.applyOrder(query)
			}
			lb.WithNamedPorts(alias, func(wq *PortQuery) {
				*wq = *query
			})
		case "loadBalancerProvider":
			var (
				alias = field.Alias
				path  = append(path, alias)
				query = (&ProviderClient{config: lb.config}).Query()
			)
			if err := query.collectField(ctx, opCtx, field, path, satisfies...); err != nil {
				return err
			}
			lb.withProvider = query
			if _, ok := fieldSeen[loadbalancer.FieldProviderID]; !ok {
				selectedFields = append(selectedFields, loadbalancer.FieldProviderID)
				fieldSeen[loadbalancer.FieldProviderID] = struct{}{}
			}
		case "createdAt":
			if _, ok := fieldSeen[loadbalancer.FieldCreatedAt]; !ok {
				selectedFields = append(selectedFields, loadbalancer.FieldCreatedAt)
				fieldSeen[loadbalancer.FieldCreatedAt] = struct{}{}
			}
		case "updatedAt":
			if _, ok := fieldSeen[loadbalancer.FieldUpdatedAt]; !ok {
				selectedFields = append(selectedFields, loadbalancer.FieldUpdatedAt)
				fieldSeen[loadbalancer.FieldUpdatedAt] = struct{}{}
			}
		case "createdBy":
			if _, ok := fieldSeen[loadbalancer.FieldCreatedBy]; !ok {
				selectedFields = append(selectedFields, loadbalancer.FieldCreatedBy)
				fieldSeen[loadbalancer.FieldCreatedBy] = struct{}{}
			}
		case "updatedBy":
			if _, ok := fieldSeen[loadbalancer.FieldUpdatedBy]; !ok {
				selectedFields = append(selectedFields, loadbalancer.FieldUpdatedBy)
				fieldSeen[loadbalancer.FieldUpdatedBy] = struct{}{}
			}
		case "deletedAt":
			if _, ok := fieldSeen[loadbalancer.FieldDeletedAt]; !ok {
				selectedFields = append(selectedFields, loadbalancer.FieldDeletedAt)
				fieldSeen[loadbalancer.FieldDeletedAt] = struct{}{}
			}
		case "deletedBy":
			if _, ok := fieldSeen[loadbalancer.FieldDeletedBy]; !ok {
				selectedFields = append(selectedFields, loadbalancer.FieldDeletedBy)
				fieldSeen[loadbalancer.FieldDeletedBy] = struct{}{}
			}
		case "name":
			if _, ok := fieldSeen[loadbalancer.FieldName]; !ok {
				selectedFields = append(selectedFields, loadbalancer.FieldName)
				fieldSeen[loadbalancer.FieldName] = struct{}{}
			}
		case "id":
		case "__typename":
		default:
			unknownSeen = true
		}
	}
	if !unknownSeen {
		lb.Select(selectedFields...)
	}
	return nil
}

type loadbalancerPaginateArgs struct {
	first, last   *int
	after, before *Cursor
	opts          []LoadBalancerPaginateOption
}

func newLoadBalancerPaginateArgs(rv map[string]any) *loadbalancerPaginateArgs {
	args := &loadbalancerPaginateArgs{}
	if rv == nil {
		return args
	}
	if v := rv[firstField]; v != nil {
		args.first = v.(*int)
	}
	if v := rv[lastField]; v != nil {
		args.last = v.(*int)
	}
	if v := rv[afterField]; v != nil {
		args.after = v.(*Cursor)
	}
	if v := rv[beforeField]; v != nil {
		args.before = v.(*Cursor)
	}
	if v, ok := rv[orderByField]; ok {
		switch v := v.(type) {
		case map[string]any:
			var (
				err1, err2 error
				order      = &LoadBalancerOrder{Field: &LoadBalancerOrderField{}, Direction: entgql.OrderDirectionAsc}
			)
			if d, ok := v[directionField]; ok {
				err1 = order.Direction.UnmarshalGQL(d)
			}
			if f, ok := v[fieldField]; ok {
				err2 = order.Field.UnmarshalGQL(f)
			}
			if err1 == nil && err2 == nil {
				args.opts = append(args.opts, WithLoadBalancerOrder(order))
			}
		case *LoadBalancerOrder:
			if v != nil {
				args.opts = append(args.opts, WithLoadBalancerOrder(v))
			}
		}
	}
	if v, ok := rv[whereField].(*LoadBalancerWhereInput); ok {
		args.opts = append(args.opts, WithLoadBalancerFilter(v.Filter))
	}
	return args
}

// CollectFields tells the query-builder to eagerly load connected nodes by resolver context.
func (o *OriginQuery) CollectFields(ctx context.Context, satisfies ...string) (*OriginQuery, error) {
	fc := graphql.GetFieldContext(ctx)
	if fc == nil {
		return o, nil
	}
	if err := o.collectField(ctx, graphql.GetOperationContext(ctx), fc.Field, nil, satisfies...); err != nil {
		return nil, err
	}
	return o, nil
}

func (o *OriginQuery) collectField(ctx context.Context, opCtx *graphql.OperationContext, collected graphql.CollectedField, path []string, satisfies ...string) error {
	path = append([]string(nil), path...)
	var (
		unknownSeen    bool
		fieldSeen      = make(map[string]struct{}, len(origin.Columns))
		selectedFields = []string{origin.FieldID}
	)
	for _, field := range graphql.CollectFields(opCtx, collected.Selections, satisfies) {
		switch field.Name {
		case "pool":
			var (
				alias = field.Alias
				path  = append(path, alias)
				query = (&PoolClient{config: o.config}).Query()
			)
			if err := query.collectField(ctx, opCtx, field, path, satisfies...); err != nil {
				return err
			}
			o.withPool = query
			if _, ok := fieldSeen[origin.FieldPoolID]; !ok {
				selectedFields = append(selectedFields, origin.FieldPoolID)
				fieldSeen[origin.FieldPoolID] = struct{}{}
			}
		case "createdAt":
			if _, ok := fieldSeen[origin.FieldCreatedAt]; !ok {
				selectedFields = append(selectedFields, origin.FieldCreatedAt)
				fieldSeen[origin.FieldCreatedAt] = struct{}{}
			}
		case "updatedAt":
			if _, ok := fieldSeen[origin.FieldUpdatedAt]; !ok {
				selectedFields = append(selectedFields, origin.FieldUpdatedAt)
				fieldSeen[origin.FieldUpdatedAt] = struct{}{}
			}
		case "name":
			if _, ok := fieldSeen[origin.FieldName]; !ok {
				selectedFields = append(selectedFields, origin.FieldName)
				fieldSeen[origin.FieldName] = struct{}{}
			}
		case "weight":
			if _, ok := fieldSeen[origin.FieldWeight]; !ok {
				selectedFields = append(selectedFields, origin.FieldWeight)
				fieldSeen[origin.FieldWeight] = struct{}{}
			}
		case "target":
			if _, ok := fieldSeen[origin.FieldTarget]; !ok {
				selectedFields = append(selectedFields, origin.FieldTarget)
				fieldSeen[origin.FieldTarget] = struct{}{}
			}
		case "portNumber":
			if _, ok := fieldSeen[origin.FieldPortNumber]; !ok {
				selectedFields = append(selectedFields, origin.FieldPortNumber)
				fieldSeen[origin.FieldPortNumber] = struct{}{}
			}
		case "active":
			if _, ok := fieldSeen[origin.FieldActive]; !ok {
				selectedFields = append(selectedFields, origin.FieldActive)
				fieldSeen[origin.FieldActive] = struct{}{}
			}
		case "poolID":
			if _, ok := fieldSeen[origin.FieldPoolID]; !ok {
				selectedFields = append(selectedFields, origin.FieldPoolID)
				fieldSeen[origin.FieldPoolID] = struct{}{}
			}
		case "id":
		case "__typename":
		default:
			unknownSeen = true
		}
	}
	if !unknownSeen {
		o.Select(selectedFields...)
	}
	return nil
}

type loadbalanceroriginPaginateArgs struct {
	first, last   *int
	after, before *Cursor
	opts          []LoadBalancerOriginPaginateOption
}

func newLoadBalancerOriginPaginateArgs(rv map[string]any) *loadbalanceroriginPaginateArgs {
	args := &loadbalanceroriginPaginateArgs{}
	if rv == nil {
		return args
	}
	if v := rv[firstField]; v != nil {
		args.first = v.(*int)
	}
	if v := rv[lastField]; v != nil {
		args.last = v.(*int)
	}
	if v := rv[afterField]; v != nil {
		args.after = v.(*Cursor)
	}
	if v := rv[beforeField]; v != nil {
		args.before = v.(*Cursor)
	}
	if v, ok := rv[orderByField]; ok {
		switch v := v.(type) {
		case map[string]any:
			var (
				err1, err2 error
				order      = &LoadBalancerOriginOrder{Field: &LoadBalancerOriginOrderField{}, Direction: entgql.OrderDirectionAsc}
			)
			if d, ok := v[directionField]; ok {
				err1 = order.Direction.UnmarshalGQL(d)
			}
			if f, ok := v[fieldField]; ok {
				err2 = order.Field.UnmarshalGQL(f)
			}
			if err1 == nil && err2 == nil {
				args.opts = append(args.opts, WithLoadBalancerOriginOrder(order))
			}
		case *LoadBalancerOriginOrder:
			if v != nil {
				args.opts = append(args.opts, WithLoadBalancerOriginOrder(v))
			}
		}
	}
	if v, ok := rv[whereField].(*LoadBalancerOriginWhereInput); ok {
		args.opts = append(args.opts, WithLoadBalancerOriginFilter(v.Filter))
	}
	return args
}

// CollectFields tells the query-builder to eagerly load connected nodes by resolver context.
func (po *PoolQuery) CollectFields(ctx context.Context, satisfies ...string) (*PoolQuery, error) {
	fc := graphql.GetFieldContext(ctx)
	if fc == nil {
		return po, nil
	}
	if err := po.collectField(ctx, graphql.GetOperationContext(ctx), fc.Field, nil, satisfies...); err != nil {
		return nil, err
	}
	return po, nil
}

func (po *PoolQuery) collectField(ctx context.Context, opCtx *graphql.OperationContext, collected graphql.CollectedField, path []string, satisfies ...string) error {
	path = append([]string(nil), path...)
	var (
		unknownSeen    bool
		fieldSeen      = make(map[string]struct{}, len(pool.Columns))
		selectedFields = []string{pool.FieldID}
	)
	for _, field := range graphql.CollectFields(opCtx, collected.Selections, satisfies) {
		switch field.Name {
		case "ports":
			var (
				alias = field.Alias
				path  = append(path, alias)
				query = (&PortClient{config: po.config}).Query()
			)
			if err := query.collectField(ctx, opCtx, field, path, satisfies...); err != nil {
				return err
			}
			po.WithNamedPorts(alias, func(wq *PortQuery) {
				*wq = *query
			})
		case "origins":
			var (
				alias = field.Alias
				path  = append(path, alias)
				query = (&OriginClient{config: po.config}).Query()
			)
			args := newLoadBalancerOriginPaginateArgs(fieldArgs(ctx, new(LoadBalancerOriginWhereInput), path...))
			if err := validateFirstLast(args.first, args.last); err != nil {
				return fmt.Errorf("validate first and last in path %q: %w", path, err)
			}
			pager, err := newLoadBalancerOriginPager(args.opts, args.last != nil)
			if err != nil {
				return fmt.Errorf("create new pager in path %q: %w", path, err)
			}
			if query, err = pager.applyFilter(query); err != nil {
				return err
			}
			ignoredEdges := !hasCollectedField(ctx, append(path, edgesField)...)
			if hasCollectedField(ctx, append(path, totalCountField)...) || hasCollectedField(ctx, append(path, pageInfoField)...) {
				hasPagination := args.after != nil || args.first != nil || args.before != nil || args.last != nil
				if hasPagination || ignoredEdges {
					query := query.Clone()
					po.loadTotal = append(po.loadTotal, func(ctx context.Context, nodes []*Pool) error {
						ids := make([]driver.Value, len(nodes))
						for i := range nodes {
							ids[i] = nodes[i].ID
						}
						var v []struct {
							NodeID gidx.PrefixedID `sql:"pool_id"`
							Count  int             `sql:"count"`
						}
						query.Where(func(s *sql.Selector) {
							s.Where(sql.InValues(s.C(pool.OriginsColumn), ids...))
						})
						if err := query.GroupBy(pool.OriginsColumn).Aggregate(Count()).Scan(ctx, &v); err != nil {
							return err
						}
						m := make(map[gidx.PrefixedID]int, len(v))
						for i := range v {
							m[v[i].NodeID] = v[i].Count
						}
						for i := range nodes {
							n := m[nodes[i].ID]
							if nodes[i].Edges.totalCount[1] == nil {
								nodes[i].Edges.totalCount[1] = make(map[string]int)
							}
							nodes[i].Edges.totalCount[1][alias] = n
						}
						return nil
					})
				} else {
					po.loadTotal = append(po.loadTotal, func(_ context.Context, nodes []*Pool) error {
						for i := range nodes {
							n := len(nodes[i].Edges.Origins)
							if nodes[i].Edges.totalCount[1] == nil {
								nodes[i].Edges.totalCount[1] = make(map[string]int)
							}
							nodes[i].Edges.totalCount[1][alias] = n
						}
						return nil
					})
				}
			}
			if ignoredEdges || (args.first != nil && *args.first == 0) || (args.last != nil && *args.last == 0) {
				continue
			}
			if query, err = pager.applyCursors(query, args.after, args.before); err != nil {
				return err
			}
			path = append(path, edgesField, nodeField)
			if field := collectedField(ctx, path...); field != nil {
				if err := query.collectField(ctx, opCtx, *field, path, mayAddCondition(satisfies, "Origin")...); err != nil {
					return err
				}
			}
			if limit := paginateLimit(args.first, args.last); limit > 0 {
				modify := limitRows(pool.OriginsColumn, limit, pager.orderExpr(query))
				query.modifiers = append(query.modifiers, modify)
			} else {
				query = pager.applyOrder(query)
			}
			po.WithNamedOrigins(alias, func(wq *OriginQuery) {
				*wq = *query
			})
		case "createdAt":
			if _, ok := fieldSeen[pool.FieldCreatedAt]; !ok {
				selectedFields = append(selectedFields, pool.FieldCreatedAt)
				fieldSeen[pool.FieldCreatedAt] = struct{}{}
			}
		case "updatedAt":
			if _, ok := fieldSeen[pool.FieldUpdatedAt]; !ok {
				selectedFields = append(selectedFields, pool.FieldUpdatedAt)
				fieldSeen[pool.FieldUpdatedAt] = struct{}{}
			}
		case "name":
			if _, ok := fieldSeen[pool.FieldName]; !ok {
				selectedFields = append(selectedFields, pool.FieldName)
				fieldSeen[pool.FieldName] = struct{}{}
			}
		case "protocol":
			if _, ok := fieldSeen[pool.FieldProtocol]; !ok {
				selectedFields = append(selectedFields, pool.FieldProtocol)
				fieldSeen[pool.FieldProtocol] = struct{}{}
			}
		case "ownerID":
			if _, ok := fieldSeen[pool.FieldOwnerID]; !ok {
				selectedFields = append(selectedFields, pool.FieldOwnerID)
				fieldSeen[pool.FieldOwnerID] = struct{}{}
			}
		case "id":
		case "__typename":
		default:
			unknownSeen = true
		}
	}
	if !unknownSeen {
		po.Select(selectedFields...)
	}
	return nil
}

type loadbalancerpoolPaginateArgs struct {
	first, last   *int
	after, before *Cursor
	opts          []LoadBalancerPoolPaginateOption
}

func newLoadBalancerPoolPaginateArgs(rv map[string]any) *loadbalancerpoolPaginateArgs {
	args := &loadbalancerpoolPaginateArgs{}
	if rv == nil {
		return args
	}
	if v := rv[firstField]; v != nil {
		args.first = v.(*int)
	}
	if v := rv[lastField]; v != nil {
		args.last = v.(*int)
	}
	if v := rv[afterField]; v != nil {
		args.after = v.(*Cursor)
	}
	if v := rv[beforeField]; v != nil {
		args.before = v.(*Cursor)
	}
	if v, ok := rv[orderByField]; ok {
		switch v := v.(type) {
		case map[string]any:
			var (
				err1, err2 error
				order      = &LoadBalancerPoolOrder{Field: &LoadBalancerPoolOrderField{}, Direction: entgql.OrderDirectionAsc}
			)
			if d, ok := v[directionField]; ok {
				err1 = order.Direction.UnmarshalGQL(d)
			}
			if f, ok := v[fieldField]; ok {
				err2 = order.Field.UnmarshalGQL(f)
			}
			if err1 == nil && err2 == nil {
				args.opts = append(args.opts, WithLoadBalancerPoolOrder(order))
			}
		case *LoadBalancerPoolOrder:
			if v != nil {
				args.opts = append(args.opts, WithLoadBalancerPoolOrder(v))
			}
		}
	}
	if v, ok := rv[whereField].(*LoadBalancerPoolWhereInput); ok {
		args.opts = append(args.opts, WithLoadBalancerPoolFilter(v.Filter))
	}
	return args
}

// CollectFields tells the query-builder to eagerly load connected nodes by resolver context.
func (po *PortQuery) CollectFields(ctx context.Context, satisfies ...string) (*PortQuery, error) {
	fc := graphql.GetFieldContext(ctx)
	if fc == nil {
		return po, nil
	}
	if err := po.collectField(ctx, graphql.GetOperationContext(ctx), fc.Field, nil, satisfies...); err != nil {
		return nil, err
	}
	return po, nil
}

func (po *PortQuery) collectField(ctx context.Context, opCtx *graphql.OperationContext, collected graphql.CollectedField, path []string, satisfies ...string) error {
	path = append([]string(nil), path...)
	var (
		unknownSeen    bool
		fieldSeen      = make(map[string]struct{}, len(port.Columns))
		selectedFields = []string{port.FieldID}
	)
	for _, field := range graphql.CollectFields(opCtx, collected.Selections, satisfies) {
		switch field.Name {
		case "pools":
			var (
				alias = field.Alias
				path  = append(path, alias)
				query = (&PoolClient{config: po.config}).Query()
			)
			if err := query.collectField(ctx, opCtx, field, path, satisfies...); err != nil {
				return err
			}
			po.WithNamedPools(alias, func(wq *PoolQuery) {
				*wq = *query
			})
		case "loadBalancer":
			var (
				alias = field.Alias
				path  = append(path, alias)
				query = (&LoadBalancerClient{config: po.config}).Query()
			)
			if err := query.collectField(ctx, opCtx, field, path, satisfies...); err != nil {
				return err
			}
			po.withLoadBalancer = query
			if _, ok := fieldSeen[port.FieldLoadBalancerID]; !ok {
				selectedFields = append(selectedFields, port.FieldLoadBalancerID)
				fieldSeen[port.FieldLoadBalancerID] = struct{}{}
			}
		case "createdAt":
			if _, ok := fieldSeen[port.FieldCreatedAt]; !ok {
				selectedFields = append(selectedFields, port.FieldCreatedAt)
				fieldSeen[port.FieldCreatedAt] = struct{}{}
			}
		case "updatedAt":
			if _, ok := fieldSeen[port.FieldUpdatedAt]; !ok {
				selectedFields = append(selectedFields, port.FieldUpdatedAt)
				fieldSeen[port.FieldUpdatedAt] = struct{}{}
			}
		case "number":
			if _, ok := fieldSeen[port.FieldNumber]; !ok {
				selectedFields = append(selectedFields, port.FieldNumber)
				fieldSeen[port.FieldNumber] = struct{}{}
			}
		case "name":
			if _, ok := fieldSeen[port.FieldName]; !ok {
				selectedFields = append(selectedFields, port.FieldName)
				fieldSeen[port.FieldName] = struct{}{}
			}
		case "loadBalancerID":
			if _, ok := fieldSeen[port.FieldLoadBalancerID]; !ok {
				selectedFields = append(selectedFields, port.FieldLoadBalancerID)
				fieldSeen[port.FieldLoadBalancerID] = struct{}{}
			}
		case "id":
		case "__typename":
		default:
			unknownSeen = true
		}
	}
	if !unknownSeen {
		po.Select(selectedFields...)
	}
	return nil
}

type loadbalancerportPaginateArgs struct {
	first, last   *int
	after, before *Cursor
	opts          []LoadBalancerPortPaginateOption
}

func newLoadBalancerPortPaginateArgs(rv map[string]any) *loadbalancerportPaginateArgs {
	args := &loadbalancerportPaginateArgs{}
	if rv == nil {
		return args
	}
	if v := rv[firstField]; v != nil {
		args.first = v.(*int)
	}
	if v := rv[lastField]; v != nil {
		args.last = v.(*int)
	}
	if v := rv[afterField]; v != nil {
		args.after = v.(*Cursor)
	}
	if v := rv[beforeField]; v != nil {
		args.before = v.(*Cursor)
	}
	if v, ok := rv[orderByField]; ok {
		switch v := v.(type) {
		case map[string]any:
			var (
				err1, err2 error
				order      = &LoadBalancerPortOrder{Field: &LoadBalancerPortOrderField{}, Direction: entgql.OrderDirectionAsc}
			)
			if d, ok := v[directionField]; ok {
				err1 = order.Direction.UnmarshalGQL(d)
			}
			if f, ok := v[fieldField]; ok {
				err2 = order.Field.UnmarshalGQL(f)
			}
			if err1 == nil && err2 == nil {
				args.opts = append(args.opts, WithLoadBalancerPortOrder(order))
			}
		case *LoadBalancerPortOrder:
			if v != nil {
				args.opts = append(args.opts, WithLoadBalancerPortOrder(v))
			}
		}
	}
	if v, ok := rv[whereField].(*LoadBalancerPortWhereInput); ok {
		args.opts = append(args.opts, WithLoadBalancerPortFilter(v.Filter))
	}
	return args
}

// CollectFields tells the query-builder to eagerly load connected nodes by resolver context.
func (pr *ProviderQuery) CollectFields(ctx context.Context, satisfies ...string) (*ProviderQuery, error) {
	fc := graphql.GetFieldContext(ctx)
	if fc == nil {
		return pr, nil
	}
	if err := pr.collectField(ctx, graphql.GetOperationContext(ctx), fc.Field, nil, satisfies...); err != nil {
		return nil, err
	}
	return pr, nil
}

func (pr *ProviderQuery) collectField(ctx context.Context, opCtx *graphql.OperationContext, collected graphql.CollectedField, path []string, satisfies ...string) error {
	path = append([]string(nil), path...)
	var (
		unknownSeen    bool
		fieldSeen      = make(map[string]struct{}, len(provider.Columns))
		selectedFields = []string{provider.FieldID}
	)
	for _, field := range graphql.CollectFields(opCtx, collected.Selections, satisfies) {
		switch field.Name {
		case "loadBalancers":
			var (
				alias = field.Alias
				path  = append(path, alias)
				query = (&LoadBalancerClient{config: pr.config}).Query()
			)
			args := newLoadBalancerPaginateArgs(fieldArgs(ctx, new(LoadBalancerWhereInput), path...))
			if err := validateFirstLast(args.first, args.last); err != nil {
				return fmt.Errorf("validate first and last in path %q: %w", path, err)
			}
			pager, err := newLoadBalancerPager(args.opts, args.last != nil)
			if err != nil {
				return fmt.Errorf("create new pager in path %q: %w", path, err)
			}
			if query, err = pager.applyFilter(query); err != nil {
				return err
			}
			ignoredEdges := !hasCollectedField(ctx, append(path, edgesField)...)
			if hasCollectedField(ctx, append(path, totalCountField)...) || hasCollectedField(ctx, append(path, pageInfoField)...) {
				hasPagination := args.after != nil || args.first != nil || args.before != nil || args.last != nil
				if hasPagination || ignoredEdges {
					query := query.Clone()
					pr.loadTotal = append(pr.loadTotal, func(ctx context.Context, nodes []*Provider) error {
						ids := make([]driver.Value, len(nodes))
						for i := range nodes {
							ids[i] = nodes[i].ID
						}
						var v []struct {
							NodeID gidx.PrefixedID `sql:"provider_id"`
							Count  int             `sql:"count"`
						}
						query.Where(func(s *sql.Selector) {
							s.Where(sql.InValues(s.C(provider.LoadBalancersColumn), ids...))
						})
						if err := query.GroupBy(provider.LoadBalancersColumn).Aggregate(Count()).Scan(ctx, &v); err != nil {
							return err
						}
						m := make(map[gidx.PrefixedID]int, len(v))
						for i := range v {
							m[v[i].NodeID] = v[i].Count
						}
						for i := range nodes {
							n := m[nodes[i].ID]
							if nodes[i].Edges.totalCount[0] == nil {
								nodes[i].Edges.totalCount[0] = make(map[string]int)
							}
							nodes[i].Edges.totalCount[0][alias] = n
						}
						return nil
					})
				} else {
					pr.loadTotal = append(pr.loadTotal, func(_ context.Context, nodes []*Provider) error {
						for i := range nodes {
							n := len(nodes[i].Edges.LoadBalancers)
							if nodes[i].Edges.totalCount[0] == nil {
								nodes[i].Edges.totalCount[0] = make(map[string]int)
							}
							nodes[i].Edges.totalCount[0][alias] = n
						}
						return nil
					})
				}
			}
			if ignoredEdges || (args.first != nil && *args.first == 0) || (args.last != nil && *args.last == 0) {
				continue
			}
			if query, err = pager.applyCursors(query, args.after, args.before); err != nil {
				return err
			}
			path = append(path, edgesField, nodeField)
			if field := collectedField(ctx, path...); field != nil {
				if err := query.collectField(ctx, opCtx, *field, path, mayAddCondition(satisfies, "LoadBalancer")...); err != nil {
					return err
				}
			}
			if limit := paginateLimit(args.first, args.last); limit > 0 {
				modify := limitRows(provider.LoadBalancersColumn, limit, pager.orderExpr(query))
				query.modifiers = append(query.modifiers, modify)
			} else {
				query = pager.applyOrder(query)
			}
			pr.WithNamedLoadBalancers(alias, func(wq *LoadBalancerQuery) {
				*wq = *query
			})
		case "createdAt":
			if _, ok := fieldSeen[provider.FieldCreatedAt]; !ok {
				selectedFields = append(selectedFields, provider.FieldCreatedAt)
				fieldSeen[provider.FieldCreatedAt] = struct{}{}
			}
		case "updatedAt":
			if _, ok := fieldSeen[provider.FieldUpdatedAt]; !ok {
				selectedFields = append(selectedFields, provider.FieldUpdatedAt)
				fieldSeen[provider.FieldUpdatedAt] = struct{}{}
			}
		case "name":
			if _, ok := fieldSeen[provider.FieldName]; !ok {
				selectedFields = append(selectedFields, provider.FieldName)
				fieldSeen[provider.FieldName] = struct{}{}
			}
		case "id":
		case "__typename":
		default:
			unknownSeen = true
		}
	}
	if !unknownSeen {
		pr.Select(selectedFields...)
	}
	return nil
}

type loadbalancerproviderPaginateArgs struct {
	first, last   *int
	after, before *Cursor
	opts          []LoadBalancerProviderPaginateOption
}

func newLoadBalancerProviderPaginateArgs(rv map[string]any) *loadbalancerproviderPaginateArgs {
	args := &loadbalancerproviderPaginateArgs{}
	if rv == nil {
		return args
	}
	if v := rv[firstField]; v != nil {
		args.first = v.(*int)
	}
	if v := rv[lastField]; v != nil {
		args.last = v.(*int)
	}
	if v := rv[afterField]; v != nil {
		args.after = v.(*Cursor)
	}
	if v := rv[beforeField]; v != nil {
		args.before = v.(*Cursor)
	}
	if v, ok := rv[orderByField]; ok {
		switch v := v.(type) {
		case map[string]any:
			var (
				err1, err2 error
				order      = &LoadBalancerProviderOrder{Field: &LoadBalancerProviderOrderField{}, Direction: entgql.OrderDirectionAsc}
			)
			if d, ok := v[directionField]; ok {
				err1 = order.Direction.UnmarshalGQL(d)
			}
			if f, ok := v[fieldField]; ok {
				err2 = order.Field.UnmarshalGQL(f)
			}
			if err1 == nil && err2 == nil {
				args.opts = append(args.opts, WithLoadBalancerProviderOrder(order))
			}
		case *LoadBalancerProviderOrder:
			if v != nil {
				args.opts = append(args.opts, WithLoadBalancerProviderOrder(v))
			}
		}
	}
	if v, ok := rv[whereField].(*LoadBalancerProviderWhereInput); ok {
		args.opts = append(args.opts, WithLoadBalancerProviderFilter(v.Filter))
	}
	return args
}

const (
	afterField     = "after"
	firstField     = "first"
	beforeField    = "before"
	lastField      = "last"
	orderByField   = "orderBy"
	directionField = "direction"
	fieldField     = "field"
	whereField     = "where"
)

func fieldArgs(ctx context.Context, whereInput any, path ...string) map[string]any {
	field := collectedField(ctx, path...)
	if field == nil || field.Arguments == nil {
		return nil
	}
	oc := graphql.GetOperationContext(ctx)
	args := field.ArgumentMap(oc.Variables)
	return unmarshalArgs(ctx, whereInput, args)
}

// unmarshalArgs allows extracting the field arguments from their raw representation.
func unmarshalArgs(ctx context.Context, whereInput any, args map[string]any) map[string]any {
	for _, k := range []string{firstField, lastField} {
		v, ok := args[k]
		if !ok {
			continue
		}
		i, err := graphql.UnmarshalInt(v)
		if err == nil {
			args[k] = &i
		}
	}
	for _, k := range []string{beforeField, afterField} {
		v, ok := args[k]
		if !ok {
			continue
		}
		c := &Cursor{}
		if c.UnmarshalGQL(v) == nil {
			args[k] = c
		}
	}
	if v, ok := args[whereField]; ok && whereInput != nil {
		if err := graphql.UnmarshalInputFromContext(ctx, v, whereInput); err == nil {
			args[whereField] = whereInput
		}
	}

	return args
}

func limitRows(partitionBy string, limit int, orderBy ...sql.Querier) func(s *sql.Selector) {
	return func(s *sql.Selector) {
		d := sql.Dialect(s.Dialect())
		s.SetDistinct(false)
		with := d.With("src_query").
			As(s.Clone()).
			With("limited_query").
			As(
				d.Select("*").
					AppendSelectExprAs(
						sql.RowNumber().PartitionBy(partitionBy).OrderExpr(orderBy...),
						"row_number",
					).
					From(d.Table("src_query")),
			)
		t := d.Table("limited_query").As(s.TableName())
		*s = *d.Select(s.UnqualifiedColumns()...).
			From(t).
			Where(sql.LTE(t.C("row_number"), limit)).
			Prefix(with)
	}
}

// mayAddCondition appends another type condition to the satisfies list
// if condition is enabled (Node/Nodes) and it does not exist in the list.
func mayAddCondition(satisfies []string, typeCond string) []string {
	if len(satisfies) == 0 {
		return satisfies
	}
	for _, s := range satisfies {
		if typeCond == s {
			return satisfies
		}
	}
	return append(satisfies, typeCond)
}
