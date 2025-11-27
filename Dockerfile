FROM golang:1.25-alpine AS builder

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY . .

RUN go build -o main ./cmd/api

RUN go install github.com/pressly/goose/v3/cmd/goose@latest

FROM alpine:latest

RUN apk add --no-cache ca-certificates tzdata postgresql-client

WORKDIR /app

COPY --from=builder /app/main .

COPY --from=builder /go/bin/goose /usr/local/bin/goose

COPY --from=builder /app/migrations ./migrations

CMD ["./main"]