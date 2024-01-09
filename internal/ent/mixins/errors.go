package mixin

import (
	"errors"
	"fmt"

	"entgo.io/ent"
)

var (
	// ErrUnexpectedMutationType is returned when an unexpected mutation type is encountered
	ErrUnexpectedMutationType = errors.New("unexpected mutation type")
)

func newUnexpectedMutationTypeError(m ent.Mutation) error {
	return fmt.Errorf("%w: %s", ErrUnexpectedMutationType, m)
}
