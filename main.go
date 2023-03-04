package main

import (
	"context"
	"fmt"
	"log"
	"net/http"

	"github.com/ServiceWeaver/weaver"
	"github.com/eminetto/microservices-serviceweaver/auth"
	"github.com/eminetto/microservices-serviceweaver/feedback"
	"github.com/eminetto/microservices-serviceweaver/vote"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

func main() {
	root := weaver.Init(context.Background())
	opts := weaver.ListenerOptions{LocalAddress: "localhost:12345"}
	lis, err := root.Listener("talk-manager", opts)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("listener available on %v\n", lis)

	fdb, err := weaver.Get[feedback.FeedbackComponent](root)
	if err != nil {
		log.Fatal(err)
	}

	vt, err := weaver.Get[vote.VoteComponent](root)
	if err != nil {
		log.Fatal(err)
	}

	us, err := weaver.Get[auth.AuthComponent](root)
	if err != nil {
		log.Fatal(err)
	}

	authMiddleware := func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			email, err := us.ValidateToken(r.Context(), r.Header.Get("Authorization"))
			if err != nil {
				w.WriteHeader(http.StatusForbidden)
				return
			}
			ctx := context.WithValue(r.Context(), "email", email)
			next.ServeHTTP(w, r.WithContext(ctx))
		})
	}

	r := chi.NewRouter()
	r.Use(middleware.Logger)
	r.Get("/health", auth.HttpHealth(us))
	r.Post("/auth", auth.HttpAuth(us))
	r.Route("/", func(r chi.Router) {
		r.Use(authMiddleware)
		r.Post("/feedback", feedback.HttpAuth(fdb))
		r.Post("/vote", vote.HttpVote(vt))
	})

	http.Serve(lis, r)
}
