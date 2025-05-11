package calculator

import (
	"errors"
)

func Evaluate(arg1, arg2 float64, operation string) (float64, error) {
	switch operation {
	case "+":
		return arg1 + arg2, nil
	case "-":
		return arg1 - arg2, nil
	case "*":
		return arg1 * arg2, nil
	case "/":
		if arg2 == 0 {
			return 0, errors.New("деление на ноль")
		}
		return arg1 / arg2, nil
	default:
		return 0, errors.New("неизвестная операция")
	}
}
