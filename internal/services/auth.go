package services

import (
	"database/sql"
	"errors"

	"golang.org/x/crypto/bcrypt"

	"github.com/Danilka776/web_go_calc_with_db/internal/database"
	"github.com/golang-jwt/jwt"
)

var (
	ErrUserExists         = errors.New("user already exists")
	ErrInvalidCredentials = errors.New("invalid login or password")
	jwtSecret             = []byte("very_secret_value")
)

func RegisterUser(login, password string) error {
	var exists int
	err := database.DB.QueryRow("SELECT COUNT(1) FROM users WHERE login = ?", login).Scan(&exists)
	if err != nil {
		return err
	}
	if exists > 0 {
		return ErrUserExists
	}
	hash, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return err
	}
	_, err = database.DB.Exec(
		"INSERT INTO users(login, password_hash) VALUES(?, ?)",
		login, string(hash),
	)
	return err
}

func Authenticate(login, password string) (string, error) {
	var hash string
	err := database.DB.QueryRow("SELECT password_hash FROM users WHERE login = ?", login).Scan(&hash)
	if err == sql.ErrNoRows {
		return "", ErrInvalidCredentials
	}
	if err != nil {
		return "", err
	}
	if bcrypt.CompareHashAndPassword([]byte(hash), []byte(password)) != nil {
		return "", ErrInvalidCredentials
	}
	// создаём токен
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"sub": login,
	})
	return token.SignedString(jwtSecret)
}
