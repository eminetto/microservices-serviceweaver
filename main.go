package main

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/ServiceWeaver/weaver"
	"github.com/eminetto/microservices-serviceweaver/auth"
	"github.com/eminetto/microservices-serviceweaver/feedback"
	"github.com/eminetto/microservices-serviceweaver/vote"
	"github.com/google/uuid"
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

	http.HandleFunc("/auth", func(w http.ResponseWriter, r *http.Request) {
		var param struct {
			Email    string `json:"email"`
			Password string `json:"password"`
		}
		err := json.NewDecoder(r.Body).Decode(&param)
		if err != nil {
			w.WriteHeader(http.StatusBadGateway)
			return
		}
		err = us.ValidateUser(r.Context(), param.Email, param.Password)
		if err != nil {
			w.WriteHeader(http.StatusForbidden)
			return
		}
		var result struct {
			Token string `json:"token"`
		}
		result.Token, err = us.GenerateToken(r.Context(), param.Email)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			return
		}

		if err := json.NewEncoder(w).Encode(result); err != nil {
			w.WriteHeader(http.StatusBadGateway)
			return
		}
		return
	})

	http.HandleFunc("/feedback", func(w http.ResponseWriter, r *http.Request) {
		var f feedback.Feedback
		err := json.NewDecoder(r.Body).Decode(&f)
		if err != nil {
			w.WriteHeader(http.StatusBadGateway)
			return
		}
		email, err := us.ValidateToken(r.Context(), r.Header.Get("Authorization"))
		if err != nil {
			w.WriteHeader(http.StatusForbidden)
			return
		}
		f.Email = email
		var result struct {
			ID uuid.UUID `json:"id"`
		}
		result.ID, err = fdb.Store(r.Context(), f)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
		if err := json.NewEncoder(w).Encode(result); err != nil {
			w.WriteHeader(http.StatusBadGateway)
			return
		}
		return
	})

	http.HandleFunc("/vote", func(w http.ResponseWriter, r *http.Request) {
		var v vote.Vote
		err := json.NewDecoder(r.Body).Decode(&v)
		if err != nil {
			w.WriteHeader(http.StatusBadGateway)
			return
		}
		email, err := us.ValidateToken(r.Context(), r.Header.Get("Authorization"))
		if err != nil {
			w.WriteHeader(http.StatusForbidden)
			return
		}
		v.Email = email
		var result struct {
			ID uuid.UUID `json:"id"`
		}
		result.ID, err = vt.Store(r.Context(), v)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
		if err := json.NewEncoder(w).Encode(result); err != nil {
			w.WriteHeader(http.StatusBadGateway)
			return
		}
		return
	})
	http.Serve(lis, nil)
}
