package auth

import (
	"context"
	"fmt"

	"github.com/ServiceWeaver/weaver"
	"github.com/eminetto/microservices-serviceweaver/auth/security"
)

type AuthComponent interface {
	ValidateUser(ctx context.Context, email, password string) error
	GenerateToken(ctx context.Context, email string) (string, error)
	ValidateToken(ctx context.Context, token string) (string, error)
}

type Service struct {
	weaver.Implements[AuthComponent]
}

func (s *Service) ValidateUser(ctx context.Context, email, password string) error {
	//@TODO create validation rules, using databases or something else
	if email == "eminetto@gmail.com" && password != "1234567" {
		return fmt.Errorf("Invalid user")
	}
	return nil
}

func (s *Service) GenerateToken(ctx context.Context, email string) (string, error) {
	return security.NewToken(email)
}

func (s *Service) ValidateToken(ctx context.Context, token string) (string, error) {
	t, err := security.ParseToken(token)
	if err != nil {
		return "", err
	}
	tData, err := security.GetClaims(t)
	if err != nil {
		return "", err
	}
	return tData["email"].(string), nil
}
