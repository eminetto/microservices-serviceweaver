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
	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
)

func main() {
	if err := weaver.Run(context.Background(), serve); err != nil {
		log.Fatal(err)
	}
}

// app is the main component of the application. weaver.Run creates
// it and passes it to serve.
type app struct {
	weaver.Implements[weaver.Main]
	feedback weaver.Ref[feedback.Writer]
	vote     weaver.Ref[vote.Writer]
	auth     weaver.Ref[auth.Auth]
	api      weaver.Listener `weaver:"api"`
}

// serve is called by weaver.Run and contains the body of the application.
func serve(ctx context.Context, app *app) error {
	var fdb feedback.Writer = app.feedback.Get()
	var vt vote.Writer = app.vote.Get()
	var us auth.Auth = app.auth.Get()

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
	r.Get("/health", auth.HealthHandler(us))
	r.Post("/auth", auth.Handler(us))
	r.Route("/", func(r chi.Router) {
		r.Use(authMiddleware)
		r.Post("/feedback", feedback.WriteHandler(fdb))
		r.Post("/vote", vote.WriterHandler(vt))
	})
	fmt.Printf("listener available on %v\n", app.api)

	http.Serve(app.api, r)
	return nil
}
