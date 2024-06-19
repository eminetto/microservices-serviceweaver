package auth

import (
	"encoding/json"
	"net/http"
)

// HealthHandler handle the health request
func HealthHandler(us Auth) http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		msg, _ := us.Health(r.Context())
		w.Write([]byte(msg))
	})
}

// Handler handle the auth request
func Handler(us Auth) http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
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
}
