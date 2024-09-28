FROM golang:1.22 AS builder

WORKDIR /backend

COPY go.mod go.sum ./
RUN go mod download

COPY . .

RUN go test ./...

RUN go build -o request-tracker .

FROM alpine:latest

RUN apk --no-cache add ca-certificates

# Копируем собранное приложение из образа сборки
COPY --from=builder /backend/request-tracker /usr/local/bin/request-tracker

CMD ["request-tracker"]
