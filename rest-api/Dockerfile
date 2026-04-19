# Этап 1: Сборка приложения
FROM golang:1.26-alpine AS builder

RUN apk add --no-cache git ca-certificates

WORKDIR /app

# Копируем go.mod и go.sum
COPY go.mod go.sum ./

# Скачиваем зависимости
RUN go mod download

# Копируем исходный код
COPY . .

# Собираем бинарный файл
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o microservice ./cmd/app

# Этап 2: Финальный образ
FROM alpine:latest

RUN apk --no-cache add ca-certificates

WORKDIR /root/

COPY --from=builder /app/microservice .

EXPOSE 8080

CMD ["./microservice"]