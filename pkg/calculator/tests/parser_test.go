package tests

import (
	"testing"

	"github.com/Danilka776/web_go_calc_with_db/pkg/calculator"
)

func TestParseExpressionPriority(t *testing.T) {
	expression := "2 + 3 * 4"
	tasks, err := calculator.ParseExpression(expression)
	if err != nil {
		t.Fatalf("Ошибка парсинга выражения: %v", err)
	}

	if len(tasks) != 2 {
		t.Errorf("Ожидалось 2 задачи, получено %d", len(tasks))
	}

	if tasks[0].Operation != "*" {
		t.Errorf("Ожидалось умножение в первой задаче, получено %s", tasks[0].Operation)
	}

	if tasks[1].Operation != "+" {
		t.Errorf("Ожидалось сложение во второй задаче, получено %s", tasks[1].Operation)
	}
}
