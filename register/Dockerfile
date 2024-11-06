# Dockerfile para servicio de registro de usuarios

FROM golang:1.23-alpine

WORKDIR /app

COPY go.mod go.sum ./

RUN go mod tidy
RUN go mod download

COPY . .

RUN go build -o register cmd/main.go

EXPOSE 8080:8080

CMD ["./register"]
