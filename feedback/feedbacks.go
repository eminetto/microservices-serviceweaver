package feedback

import (
	"context"

	"github.com/ServiceWeaver/weaver"
	"github.com/google/uuid"
)

type FeedbackComponent interface {
	Store(ctx context.Context, f Feedback) (uuid.UUID, error)
}

type Service struct {
	weaver.Implements[FeedbackComponent]
}

func (s *Service) Store(ctx context.Context, f Feedback) (uuid.UUID, error) {
	//@TODO create store rules, using databases or something else
	return uuid.New(), nil
}
