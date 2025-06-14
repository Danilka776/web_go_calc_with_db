package api

import (
	"net/http"

	"github.com/gorilla/mux"
)

func SetupRouter() http.Handler {
	r := mux.NewRouter()

	r.HandleFunc("/api/v1/calculate", ServeCalculatorForm).Methods("GET")
	//r.HandleFunc("/api/v1/calculate", SubmitExpression).Methods("POST")

	r.HandleFunc("/api/v1/register", RegisterHandler).Methods("POST")
	r.HandleFunc("/api/v1/login", LoginHandler).Methods("POST")

	auth := JWTMiddleware([]byte("very_secret_value"))
	r.Handle("/api/v1/calculate", auth(http.HandlerFunc(SubmitExpression))).Methods("POST")
	r.HandleFunc("/api/v1/expressions", GetExpressions).Methods("GET")
	r.HandleFunc("/api/v1/expressions/{id}", GetExpressionByID).Methods("GET")
	r.HandleFunc("/internal/task", GetTask).Methods("GET")
	r.HandleFunc("/internal/task", SubmitResult).Methods("POST")

	r.PathPrefix("/public/").Handler(http.StripPrefix("/public/", http.FileServer(http.Dir("../../public/"))))

	return r
}
