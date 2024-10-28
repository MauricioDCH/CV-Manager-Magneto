# Dockerfile para servicio de rellanado de hojas de vida.

FROM golang:1.23-alpine

WORKDIR /app

COPY go.mod go.sum ./

RUN go mod tidy
RUN go mod download

COPY . .

RUN go build -o extension-server cmd/main.go

EXPOSE 5000

CMD ["./extension-server"]