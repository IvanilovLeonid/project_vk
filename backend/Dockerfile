# Используем официальный образ Go
FROM golang:1.23-alpine AS builder

# Устанавливаем рабочую директорию внутри контейнера
WORKDIR /app

# Копируем go.mod и go.sum из корня проекта в контейнер
COPY ../go.mod ../go.sum ./

# Загружаем все зависимости
RUN go mod tidy

# Копируем весь проект, включая папку backend
COPY ../ ./

# Переходим в каталог backend
WORKDIR /app/backend

# Загружаем зависимости и выполняем сборку проекта
RUN go mod tidy && go build -o backend .

# Запуск приложения
CMD ["./backend"]
