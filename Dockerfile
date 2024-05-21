# Используем официальный образ Go для сборки приложения
FROM golang:1.18-alpine as builder

# Устанавливаем рабочую директорию внутри контейнера
WORKDIR /app

# Копируем go.mod и go.sum файлы
COPY go.mod go.sum ./

# Загружаем зависимости
RUN go mod download

# Копируем остальные файлы
COPY . .

# Создаем исполняемый файл
RUN go build -o text-style-service ./cmd

# Используем минимальный образ для запуска приложения
FROM alpine:latest

# Устанавливаем рабочую директорию внутри контейнера
WORKDIR /root/

# Копируем исполняемый файл из стадии сборки
COPY --from=builder /app/text-style-service .

# Пробрасываем порт для приложения
EXPOSE 3000

# Устанавливаем команду запуска контейнера
CMD ["./text-style-service"]

