FROM golang:latest AS builder

WORKDIR /app

COPY ../go.mod ../go.sum ./

RUN go mod download

COPY ../Pinger ./Pinger

WORKDIR /app/Pinger

RUN go build -o pinger cmd/app/main.go


FROM debian:bookworm-slim

WORKDIR /app/Pinger

COPY --from=builder /app/Pinger/pinger .

COPY ../Pinger/config/.env .

RUN apt update && apt install -y cron

RUN chmod +x /app/Pinger/pinger

RUN echo "* * * * * cd /app/Pinger && /app/Pinger/pinger >> /var/log/pinger.log 2>&1" > /etc/cron.d/pinger-cron \
    && chmod 0644 /etc/cron.d/pinger-cron \
    && crontab /etc/cron.d/pinger-cron

RUN touch /var/log/pinger.log

CMD ["cron", "-f"]