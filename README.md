# quote_book
- REST API-сервис на Go для хранения и управления цитатами.

### Функционал
1. Добавление новой цитаты (POST /quotes)
2. Получение всех цитат (GET /quotes)
3. Получение случайной цитаты (GET /quotes/random)
4. Фильтрация по автору (GET /quotes?author=Confucius)
5. Удаление цитаты по ID (DELETE /quotes/{id})

### Запуск
1. Необходим `go 1.24.*`
2. Клонировать репозиторий:
    ```bash
    git clone https://github.com/yermaka-a/quote_book.git
    # Перейти в папку с проектом
    cd quote_book
    ```
3. 
    - Запусить с помощью Makefile командой
    `make`
    - Запустить командой `go run ./cmd/main.go`
4. Проект стартует на порту `8080` по умолчанию

### Использовать с помощью curl
```bash
# 1. Confucius
curl -X POST http://localhost:8080/quotes \
  -H "Content-Type: application/json" \
  -d '{"author":"Confucius", "quote":"Life is really simple, but we insist on making it complicated."}'

curl http://localhost:8080/quotes 

curl http://localhost:8080/quotes/random

curl http://localhost:8080/quotes?author=Confucius

curl -X DELETE http://localhost:8080/quotes/1
```

