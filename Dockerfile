FROM golang:1.25-alpine AS builder

WORKDIR /app

RUN apk add --no-cache git

COPY go.mod go.sum ./
RUN go mod download

COPY . .

RUN go install github.com/pressly/goose/v3/cmd/goose@latest

RUN CGO_ENABLED=0 GOOS=linux go build -o qna-app ./cmd/api

FROM alpine:3.18

WORKDIR /app

RUN apk add --no-cache postgresql-client

COPY --from=builder /app/qna-app .
COPY --from=builder /go/bin/goose /usr/local/bin/goose
COPY --from=builder /app/migrations ./migrations

RUN ls -la /app/migrations/

EXPOSE 8080 

CMD ["/app/qna-app"]