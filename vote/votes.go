package vote

import (
	"context"

	"github.com/ServiceWeaver/weaver"
	"github.com/google/uuid"
)

type VoteComponent interface {
	Store(ctx context.Context, v Vote) (uuid.UUID, error)
}

type Service struct {
	weaver.Implements[VoteComponent]
}

func (s *Service) Store(ctx context.Context, v Vote) (uuid.UUID, error) {
	//@TODO create store rules, using databases or something else
	return uuid.New(), nil
}
