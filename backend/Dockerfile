FROM golang:1.22 AS builder

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY . .

RUN go test .

RUN CGO_ENABLED=0 GOOS=linux go build -o myapp .

FROM alpine:latest

RUN apk --no-cache add ca-certificates

WORKDIR /

# Копируем скомпилированное приложение из стадии сборки
COPY --from=builder /app/myapp .

EXPOSE 4444

# Указываем команду для запуска приложения
CMD ["./myapp"]
