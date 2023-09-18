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

package pool

import (
	"fmt"
	"io"
	"strconv"
	"time"

	"entgo.io/ent"
	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"go.infratographer.com/x/gidx"
)

const (
	// Label holds the string label denoting the pool type in the database.
	Label = "pool"
	// FieldID holds the string denoting the id field in the database.
	FieldID = "id"
	// FieldCreatedAt holds the string denoting the created_at field in the database.
	FieldCreatedAt = "created_at"
	// FieldUpdatedAt holds the string denoting the updated_at field in the database.
	FieldUpdatedAt = "updated_at"
	// FieldDeletedAt holds the string denoting the deleted_at field in the database.
	FieldDeletedAt = "deleted_at"
	// FieldName holds the string denoting the name field in the database.
	FieldName = "name"
	// FieldProtocol holds the string denoting the protocol field in the database.
	FieldProtocol = "protocol"
	// FieldOwnerID holds the string denoting the owner_id field in the database.
	FieldOwnerID = "owner_id"
	// EdgePorts holds the string denoting the ports edge name in mutations.
	EdgePorts = "ports"
	// EdgeOrigins holds the string denoting the origins edge name in mutations.
	EdgeOrigins = "origins"
	// Table holds the table name of the pool in the database.
	Table = "pools"
	// PortsTable is the table that holds the ports relation/edge. The primary key declared below.
	PortsTable = "pool_ports"
	// PortsInverseTable is the table name for the Port entity.
	// It exists in this package in order to avoid circular dependency with the "port" package.
	PortsInverseTable = "ports"
	// OriginsTable is the table that holds the origins relation/edge.
	OriginsTable = "origins"
	// OriginsInverseTable is the table name for the Origin entity.
	// It exists in this package in order to avoid circular dependency with the "origin" package.
	OriginsInverseTable = "origins"
	// OriginsColumn is the table column denoting the origins relation/edge.
	OriginsColumn = "pool_id"
)

// Columns holds all SQL columns for pool fields.
var Columns = []string{
	FieldID,
	FieldCreatedAt,
	FieldUpdatedAt,
	FieldDeletedAt,
	FieldName,
	FieldProtocol,
	FieldOwnerID,
}

var (
	// PortsPrimaryKey and PortsColumn2 are the table columns denoting the
	// primary key for the ports relation (M2M).
	PortsPrimaryKey = []string{"pool_id", "port_id"}
)

// ValidColumn reports if the column name is valid (part of the table columns).
func ValidColumn(column string) bool {
	for i := range Columns {
		if column == Columns[i] {
			return true
		}
	}
	return false
}

// Note that the variables below are initialized by the runtime
// package on the initialization of the application. Therefore,
// it should be imported in the main as follows:
//
//	import _ "go.infratographer.com/load-balancer-api/internal/ent/generated/runtime"
var (
	Hooks        [1]ent.Hook
	Interceptors [1]ent.Interceptor
	// DefaultCreatedAt holds the default value on creation for the "created_at" field.
	DefaultCreatedAt func() time.Time
	// DefaultUpdatedAt holds the default value on creation for the "updated_at" field.
	DefaultUpdatedAt func() time.Time
	// UpdateDefaultUpdatedAt holds the default value on update for the "updated_at" field.
	UpdateDefaultUpdatedAt func() time.Time
	// NameValidator is a validator for the "name" field. It is called by the builders before save.
	NameValidator func(string) error
	// OwnerIDValidator is a validator for the "owner_id" field. It is called by the builders before save.
	OwnerIDValidator func(string) error
	// DefaultID holds the default value on creation for the "id" field.
	DefaultID func() gidx.PrefixedID
)

// Protocol defines the type for the "protocol" enum field.
type Protocol string

// Protocol values.
const (
	ProtocolTCP Protocol = "tcp"
	ProtocolUDP Protocol = "udp"
)

func (pr Protocol) String() string {
	return string(pr)
}

// ProtocolValidator is a validator for the "protocol" field enum values. It is called by the builders before save.
func ProtocolValidator(pr Protocol) error {
	switch pr {
	case ProtocolTCP, ProtocolUDP:
		return nil
	default:
		return fmt.Errorf("pool: invalid enum value for protocol field: %q", pr)
	}
}

// OrderOption defines the ordering options for the Pool queries.
type OrderOption func(*sql.Selector)

// ByID orders the results by the id field.
func ByID(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldID, opts...).ToFunc()
}

// ByCreatedAt orders the results by the created_at field.
func ByCreatedAt(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldCreatedAt, opts...).ToFunc()
}

// ByUpdatedAt orders the results by the updated_at field.
func ByUpdatedAt(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldUpdatedAt, opts...).ToFunc()
}

// ByDeletedAt orders the results by the deleted_at field.
func ByDeletedAt(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldDeletedAt, opts...).ToFunc()
}

// ByName orders the results by the name field.
func ByName(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldName, opts...).ToFunc()
}

// ByProtocol orders the results by the protocol field.
func ByProtocol(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldProtocol, opts...).ToFunc()
}

// ByOwnerID orders the results by the owner_id field.
func ByOwnerID(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldOwnerID, opts...).ToFunc()
}

// ByPortsCount orders the results by ports count.
func ByPortsCount(opts ...sql.OrderTermOption) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborsCount(s, newPortsStep(), opts...)
	}
}

// ByPorts orders the results by ports terms.
func ByPorts(term sql.OrderTerm, terms ...sql.OrderTerm) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborTerms(s, newPortsStep(), append([]sql.OrderTerm{term}, terms...)...)
	}
}

// ByOriginsCount orders the results by origins count.
func ByOriginsCount(opts ...sql.OrderTermOption) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborsCount(s, newOriginsStep(), opts...)
	}
}

// ByOrigins orders the results by origins terms.
func ByOrigins(term sql.OrderTerm, terms ...sql.OrderTerm) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborTerms(s, newOriginsStep(), append([]sql.OrderTerm{term}, terms...)...)
	}
}
func newPortsStep() *sqlgraph.Step {
	return sqlgraph.NewStep(
		sqlgraph.From(Table, FieldID),
		sqlgraph.To(PortsInverseTable, FieldID),
		sqlgraph.Edge(sqlgraph.M2M, false, PortsTable, PortsPrimaryKey...),
	)
}
func newOriginsStep() *sqlgraph.Step {
	return sqlgraph.NewStep(
		sqlgraph.From(Table, FieldID),
		sqlgraph.To(OriginsInverseTable, FieldID),
		sqlgraph.Edge(sqlgraph.O2M, true, OriginsTable, OriginsColumn),
	)
}

// MarshalGQL implements graphql.Marshaler interface.
func (e Protocol) MarshalGQL(w io.Writer) {
	io.WriteString(w, strconv.Quote(e.String()))
}

// UnmarshalGQL implements graphql.Unmarshaler interface.
func (e *Protocol) UnmarshalGQL(val interface{}) error {
	str, ok := val.(string)
	if !ok {
		return fmt.Errorf("enum %T must be a string", val)
	}
	*e = Protocol(str)
	if err := ProtocolValidator(*e); err != nil {
		return fmt.Errorf("%s is not a valid Protocol", str)
	}
	return nil
}
