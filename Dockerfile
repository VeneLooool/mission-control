FROM golang:1.24 AS builder

WORKDIR /app
COPY . .

# ВАЖНО: явно указываем GOOS и GOARCH
ENV CGO_ENABLED=0 GOOS=linux GOARCH=amd64

RUN go build -o mission-control ./cmd/mission-control

FROM debian:bullseye-slim
WORKDIR /app

COPY --from=builder /app/mission-control .

COPY swagger-ui ./swagger-ui/
COPY internal ./internal/

EXPOSE 8084 50055

CMD ["./mission-control"]
