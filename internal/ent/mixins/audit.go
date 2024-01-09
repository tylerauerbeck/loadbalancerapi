package mixin

import (
	"context"
	"time"

	"entgo.io/contrib/entgql"
	"entgo.io/ent"
	"entgo.io/ent/schema/field"
	"entgo.io/ent/schema/index"
	"entgo.io/ent/schema/mixin"
	"go.infratographer.com/x/echojwtx"
)

// AuditMixin defines an interface of a Mixin that provides the created_at and updated_at timestamp fields
type AuditMixin struct {
	mixin.Schema
}

// Fields of the AuditMixin.
func (AuditMixin) Fields() []ent.Field {
	return []ent.Field{
		field.Time("created_by").
			Optional().
			Immutable().
			Annotations(
				entgql.Skip(entgql.SkipMutationCreateInput, entgql.SkipMutationUpdateInput),
				entgql.OrderField("CREATED_BY"),
			),
		field.String("updated_by").
			Optional().
			Annotations(
				entgql.Skip(entgql.SkipMutationCreateInput, entgql.SkipMutationUpdateInput),
				entgql.OrderField("UPDATED_BY"),
			),
	}
}

// Indexes provides indexes on both created_at and updated_at fields
func (AuditMixin) Indexes() []ent.Index {
	return []ent.Index{
		index.Fields("created_by"),
		index.Fields("updated_by"),
	}
}

// Hooks of the AuditMixin.
func (AuditMixin) Hooks() []ent.Hook {
	return []ent.Hook{
		AuditHook,
	}
}

// AuditHook sets and returns the created_at, updated_at, etc., fields
func AuditHook(next ent.Mutator) ent.Mutator {
	type AuditLogger interface {
		SetCreatedAt(time.Time)
		CreatedAt() (v time.Time, exists bool) // exists if present before this hook
		SetUpdatedAt(time.Time)
		UpdatedAt() (v time.Time, exists bool)
		SetCreatedBy(string)
		CreatedBy() (id string, exists bool)
		SetUpdatedBy(string)
		UpdatedBy() (id string, exists bool)
	}

	return ent.MutateFunc(func(ctx context.Context, m ent.Mutation) (ent.Value, error) {
		ml, ok := m.(AuditLogger)
		if !ok {
			return nil, newUnexpectedMutationTypeError(m)
		}

		actor := "unknown"

		id, ok := ctx.Value(echojwtx.ActorCtxKey).(string)
		if ok {
			actor = id
		}

		switch op := m.Op(); {
		case op.Is(ent.OpCreate):
			ml.SetCreatedAt(time.Now())
			ml.SetCreatedBy(actor)
			ml.SetUpdatedBy(actor)

		case op.Is(ent.OpUpdateOne | ent.OpUpdate):
			ml.SetUpdatedAt(time.Now())
			ml.SetUpdatedBy(actor)
		}

		return next.Mutate(ctx, m)
	})
}
