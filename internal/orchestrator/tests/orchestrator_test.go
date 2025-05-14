package tests

import (
	"testing"

	"github.com/Danilka776/web_go_calc_with_db/internal/orchestrator/services"
)

func TestOrchestratorFlow(t *testing.T) {

	id, err := services.AddExpression("2 + 3 * 4")
	if err != nil {
		t.Fatalf("Ошибка добавления выражения: %v", err)
	}

	if id == "" {
		t.Fatal("ID выражения не должен быть пустым")
	}

	task, err := services.GetTask()
	if err != nil {
		t.Fatalf("Ошибка получения задачи: %v", err)
	}

	if task.Operation != "*" {
		t.Errorf("Первая задача должна быть умножением, получена %s", task.Operation)
	}

	err = services.SubmitResult(task.ID, 12)
	if err != nil {
		t.Errorf("Ошибка отправки результата: %v", err)
	}
}
