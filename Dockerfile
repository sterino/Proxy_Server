FROM golang:1.22.5-alpine

WORKDIR /app

COPY go.mod go.sum ./

RUN go mod download

COPY . .

RUN go build -o /handler-server ./cmd/main.go

Expose 8080

CMD ["/proxy-server"]