package api

import (
	"encoding/json"
	"net/http"

	"github.com/Danilka776/web_go_calc_with_db/internal/models"
	"github.com/Danilka776/web_go_calc_with_db/internal/services"
)

func RegisterHandler(w http.ResponseWriter, r *http.Request) {
	var req models.RegisterRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "invalid request", http.StatusUnprocessableEntity)
		return
	}
	if err := services.RegisterUser(req.Login, req.Password); err != nil {
		if err == services.ErrUserExists {
			http.Error(w, "user already exists", http.StatusBadRequest)
		} else {
			http.Error(w, "internal error", http.StatusInternalServerError)
		}
		return
	}
	w.WriteHeader(http.StatusOK)
}

func LoginHandler(w http.ResponseWriter, r *http.Request) {
	var req models.LoginRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "invalid request", http.StatusUnprocessableEntity)
		return
	}
	token, err := services.Authenticate(req.Login, req.Password)
	if err != nil {
		if err == services.ErrInvalidCredentials {
			http.Error(w, "invalid credentials", http.StatusUnauthorized)
		} else {
			http.Error(w, "internal error", http.StatusInternalServerError)
		}
		return
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(models.AuthResponse{Token: token})
}
