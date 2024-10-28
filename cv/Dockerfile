# Dockerfile para servicio de cv.

# Usar una imagen base de Go
FROM golang:1.23-alpine

# Establecer el directorio de trabajo
WORKDIR /app

# Copia los archivos go.mod y go.sum primero para aprovechar la caché de Docker si las dependencias no han cambiado
COPY go.mod go.sum ./

# Ejecutar go mod tidy para gestionar las dependencias
RUN go mod tidy

# Descarga las dependencias después de hacer tidy para asegurar consistencia
RUN go mod download

# Copiar el código fuente
COPY . .

# Compila el proyecto
RUN go build -o cv cmd/main.go

# Exponer el puerto
EXPOSE 8008

# Comando para ejecutar el servicio
CMD ["./cv"]