package graphapi

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.
// Code generated by github.com/99designs/gqlgen version v0.17.34

import (
	"context"
	"database/sql"

	"go.infratographer.com/load-balancer-api/internal/ent/generated"
	"go.infratographer.com/load-balancer-api/internal/ent/generated/port"
	"go.infratographer.com/load-balancer-api/internal/ent/generated/predicate"
	"go.infratographer.com/permissions-api/pkg/permissions"
	"go.infratographer.com/x/gidx"
)

// LoadBalancerCreate is the resolver for the loadBalancerCreate field.
func (r *mutationResolver) LoadBalancerCreate(ctx context.Context, input generated.CreateLoadBalancerInput) (*LoadBalancerCreatePayload, error) {
	if err := permissions.CheckAccess(ctx, input.OwnerID, actionLoadBalancerCreate); err != nil {
		return nil, err
	}

	lb, err := r.client.LoadBalancer.Create().SetInput(input).Save(ctx)
	if err != nil {
		return nil, err
	}

	return &LoadBalancerCreatePayload{LoadBalancer: lb}, nil
}

// LoadBalancerUpdate is the resolver for the loadBalancerUpdate field.
func (r *mutationResolver) LoadBalancerUpdate(ctx context.Context, id gidx.PrefixedID, input generated.UpdateLoadBalancerInput) (*LoadBalancerUpdatePayload, error) {
	if err := permissions.CheckAccess(ctx, id, actionLoadBalancerUpdate); err != nil {
		return nil, err
	}

	lb, err := r.client.LoadBalancer.Get(ctx, id)
	if err != nil {
		return nil, err
	}

	lb, err = lb.Update().SetInput(input).Save(ctx)
	if err != nil {
		return nil, err
	}

	return &LoadBalancerUpdatePayload{LoadBalancer: lb}, nil
}

// LoadBalancerDelete is the resolver for the loadBalancerDelete field.
func (r *mutationResolver) LoadBalancerDelete(ctx context.Context, id gidx.PrefixedID) (*LoadBalancerDeletePayload, error) {
	if err := permissions.CheckAccess(ctx, id, actionLoadBalancerDelete); err != nil {
		return nil, err
	}

	tx, err := r.client.BeginTx(ctx, &sql.TxOptions{})
	if err != nil {
		return nil, err
	}

	// cleanup ports associated with loadbalancer
	ports, err := tx.Port.Query().Where(predicate.Port(port.LoadBalancerIDEQ(id))).All(ctx)
	if err != nil {
		r.logger.Errorw("failed to query ports", "error", err)
		if rerr := tx.Rollback(); rerr != nil {
			r.logger.Errorw("failed to rollback transaction", "error", rerr, "stage", "query ports")
		}
		return nil, err
	}

	for _, p := range ports {
		if err = tx.Port.DeleteOne(p).Exec(ctx); err != nil {
			r.logger.Errorw("failed to delete port", "port", p.ID, "error", err)
			if rerr := tx.Rollback(); rerr != nil {
				r.logger.Errorw("failed to rollback transaction", "error", rerr, "stage", "delete port")
			}
			return nil, err
		}
	}

	// delete loadbalancer
	// if err = tx.LoadBalancer.DeleteOneID(id).Exec(ctx); err != nil {
	// 	r.logger.Errorw("failed to delete loadbalancer", "loadbalancer", id.String(), "error", err)
	// 	if rerr := tx.Rollback(); rerr != nil {
	// 		r.logger.Errorw("failed to rollback transaction", "error", rerr, "stage", "delete loadbalancer")
	// 	}
	// 	return nil, err
	// }

	// softdelete loadbalancer
	// if err := tx.LoadBalancer.u

	if err = tx.Commit(); err != nil {
		r.logger.Errorw("failed to commit transaction", "error", err)
		if rerr := tx.Rollback(); rerr != nil {
			r.logger.Errorw("failed to rollback transaction", "error", rerr, "stage", "commit transaction")
		}
		return nil, err
	}

	return &LoadBalancerDeletePayload{DeletedID: id}, nil
}

// LoadBalancer is the resolver for the loadBalancer field.
func (r *queryResolver) LoadBalancer(ctx context.Context, id gidx.PrefixedID) (*generated.LoadBalancer, error) {
	if err := permissions.CheckAccess(ctx, id, actionLoadBalancerGet); err != nil {
		return nil, err
	}

	return r.client.LoadBalancer.Get(ctx, id)
}
