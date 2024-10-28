# Dockerfile para servicio de create-cv.

FROM golang:1.23-alpine

WORKDIR /app

COPY go.mod go.sum ./

RUN go mod tidy
RUN go mod download

COPY . .

RUN go build -o create-cv cmd/main.go

EXPOSE 8081

CMD ["./create-cv"]