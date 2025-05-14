# 1) Успешная регистрация
curl -i -X POST localhost:8080/api/v1/register \
  -H 'Content-Type: application/json' \
  -d '{"login":"danil","password":"pass123"}'

# 2) Ошибка: пользователь уже существует
curl -i -X POST localhost:8080/api/v1/register \
  -H 'Content-Type: application/json' \
  -d '{"login":"danil","password":"pass123"}'

# 3) Успешный вход
curl -i -X POST localhost:8080/api/v1/login \
  -H 'Content-Type: application/json' \
  -d '{"login":"danil","password":"pass123"}'

# 4) Неверные креды
curl -i -X POST localhost:8080/api/v1/login \
  -H 'Content-Type: application/json' \
  -d '{"login":"danil","password":"wrong"}'

# 5) Вызов калькулятора с токеном (успех)
curl -i -X POST localhost:8080/api/v1/calculate \
  -H 'Content-Type: application/json' \
  -H 'Authorization: Bearer <JWT_токен_после_авторизации>' \
  -d '{"expression":"2+2*2"}'

# 6) Вызов без токена (ошибка)
curl -i -X POST localhost:8080/api/v1/calculate \
  -H 'Content-Type: application/json' \
  -d '{"expression":"2+2*2"}'

# 7) Вызов калькулятора с неправильным токеном (ошибка)
curl -i -X POST localhost:8080/api/v1/calculate \
  -H 'Content-Type: application/json' \
  -H 'Authorization: Bearer <JWT_токен_после_авторизации_не_правильный>' \
  -d '{"expression":"2+2*2"}'