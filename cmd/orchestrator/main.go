package main

import (
	"log"
	"net/http"

	"github.com/Danilka776/web_go_calc_2/internal/orchestrator/api"
)

func main() {
	log.Println("Запуск оркестратора...")
	router := api.SetupRouter()
	if err := http.ListenAndServe(":8080", router); err != nil {
		log.Fatalf("Ошибка при запуске оркестратора: %v", err)
	}
}
