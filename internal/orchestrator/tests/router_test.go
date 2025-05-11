package tests

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/Danilka776/web_go_calc_2/internal/orchestrator/api"
)

func TestRouter(t *testing.T) {
	router := api.SetupRouter()

	req, _ := http.NewRequest("GET", "/api/v1/expressions", nil)
	rr := httptest.NewRecorder()

	router.ServeHTTP(rr, req)

	if rr.Code != http.StatusOK {
		t.Errorf("Ожидался код 200, получен %d", rr.Code)
	}
}
