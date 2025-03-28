
FROM golang:1.24-alpine

# Establece el directorio de trabajo
WORKDIR /app

# Instala dependencias del sistema 
RUN apk add --no-cache git gcc musl-dev

# Copia los archivos de dependencias
COPY go.mod go.sum ./

# Descarga las dependencias
RUN go mod download

# Copia todo el código f
COPY . .

# Compila la aplicación
RUN go build -o main .

# Expone el puerto que usa tu backend
EXPOSE 8080

# Comando para ejecutar la aplicación
CMD ["./main"]