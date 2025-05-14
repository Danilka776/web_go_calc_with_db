package api

import (
	"bytes"
	"encoding/json"
	"net/http/httptest"
	"os"
	"testing"

	"github.com/Danilka776/web_go_calc_with_db/internal/database"
)

func setup() {
	os.Remove("test.db")
	database.Init("test.db")
}

func TestRegisterHandler(t *testing.T) {
	setup()
	body, _ := json.Marshal(map[string]string{"login": "u1", "password": "p1"})
	req := httptest.NewRequest("POST", "/api/v1/register", bytes.NewReader(body))
	w := httptest.NewRecorder()
	RegisterHandler(w, req)
	if w.Code != 200 {
		t.Fatalf("expected 200, got %d", w.Code)
	}
	// повторная регистрация
	req = httptest.NewRequest("POST", "/api/v1/register", bytes.NewReader(body))
	w = httptest.NewRecorder()
	RegisterHandler(w, req)
	if w.Code != 400 {
		t.Fatalf("expected 400, got %d", w.Code)
	}
}

func TestLoginHandler(t *testing.T) {
	setup()
	// сначала зарегистрируем
	body, _ := json.Marshal(map[string]string{"login": "u2", "password": "p2"})
	httptest.NewRecorder()
	RegisterHandler(httptest.NewRecorder(), httptest.NewRequest("POST", "/api/v1/register", bytes.NewReader(body)))
	// правильный логин
	req := httptest.NewRequest("POST", "/api/v1/login", bytes.NewReader(body))
	w := httptest.NewRecorder()
	LoginHandler(w, req)
	if w.Code != 200 {
		t.Fatalf("expected 200, got %d", w.Code)
	}
	var resp map[string]string
	json.NewDecoder(w.Body).Decode(&resp)
	if resp["token"] == "" {
		t.Error("expected token in response")
	}
	// неверный пароль
	bad, _ := json.Marshal(map[string]string{"login": "u2", "password": "bad"})
	req = httptest.NewRequest("POST", "/api/v1/login", bytes.NewReader(bad))
	w = httptest.NewRecorder()
	LoginHandler(w, req)
	if w.Code != 401 {
		t.Fatalf("expected 401, got %d", w.Code)
	}
}
