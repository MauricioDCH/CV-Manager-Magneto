# Dockerfile para servicio de cv.

FROM golang:1.23-alpine

WORKDIR /app

COPY go.mod go.sum ./

RUN go mod tidy
RUN go mod download

COPY . .

RUN go build -o cv cmd/main.go

EXPOSE 8008:8008

CMD ["./cv"]