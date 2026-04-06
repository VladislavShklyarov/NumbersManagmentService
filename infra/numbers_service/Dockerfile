# ====== BUILD STAGE ======
FROM golang:1.25.5-alpine AS builder

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY . .

RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 \
    go build -o numbers_service ./cmd/main.go


# ====== RUNTIME STAGE ======
FROM alpine:3.20

RUN apk add --no-cache ca-certificates

WORKDIR /app

COPY --from=builder /app/numbers_service .


COPY infra/numbers_service/migrations /app/migrations

COPY configs /app/configs

EXPOSE 8080

RUN adduser -D appuser
USER appuser

CMD ["./numbers_service"]