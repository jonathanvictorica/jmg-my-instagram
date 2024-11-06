# Usar una imagen oficial de Go como base
FROM golang:1.23-alpine AS builder

# Establecer el directorio de trabajo dentro del contenedor
WORKDIR /app

# Copiar el go.mod y go.sum para instalar las dependencias
COPY go.mod go.sum ./

# Descargar las dependencias (esto se hace antes de copiar todo el código para aprovechar la cache de Docker)
RUN go mod tidy

# Copiar el resto del código fuente
COPY . .

# Compilar la aplicación (cambia el nombre de tu ejecutable si es necesario)
RUN GOOS=linux GOARCH=amd64 go build -o jmg-my-instagram .

# Usar una imagen más ligera de Alpine para ejecutar la aplicación
FROM alpine:latest

# Instalar dependencias mínimas necesarias (si tienes dependencias nativas, como librerías de SSL, etc.)
RUN apk --no-cache add ca-certificates

# Establecer el directorio de trabajo dentro del contenedor
WORKDIR /root/

# Copiar el binario de la etapa de construcción al contenedor
COPY --from=builder /app/jmg-my-instagram .

# Exponer el puerto en el que se ejecutará el microservicio
EXPOSE 8080

# Comando para ejecutar el microservicio
CMD ["./jmg-my-instagram"]
