FROM golang:1.21.0

WORKDIR /app

COPY ./go.mod ./go.sum ./

RUN go mod download

COPY . .

RUN CGO_ENABLED=0 go build -a -installsuffix cgo -o main ./cmd/server/server.go

EXPOSE 8080

CMD ["./main"]