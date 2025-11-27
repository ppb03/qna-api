# Q&A API
*Тестовое задание*

RESTful API для системы вопросов и ответов. Позволяет создавать вопросы, добавлять ответы, просматривать и удалять их. Реализовано на Go с использованием PostgreSQL в качестве базы данных и goose в качестве инструмента для миграций. Из внешних библиотек были использованы `gorm.io/gorm` и `github.com/stretchr/testify`.

Тесты содержатся в `internal/handler/handler_test.go`. В нем покрыты варианты использования (в т.ч. ошибочные) как хендлеров, так и сервисного слоя. Слой репозиториев тестами не покрыт, вместо него используются моки из `internal/mocks/repository.go`.

## Установка и запуск

**1. Склонируйте этот репозиторий:**
```sh
git clone https://github.com/ppb03/qna-api.git
```

**2. Создайте `.env` файл, можно использовать файл-пример `.env.example`:**
```env
DB_USER=postgres
DB_PASSWORD=password
DB_NAME=qna_db
DB_HOST=localhost
DB_PORT=5432
DB_SSLMODE=disable
```

**3. Запустите сборку приложения и создание БД:**
```sh
docker compose build
```

**4. Запустите миграции через goose:**
```sh
docker compose --profile migrations run migrations
```

**5. Запустите приложение:**
```sh
docker compose up
```

**\* Примеры запросов к API:**
```bash
# Создание вопроса
curl -X POST http://localhost:8080/questions/ \
  -H "Content-Type: application/json" \
  -d '{
    "text": "Как работает Docker Compose?"
  }'

# Получение списка всех ответов
curl -X GET http://localhost:8080/questions/

# Получение конкретного вопроса с ответами
curl -X GET http://localhost:8080/questions/1/
```

**\* Миграции могут быть запущены отдельно:**
```sh
docker compose build migrations
```

## Методы API

### 1. Вопросы (Questions):

- `POST /questions/` - создать новый вопрос
  - **Тело запроса:**
  - `text`: строка, не может быть пустой
- `DELETE /questions/{id}` - удалить вопрос (вместе с ответами)
- `GET /questions/` - список всех вопросов
- `GET /questions/{id}` - получить вопрос и все ответы на него

### 2. Ответы (Answers):

- `POST /questions/{id}/answers/` - добавить ответ к вопросу
  - **Тело запроса:**
  - `user_id`: строка, обязана соответствовать формату UUID, не может быть пустой
  - `text`: строка, не может быть пустой
- `DELETE /answers/{id}` - удалить ответ
- `GET /answers/{id}` - получить конкретный ответ

**Логика:**
- *Нельзя создать ответ к несуществующему вопросу.*
- *Один и тот же пользователь может оставлять несколько ответов на один вопрос.*
- *При удалении вопроса должны удаляться все его ответы (каскадно).*