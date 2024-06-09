# Stage 1: Build the Go application
FROM golang:1.22.4-bullseye AS builder

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY . ./

RUN apt-get update && apt-get install -y gcc sqlite3 libsqlite3-dev
RUN CGO_ENABLED=1 GOOS=linux go build -o /expenses-api

# Stage 2: Run the application
FROM debian:bullseye-slim

COPY --from=builder /expenses-api /expenses-api
COPY --from=builder /app/expenses.db /expenses.db

EXPOSE 8080

CMD ["/expenses-api"]
