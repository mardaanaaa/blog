# Этап сборки
FROM golang:1.23.4 AS builder

# Устанавливаем рабочую директорию для сборки
WORKDIR /app

# Копируем go.mod и загружаем зависимости
COPY go.mod ./
RUN go mod download

# Копируем весь исходный код проекта
COPY . ./

# Строим приложение
RUN go build -o main ./cmd/main.go

# Финальный образ на основе Ubuntu
FROM ubuntu:22.04

# Копируем wait-for-it.sh в контейнер
COPY wait-for-it.sh /app/wait-for-it.sh
RUN chmod +x /app/wait-for-it.sh

# Копируем миграции
COPY ./internal/db/migrations /app/internal/db/migrations

# Устанавливаем рабочую директорию
WORKDIR /app

# Копируем скомпилированный файл из этапа сборки
COPY --from=builder /app/main ./

# Открываем порт для приложения
EXPOSE 8080

# Команда для запуска приложения
CMD ["/app/wait-for-it.sh", "my_postgres_db:5432", "--", "/app/main"]
