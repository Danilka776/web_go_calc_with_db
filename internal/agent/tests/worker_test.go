package tests

import (
	"testing"

	"github.com/Danilka776/web_go_calc_with_db/internal/agent"
	"github.com/Danilka776/web_go_calc_with_db/internal/models"
)

func TestEvaluateTask(t *testing.T) {
	task := models.Task{Arg1: 10, Arg2: 50, Operation: "+"}
	result := agent.EvaluateTask(task)
	if result != 60 {
		t.Errorf("Ожидалось получить 60, а получено %v", result)
	}
}
