package main

import (
	"log"
	"os"
	"strconv"

	"github.com/Danilka776/web_go_calc_with_db/internal/agent"
)

func main() {
	log.Println("Запуск агента...")

	orchestratorURL := os.Getenv("ORCHESTRATOR_URL")
	if orchestratorURL == "" {
		orchestratorURL = "http://localhost:8080"
	}

	computingPowerStr := os.Getenv("COMPUTING_POWER")
	if computingPowerStr == "" {
		computingPowerStr = "2"
	}

	computingPower, err := strconv.Atoi(computingPowerStr)
	if err != nil || computingPower <= 0 {
		log.Fatalf("Неверное значение COMPUTING_POWER: %s", computingPowerStr)
	}

	if err := agent.Start(orchestratorURL, computingPower); err != nil {
		log.Fatalf("Ошибка при запуске агента: %v", err)
	}
}
