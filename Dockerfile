FROM golang:1.25.1-alpine AS builder

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY main/ ./main/
COPY miniapp/ ./miniapp/
COPY bot/ ./bot/
COPY shared/ ./shared/

RUN CGO_ENABLED=0 GOOS=linux go build -o main ./main/cmd/main.go

FROM alpine:3.22

RUN mkdir -p /static/img 

COPY --from=builder /app/main .
COPY ./config/deploy.yaml ./config/local.yaml

EXPOSE 8080

CMD [ "./main" ]