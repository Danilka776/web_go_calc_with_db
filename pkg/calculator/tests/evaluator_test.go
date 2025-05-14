package tests

import (
	"testing"

	"github.com/Danilka776/web_go_calc_with_db/pkg/calculator"
)

func TestEvaluateOperations(t *testing.T) {
	testCases := []struct {
		arg1, arg2  float64
		op          string
		expected    float64
		expectError bool
	}{
		{10, 5, "+", 15, false},
		{10, 5, "-", 5, false},
		{10, 5, "*", 50, false},
		{10, 5, "/", 2, false},
		{10, 0, "/", 0, true},
	}

	for _, tc := range testCases {
		result, err := calculator.Evaluate(tc.arg1, tc.arg2, tc.op)
		if tc.expectError {
			if err == nil {
				t.Errorf("%v %s %v: ожидалась ошибка", tc.arg1, tc.op, tc.arg2)
			}
		} else {
			if err != nil {
				t.Errorf("%v %s %v: неожиданная ошибка: %v", tc.arg1, tc.op, tc.arg2, err)
			}
			if result != tc.expected {
				t.Errorf("%v %s %v = %v, ожидалось %v", tc.arg1, tc.op, tc.arg2, result, tc.expected)
			}
		}
	}
}
