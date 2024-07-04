FROM golang:1.18-alpine

WORKDIR /app

COPY go.mod go.sum ./

RUN go mod download

COPY . .

RUN go build -o /proxy-server

Expose 8080

CMD ["/proxy-server"]