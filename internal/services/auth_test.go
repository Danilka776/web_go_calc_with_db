package services

import (
	"os"
	"testing"

	"github.com/Danilka776/web_go_calc_with_db/internal/database"
)

func TestRegisterAndAuthenticate(t *testing.T) {
	os.Remove("test.db")
	database.Init("test.db")
	// успешная регистрация
	if err := RegisterUser("danil", "pass123"); err != nil {
		t.Fatalf("expected no error, got %v", err)
	}
	// дублирование
	if err := RegisterUser("danil", "pass1234"); err != ErrUserExists {
		t.Fatalf("expected ErrUserExists, got %v", err)
	}
	// правильная аутентификация
	token, err := Authenticate("danil", "pass123")
	if err != nil || token == "" {
		t.Fatalf("expected valid token, got err=%v token=%s", err, token)
	}
	// неправильный пароль
	if _, err := Authenticate("danil", "wrong"); err != ErrInvalidCredentials {
		t.Fatalf("expected ErrInvalidCredentials, got %v", err)
	}
	// несуществующий пользователь
	if _, err := Authenticate("stive", "pass"); err != ErrInvalidCredentials {
		t.Fatalf("expected ErrInvalidCredentials, got %v", err)
	}
}
