package services

import (
	"errors"
	"sync"
	"time"

	"github.com/Danilka776/web_go_calc_2/internal/models"
	"github.com/Danilka776/web_go_calc_2/pkg/calculator"
)

var (
	expressions = make(map[string]models.ExpressionResponse)
	tasks       []calculator.Task
	results     = make(map[string]float64)
	queueMutex  sync.Mutex
	exprMutex   sync.Mutex
	resultMutex sync.Mutex
)

var ErrExpressionNotFound = errors.New("выражение не найдено")
var ErrNoTasks = errors.New("нет доступных задач")
var ErrTaskNotFound = errors.New("задача не найдена")

func AddExpression(value string) (string, error) {
	parsedTasks, err := calculator.ParseExpression(value)
	if err != nil {
		return "", err
	}

	id := generateID()
	expr := models.ExpressionResponse{
		ID:     id,
		Status: "pending",
		Result: 0,
	}

	queueMutex.Lock()
	expressions[id] = expr
	tasks = append(tasks, parsedTasks...)
	queueMutex.Unlock()

	return id, nil
}

func GetExpressions() []models.ExpressionResponse {
	exprMutex.Lock()
	defer exprMutex.Unlock()

	var exprs []models.ExpressionResponse
	for _, expr := range expressions {
		exprs = append(exprs, expr)
	}
	return exprs
}

func GetExpressionByID(id string) (models.ExpressionResponse, error) {
	exprMutex.Lock()
	defer exprMutex.Unlock()

	expr, exists := expressions[id]
	if !exists {
		return models.ExpressionResponse{}, ErrExpressionNotFound
	}
	return expr, nil
}

func GetTask() (models.Task, error) {
	queueMutex.Lock()
	defer queueMutex.Unlock()

	if len(tasks) == 0 {
		return models.Task{}, ErrNoTasks
	}

	task := tasks[0]
	tasks = tasks[1:]
	return models.Task(task), nil
}

func SubmitResult(taskID string, result float64) error {
	resultMutex.Lock()
	defer resultMutex.Unlock()

	found := false
	for _, task := range tasks {
		if task.ID == taskID {
			found = true
			break
		}
	}
	if !found {
		return ErrTaskNotFound
	}

	results[taskID] = result
	return nil
}

func generateID() string {
	return "expr_" + time.Now().Format("20060102150405")
}
