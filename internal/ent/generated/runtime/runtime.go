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

package runtime

import (
	"time"

	"go.infratographer.com/load-balancer-api/internal/ent/generated/loadbalancer"
	"go.infratographer.com/load-balancer-api/internal/ent/generated/origin"
	"go.infratographer.com/load-balancer-api/internal/ent/generated/pool"
	"go.infratographer.com/load-balancer-api/internal/ent/generated/port"
	"go.infratographer.com/load-balancer-api/internal/ent/generated/provider"
	"go.infratographer.com/load-balancer-api/internal/ent/schema"
	"go.infratographer.com/x/gidx"
)

// The init function reads all schema descriptors with runtime code
// (default values, validators, hooks and policies) and stitches it
// to their package variables.
func init() {
	loadbalancerMixin := schema.LoadBalancer{}.Mixin()
	loadbalancerMixinFields0 := loadbalancerMixin[0].Fields()
	_ = loadbalancerMixinFields0
	loadbalancerFields := schema.LoadBalancer{}.Fields()
	_ = loadbalancerFields
	// loadbalancerDescCreatedAt is the schema descriptor for created_at field.
	loadbalancerDescCreatedAt := loadbalancerMixinFields0[0].Descriptor()
	// loadbalancer.DefaultCreatedAt holds the default value on creation for the created_at field.
	loadbalancer.DefaultCreatedAt = loadbalancerDescCreatedAt.Default.(func() time.Time)
	// loadbalancerDescUpdatedAt is the schema descriptor for updated_at field.
	loadbalancerDescUpdatedAt := loadbalancerMixinFields0[1].Descriptor()
	// loadbalancer.DefaultUpdatedAt holds the default value on creation for the updated_at field.
	loadbalancer.DefaultUpdatedAt = loadbalancerDescUpdatedAt.Default.(func() time.Time)
	// loadbalancer.UpdateDefaultUpdatedAt holds the default value on update for the updated_at field.
	loadbalancer.UpdateDefaultUpdatedAt = loadbalancerDescUpdatedAt.UpdateDefault.(func() time.Time)
	// loadbalancerDescName is the schema descriptor for name field.
	loadbalancerDescName := loadbalancerFields[1].Descriptor()
	// loadbalancer.NameValidator is a validator for the "name" field. It is called by the builders before save.
	loadbalancer.NameValidator = loadbalancerDescName.Validators[0].(func(string) error)
	// loadbalancerDescOwnerID is the schema descriptor for owner_id field.
	loadbalancerDescOwnerID := loadbalancerFields[2].Descriptor()
	// loadbalancer.OwnerIDValidator is a validator for the "owner_id" field. It is called by the builders before save.
	loadbalancer.OwnerIDValidator = loadbalancerDescOwnerID.Validators[0].(func(string) error)
	// loadbalancerDescLocationID is the schema descriptor for location_id field.
	loadbalancerDescLocationID := loadbalancerFields[3].Descriptor()
	// loadbalancer.LocationIDValidator is a validator for the "location_id" field. It is called by the builders before save.
	loadbalancer.LocationIDValidator = loadbalancerDescLocationID.Validators[0].(func(string) error)
	// loadbalancerDescProviderID is the schema descriptor for provider_id field.
	loadbalancerDescProviderID := loadbalancerFields[4].Descriptor()
	// loadbalancer.ProviderIDValidator is a validator for the "provider_id" field. It is called by the builders before save.
	loadbalancer.ProviderIDValidator = loadbalancerDescProviderID.Validators[0].(func(string) error)
	// loadbalancerDescID is the schema descriptor for id field.
	loadbalancerDescID := loadbalancerFields[0].Descriptor()
	// loadbalancer.DefaultID holds the default value on creation for the id field.
	loadbalancer.DefaultID = loadbalancerDescID.Default.(func() gidx.PrefixedID)
	originMixin := schema.Origin{}.Mixin()
	originMixinFields0 := originMixin[0].Fields()
	_ = originMixinFields0
	originFields := schema.Origin{}.Fields()
	_ = originFields
	// originDescCreatedAt is the schema descriptor for created_at field.
	originDescCreatedAt := originMixinFields0[0].Descriptor()
	// origin.DefaultCreatedAt holds the default value on creation for the created_at field.
	origin.DefaultCreatedAt = originDescCreatedAt.Default.(func() time.Time)
	// originDescUpdatedAt is the schema descriptor for updated_at field.
	originDescUpdatedAt := originMixinFields0[1].Descriptor()
	// origin.DefaultUpdatedAt holds the default value on creation for the updated_at field.
	origin.DefaultUpdatedAt = originDescUpdatedAt.Default.(func() time.Time)
	// origin.UpdateDefaultUpdatedAt holds the default value on update for the updated_at field.
	origin.UpdateDefaultUpdatedAt = originDescUpdatedAt.UpdateDefault.(func() time.Time)
	// originDescName is the schema descriptor for name field.
	originDescName := originFields[1].Descriptor()
	// origin.NameValidator is a validator for the "name" field. It is called by the builders before save.
	origin.NameValidator = originDescName.Validators[0].(func(string) error)
	// originDescTarget is the schema descriptor for target field.
	originDescTarget := originFields[2].Descriptor()
	// origin.TargetValidator is a validator for the "target" field. It is called by the builders before save.
	origin.TargetValidator = func() func(string) error {
		validators := originDescTarget.Validators
		fns := [...]func(string) error{
			validators[0].(func(string) error),
			validators[1].(func(string) error),
		}
		return func(target string) error {
			for _, fn := range fns {
				if err := fn(target); err != nil {
					return err
				}
			}
			return nil
		}
	}()
	// originDescPortNumber is the schema descriptor for port_number field.
	originDescPortNumber := originFields[3].Descriptor()
	// origin.PortNumberValidator is a validator for the "port_number" field. It is called by the builders before save.
	origin.PortNumberValidator = func() func(int) error {
		validators := originDescPortNumber.Validators
		fns := [...]func(int) error{
			validators[0].(func(int) error),
			validators[1].(func(int) error),
		}
		return func(port_number int) error {
			for _, fn := range fns {
				if err := fn(port_number); err != nil {
					return err
				}
			}
			return nil
		}
	}()
	// originDescActive is the schema descriptor for active field.
	originDescActive := originFields[4].Descriptor()
	// origin.DefaultActive holds the default value on creation for the active field.
	origin.DefaultActive = originDescActive.Default.(bool)
	// originDescPoolID is the schema descriptor for pool_id field.
	originDescPoolID := originFields[5].Descriptor()
	// origin.PoolIDValidator is a validator for the "pool_id" field. It is called by the builders before save.
	origin.PoolIDValidator = originDescPoolID.Validators[0].(func(string) error)
	// originDescID is the schema descriptor for id field.
	originDescID := originFields[0].Descriptor()
	// origin.DefaultID holds the default value on creation for the id field.
	origin.DefaultID = originDescID.Default.(func() gidx.PrefixedID)
	poolMixin := schema.Pool{}.Mixin()
	poolMixinHooks1 := poolMixin[1].Hooks()
	pool.Hooks[0] = poolMixinHooks1[0]
	poolMixinInters1 := poolMixin[1].Interceptors()
	pool.Interceptors[0] = poolMixinInters1[0]
	poolMixinFields0 := poolMixin[0].Fields()
	_ = poolMixinFields0
	poolFields := schema.Pool{}.Fields()
	_ = poolFields
	// poolDescCreatedAt is the schema descriptor for created_at field.
	poolDescCreatedAt := poolMixinFields0[0].Descriptor()
	// pool.DefaultCreatedAt holds the default value on creation for the created_at field.
	pool.DefaultCreatedAt = poolDescCreatedAt.Default.(func() time.Time)
	// poolDescUpdatedAt is the schema descriptor for updated_at field.
	poolDescUpdatedAt := poolMixinFields0[1].Descriptor()
	// pool.DefaultUpdatedAt holds the default value on creation for the updated_at field.
	pool.DefaultUpdatedAt = poolDescUpdatedAt.Default.(func() time.Time)
	// pool.UpdateDefaultUpdatedAt holds the default value on update for the updated_at field.
	pool.UpdateDefaultUpdatedAt = poolDescUpdatedAt.UpdateDefault.(func() time.Time)
	// poolDescName is the schema descriptor for name field.
	poolDescName := poolFields[1].Descriptor()
	// pool.NameValidator is a validator for the "name" field. It is called by the builders before save.
	pool.NameValidator = poolDescName.Validators[0].(func(string) error)
	// poolDescOwnerID is the schema descriptor for owner_id field.
	poolDescOwnerID := poolFields[3].Descriptor()
	// pool.OwnerIDValidator is a validator for the "owner_id" field. It is called by the builders before save.
	pool.OwnerIDValidator = poolDescOwnerID.Validators[0].(func(string) error)
	// poolDescID is the schema descriptor for id field.
	poolDescID := poolFields[0].Descriptor()
	// pool.DefaultID holds the default value on creation for the id field.
	pool.DefaultID = poolDescID.Default.(func() gidx.PrefixedID)
	portMixin := schema.Port{}.Mixin()
	portMixinFields0 := portMixin[0].Fields()
	_ = portMixinFields0
	portFields := schema.Port{}.Fields()
	_ = portFields
	// portDescCreatedAt is the schema descriptor for created_at field.
	portDescCreatedAt := portMixinFields0[0].Descriptor()
	// port.DefaultCreatedAt holds the default value on creation for the created_at field.
	port.DefaultCreatedAt = portDescCreatedAt.Default.(func() time.Time)
	// portDescUpdatedAt is the schema descriptor for updated_at field.
	portDescUpdatedAt := portMixinFields0[1].Descriptor()
	// port.DefaultUpdatedAt holds the default value on creation for the updated_at field.
	port.DefaultUpdatedAt = portDescUpdatedAt.Default.(func() time.Time)
	// port.UpdateDefaultUpdatedAt holds the default value on update for the updated_at field.
	port.UpdateDefaultUpdatedAt = portDescUpdatedAt.UpdateDefault.(func() time.Time)
	// portDescNumber is the schema descriptor for number field.
	portDescNumber := portFields[1].Descriptor()
	// port.NumberValidator is a validator for the "number" field. It is called by the builders before save.
	port.NumberValidator = func() func(int) error {
		validators := portDescNumber.Validators
		fns := [...]func(int) error{
			validators[0].(func(int) error),
			validators[1].(func(int) error),
			validators[2].(func(int) error),
		}
		return func(number int) error {
			for _, fn := range fns {
				if err := fn(number); err != nil {
					return err
				}
			}
			return nil
		}
	}()
	// portDescName is the schema descriptor for name field.
	portDescName := portFields[2].Descriptor()
	// port.NameValidator is a validator for the "name" field. It is called by the builders before save.
	port.NameValidator = portDescName.Validators[0].(func(string) error)
	// portDescLoadBalancerID is the schema descriptor for load_balancer_id field.
	portDescLoadBalancerID := portFields[3].Descriptor()
	// port.LoadBalancerIDValidator is a validator for the "load_balancer_id" field. It is called by the builders before save.
	port.LoadBalancerIDValidator = portDescLoadBalancerID.Validators[0].(func(string) error)
	// portDescID is the schema descriptor for id field.
	portDescID := portFields[0].Descriptor()
	// port.DefaultID holds the default value on creation for the id field.
	port.DefaultID = portDescID.Default.(func() gidx.PrefixedID)
	providerMixin := schema.Provider{}.Mixin()
	providerMixinFields0 := providerMixin[0].Fields()
	_ = providerMixinFields0
	providerFields := schema.Provider{}.Fields()
	_ = providerFields
	// providerDescCreatedAt is the schema descriptor for created_at field.
	providerDescCreatedAt := providerMixinFields0[0].Descriptor()
	// provider.DefaultCreatedAt holds the default value on creation for the created_at field.
	provider.DefaultCreatedAt = providerDescCreatedAt.Default.(func() time.Time)
	// providerDescUpdatedAt is the schema descriptor for updated_at field.
	providerDescUpdatedAt := providerMixinFields0[1].Descriptor()
	// provider.DefaultUpdatedAt holds the default value on creation for the updated_at field.
	provider.DefaultUpdatedAt = providerDescUpdatedAt.Default.(func() time.Time)
	// provider.UpdateDefaultUpdatedAt holds the default value on update for the updated_at field.
	provider.UpdateDefaultUpdatedAt = providerDescUpdatedAt.UpdateDefault.(func() time.Time)
	// providerDescName is the schema descriptor for name field.
	providerDescName := providerFields[1].Descriptor()
	// provider.NameValidator is a validator for the "name" field. It is called by the builders before save.
	provider.NameValidator = providerDescName.Validators[0].(func(string) error)
	// providerDescOwnerID is the schema descriptor for owner_id field.
	providerDescOwnerID := providerFields[2].Descriptor()
	// provider.OwnerIDValidator is a validator for the "owner_id" field. It is called by the builders before save.
	provider.OwnerIDValidator = providerDescOwnerID.Validators[0].(func(string) error)
	// providerDescID is the schema descriptor for id field.
	providerDescID := providerFields[0].Descriptor()
	// provider.DefaultID holds the default value on creation for the id field.
	provider.DefaultID = providerDescID.Default.(func() gidx.PrefixedID)
}

const (
	Version = "v0.12.4-0.20230503082810-f251400818ea"           // Version of ent codegen.
	Sum     = "h1:R0Rq0neRfHnux+a4NrCgNncloOQQaNAOvEC/YY5+Ox0=" // Sum of ent codegen.
)
