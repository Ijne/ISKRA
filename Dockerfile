# builder stage
FROM golang:1.25.1-alpine AS builder

# Установка зависимостей для сборки (если нужны)
# RUN apk add --no-cache git ca-certificates

WORKDIR /app

# Копируем файлы зависимостей
COPY go.mod go.sum ./
RUN go mod download

COPY main/ ./main/
COPY miniapp/ ./miniapp/
COPY bot/ ./bot/
COPY shared/ ./shared/
# COPY config/ ./config/

# Собираем приложение
# RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o main ./main/cmd
# RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o main ./main/cmd/main.go
RUN CGO_ENABLED=0 GOOS=linux go build -o main ./main/cmd/main.go

# runner stage
FROM alpine:3.22

# Устанавливаем зависимости времени выполнения (если нужны)
# RUN apk --no-cache add ca-certificates

#WORKDIR /root
RUN mkdir -p /static/img 

# COPY --from=builder /app/main/cmd/main .
COPY --from=builder /app/main .
# COPY ./config/local.yaml ./config/
COPY ./config/deploy.yaml ./config/local.yaml

EXPOSE 8080

# Команда для запуска
CMD [ "./main" ]
# CMD [ "tree" ]