FROM golang:1.22-alpine AS builder
WORKDIR /app
COPY . .

RUN go build -o main ./cmd/web/main.go

FROM alpine
WORKDIR /app
COPY --from=builder /app/main .

COPY notification_config.yaml .

EXPOSE 8081
CMD ["/app/main"]