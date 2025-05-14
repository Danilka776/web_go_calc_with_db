package agent

import (
	"bytes"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
	"strconv"
	"time"

	"github.com/Danilka776/web_go_calc_with_db/internal/models"
)

func EvaluateTask(task models.Task) float64 {
	return evaluateTask(task)
}

func Start(orchestratorURL string, computingPower int) error {
	for i := 0; i < computingPower; i++ {
		go func(id int) {
			for {
				resp, err := http.Get(fmt.Sprintf("%s/internal/task", orchestratorURL))
				if err != nil {
					log.Printf("Агент %d: Ошибка при получении задачи: %v", id, err)
					time.Sleep(time.Second)
					continue
				}

				var task models.Task
				if err := json.NewDecoder(resp.Body).Decode(&task); err != nil {
					log.Printf("Агент %d: Ошибка при декодировании задачи: %v", id, err)
					resp.Body.Close()
					continue
				}
				resp.Body.Close()

				result := evaluateTask(task)
				sendResult(orchestratorURL, task.ID, result)
			}
		}(i)
	}
	return nil
}

func evaluateTask(task models.Task) float64 {
	time.Sleep(getOperationTime(task.Operation))

	switch task.Operation {
	case "+":
		return task.Arg1 + task.Arg2
	case "-":
		return task.Arg1 - task.Arg2
	case "*":
		return task.Arg1 * task.Arg2
	case "/":
		if task.Arg2 == 0 {
			return 0
		}
		return task.Arg1 / task.Arg2
	default:
		return 0
	}
}

func sendResult(orchestratorURL, taskID string, result float64) {
	client := &http.Client{}
	reqBody := map[string]interface{}{"id": taskID, "result": result}
	body, _ := json.Marshal(reqBody)

	req, _ := http.NewRequest("POST", fmt.Sprintf("%s/internal/task", orchestratorURL), bytes.NewBuffer(body))
	req.Header.Set("Content-Type", "application/json")
	client.Do(req)
}

func getOperationTime(operation string) time.Duration {
	switch operation {
	case "+":
		return time.Millisecond * time.Duration(getEnvInt("TIME_ADDITION_MS", 500))
	case "-":
		return time.Millisecond * time.Duration(getEnvInt("TIME_SUBTRACTION_MS", 500))
	case "*":
		return time.Millisecond * time.Duration(getEnvInt("TIME_MULTIPLICATION_MS", 1000))
	case "/":
		return time.Millisecond * time.Duration(getEnvInt("TIME_DIVISION_MS", 1500))
	default:
		return time.Millisecond * 500
	}
}

func getEnvInt(key string, defaultValue int) int {
	valueStr := os.Getenv(key)
	if valueStr == "" {
		return defaultValue
	}
	value, _ := strconv.Atoi(valueStr)
	return value
}
