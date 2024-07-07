FROM golang:1.22.5-alpine

WORKDIR /app

COPY go.mod go.sum ./

RUN go mod download

COPY . .

RUN go build -o app/proxy-server ./cmd/main.go

Expose 8080

CMD ["app/proxy-server"]