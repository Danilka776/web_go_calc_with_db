package tests

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/Danilka776/web_go_calc_2/internal/orchestrator/api"
)

func TestAPIHandlers(t *testing.T) {
	expr := map[string]string{"expression": "2 + 2"}
	body, _ := json.Marshal(expr)

	req, _ := http.NewRequest("POST", "/api/v1/calculate", bytes.NewBuffer(body))
	rr := httptest.NewRecorder()

	api.SubmitExpression(rr, req)

	if rr.Code != http.StatusCreated {
		t.Errorf("Ожидался код 201, получен %d", rr.Code)
	}

	var resp map[string]string
	json.Unmarshal(rr.Body.Bytes(), &resp)
	exprID := resp["id"]

	req, _ = http.NewRequest("GET", "/api/v1/expressions", nil)
	rr = httptest.NewRecorder()

	api.GetExpressions(rr, req)

	if rr.Code != http.StatusOK {
		t.Errorf("Ожидался код 200, получен %d", rr.Code)
	}

	req, _ = http.NewRequest("GET", "/api/v1/expressions/"+exprID, nil)
	rr = httptest.NewRecorder()

	api.GetExpressionByID(rr, req)

	if rr.Code != http.StatusOK {
		t.Errorf("Ожидался код 200, получен %d", rr.Code)
	}
}
