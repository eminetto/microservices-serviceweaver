package vote

import (
	"context"

	"github.com/ServiceWeaver/weaver"
	"github.com/google/uuid"
)

// Writer is the interface that provides write methods.
type Writer interface {
	Write(ctx context.Context, v *Vote) (uuid.UUID, error)
}

type writer struct {
	weaver.Implements[Writer]
}

func (s *writer) Write(ctx context.Context, v *Vote) (uuid.UUID, error) {
	//@TODO create store rules, using databases or something else
	return uuid.New(), nil
}
