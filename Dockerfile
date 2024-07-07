FROM golang:1.22.5-alpine

WORKDIR /app

COPY go.mod go.sum ./

RUN go mod download

COPY . .

RUN go build -o /proxy-server ./cmd/main.go

Expose 8080

CMD ["go", "run", "/proxy-server", "."]