package tests

import (
	"testing"

	"github.com/Danilka776/web_go_calc_2/internal/models"
)

func TestTask(t *testing.T) {
	task := models.Task{ID: "1", Arg1: 2, Arg2: 3, Operation: "+", Priority: 0, OperationTime: 500}
	if task.ID != "1" || task.Arg1 != 2 || task.Arg2 != 3 || task.Operation != "+" || task.Priority != 0 || task.OperationTime != 500 {
		t.Errorf("Некорректные данные задачи")
	}
}
