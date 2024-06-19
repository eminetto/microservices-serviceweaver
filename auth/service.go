package auth

import (
	"context"
	"fmt"

	"github.com/ServiceWeaver/weaver"
	"github.com/eminetto/microservices-serviceweaver/auth/security"
)

// Auth is the interface that provides auth methods.
type Auth interface {
	ValidateUser(ctx context.Context, email, password string) error
	GenerateToken(ctx context.Context, email string) (string, error)
	ValidateToken(ctx context.Context, token string) (string, error)
	Health(ctx context.Context) (string, error)
}

type auth struct {
	weaver.Implements[Auth]
}

func (s *auth) Health(ctx context.Context) (string, error) {
	return "ok", nil
}

func (s *auth) ValidateUser(ctx context.Context, email, password string) error {
	//@TODO create validation rules, using databases or something else
	if email == "eminetto@gmail.com" && password != "1234567" {
		return fmt.Errorf("Invalid user")
	}
	return nil
}

func (s *auth) GenerateToken(ctx context.Context, email string) (string, error) {
	return security.NewToken(email)
}

func (s *auth) ValidateToken(ctx context.Context, token string) (string, error) {
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
