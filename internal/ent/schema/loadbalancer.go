package schema

import (
	"context"
	"fmt"

	"entgo.io/contrib/entgql"
	"entgo.io/ent"
	"entgo.io/ent/schema"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	"entgo.io/ent/schema/index"

	gen "go.infratographer.com/load-balancer-api/internal/ent/generated"
	"go.infratographer.com/load-balancer-api/internal/ent/generated/hook"
	"go.infratographer.com/load-balancer-api/internal/pubsub"
	"go.infratographer.com/x/entx"
	"go.infratographer.com/x/gidx"
)

// LoadBalancer holds the schema definition for the LoadBalancer entity.
type LoadBalancer struct {
	ent.Schema
}

// Mixin to use for LoadBalancer type
func (LoadBalancer) Mixin() []ent.Mixin {
	return []ent.Mixin{
		entx.NewTimestampMixin(),
		// softdelete.Mixin{},
	}
}

// Fields of the Instance.
func (LoadBalancer) Fields() []ent.Field {
	return []ent.Field{
		field.String("id").
			GoType(gidx.PrefixedID("")).
			DefaultFunc(func() gidx.PrefixedID { return gidx.MustNewID(LoadBalancerPrefix) }).
			Unique().
			Immutable().
			Comment("The ID for the load balancer.").
			Annotations(
				entgql.OrderField("ID"),
			),
		field.Text("name").
			NotEmpty().
			Comment("The name of the load balancer.").
			Annotations(
				entgql.OrderField("NAME"),
			),
		field.String("tenant_id").
			GoType(gidx.PrefixedID("")).
			Immutable().
			Comment("The ID for the tenant for this load balancer.").
			Annotations(
				entgql.QueryField(),
				entgql.Type("ID"),
				entgql.Skip(entgql.SkipWhereInput, entgql.SkipMutationUpdateInput, entgql.SkipType),
				entgql.OrderField("TENANT"),
			),
		field.String("location_id").
			GoType(gidx.PrefixedID("")).
			Immutable().
			NotEmpty().
			Comment("The ID for the location of this load balancer.").
			Annotations(
				entgql.Type("ID"),
				entgql.Skip(^entgql.SkipMutationCreateInput),
			),
		field.String("provider_id").
			GoType(gidx.PrefixedID("")).
			Immutable().
			NotEmpty().
			Comment("The ID for the load balancer provider for this load balancer.").
			Annotations(
				entgql.Type("ID"),
				entgql.Skip(^entgql.SkipMutationCreateInput),
			),
	}
}

// Edges of the Instance.
func (LoadBalancer) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("annotations", LoadBalancerAnnotation.Type).
			Ref("load_balancer").
			Comment("Annotations for the load balancer.").
			Annotations(
				entgql.RelayConnection(),
				entgql.Skip(entgql.SkipMutationCreateInput, entgql.SkipMutationUpdateInput),
			),
		edge.From("statuses", LoadBalancerStatus.Type).
			Ref("load_balancer").
			Comment("Statuses for the load balancer.").
			Annotations(
				entgql.RelayConnection(),
				entgql.Skip(entgql.SkipMutationCreateInput, entgql.SkipMutationUpdateInput),
			),
		edge.From("ports", Port.Type).
			Ref("load_balancer").
			Annotations(
				entgql.RelayConnection(),
			),
		edge.To("provider", Provider.Type).
			Unique().
			Required().
			Immutable().
			Field("provider_id").
			Comment("The load balancer provider for the load balancer.").
			Annotations(
				entgql.MapsTo("loadBalancerProvider"),
			),
	}
}

// Indexes of the LoadBalancer
func (LoadBalancer) Indexes() []ent.Index {
	return []ent.Index{
		index.Fields("provider_id"),
		index.Fields("location_id"),
		index.Fields("tenant_id"),
	}
}

// Annotations for the LoadBalancer
func (LoadBalancer) Annotations() []schema.Annotation {
	return []schema.Annotation{
		entx.GraphKeyDirective("id"),
		schema.Comment("Representation of a load balancer."),
		entgql.Implements("IPv4Addressable"),
		entgql.RelayConnection(),
		entgql.Mutations(
			entgql.MutationCreate().Description("Input information to create a load balancer."),
			entgql.MutationUpdate().Description("Input information to update a load balancer."),
		),
	}
}

// Hooks might do some things
func (LoadBalancer) Hooks() []ent.Hook {
	return []ent.Hook{
		hook.On(
			func(next ent.Mutator) ent.Mutator {
				return hook.LoadBalancerFunc(func(ctx context.Context, m *gen.LoadBalancerMutation) (ent.Value, error) {
					val, err := next.Mutate(ctx, m)
					if err != nil {
						return nil, err
					}

					fieldMap := make(map[string]ent.Value)
					for _, field := range m.Fields() {
						v, _ := m.Field(field)
						fieldMap[field] = v
					}

					msg, err := pubsub.NewMessage(fieldMap["tenant_id"].(gidx.PrefixedID).String(), pubsub.WithEventType("create"), pubsub.WithSource("load-balancer-api"))

					fmt.Println(msg)

					return val, err
				})
			},
			// Limit the hook only for these operations.
			ent.OpCreate|ent.OpUpdate|ent.OpUpdateOne,
		),
	}
}
