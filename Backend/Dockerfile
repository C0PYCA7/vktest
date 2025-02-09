FROM golang:latest AS builder

WORKDIR /app

COPY ../go.mod ../go.sum ./

RUN go mod download

COPY ../Backend ./Backend

WORKDIR /app/Backend

RUN go build -o backend cmd/app/main.go

FROM debian:bookworm-slim

WORKDIR /app/Backend

COPY --from=builder /app/Backend/backend .

COPY ../Backend/config/.env .

RUN apt-get update && apt-get install -y curl postgresql-client bash && \
    curl -sSL https://github.com/pressly/goose/releases/download/v3.15.1/goose_linux_x86_64 -o /usr/local/bin/goose && \
    chmod +x /usr/local/bin/goose

COPY ../Backend/internal/storage/migrations /migrations

CMD ["sh", "-c", "goose -dir /migrations postgres 'host=db port=5432 user=postgres password=1977 dbname=vk sslmode=disable' up && ./backend"]