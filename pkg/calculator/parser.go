package calculator

import (
	"go/ast"
	"go/parser"
	"strconv"
	"strings"
)

type Task struct {
	ID            string  `json:"id"`
	Arg1          float64 `json:"arg1"`
	Arg2          float64 `json:"arg2"`
	Operation     string  `json:"operation"`
	Priority      int     `json:"priority"`
	OperationTime int     `json:"operation_time"`
}

func ParseExpression(expression string) ([]Task, error) {
	expression = strings.ReplaceAll(expression, ",", ".")

	expr, err := parser.ParseExpr(expression)
	if err != nil {
		return nil, err
	}

	var tasks []Task
	err = buildTasks(expr, &tasks)
	if err != nil {
		return nil, err
	}

	sortTasks(tasks)
	return tasks, nil
}

func buildTasks(node ast.Node, tasks *[]Task) error {
	switch n := node.(type) {
	case *ast.BinaryExpr:
		err := buildTasks(n.X, tasks)
		if err != nil {
			return err
		}
		err = buildTasks(n.Y, tasks)
		if err != nil {
			return err
		}

		priority := 0
		if n.Op.String() == "*" || n.Op.String() == "/" {
			priority = 1
		}

		*tasks = append(*tasks, Task{
			Arg1:      extractValue(n.X),
			Arg2:      extractValue(n.Y),
			Operation: n.Op.String(),
			Priority:  priority,
		})
	}
	return nil
}

func extractValue(node ast.Node) float64 {
	if lit, ok := node.(*ast.BasicLit); ok {
		value, _ := strconv.ParseFloat(lit.Value, 64)
		return value
	}
	return 0
}

func sortTasks(tasks []Task) {
	for i := 0; i < len(tasks)-1; i++ {
		for j := i + 1; j < len(tasks); j++ {
			if tasks[i].Priority < tasks[j].Priority {
				tasks[i], tasks[j] = tasks[j], tasks[i]
			}
		}
	}
}
