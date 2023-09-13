## build stage ##
FROM golang:1.21.0 AS builder
WORKDIR /app
COPY go.mod go.sum ./
RUN go mod download
COPY . .
RUN CGO_ENABLED=0 go build -a -installsuffix cgo -o main ./cmd/server/server.go

## run stage ##
FROM ubuntu:latest
ENV DB_HOST="mysql_db"
ENV DB_USER="root"
ENV DB_PASSWORD="password"
ENV DB_NAME="test_db"
WORKDIR /app
COPY --from=builder /app .
EXPOSE 8080
CMD [ "/app/main" ]