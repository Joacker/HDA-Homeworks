#Se selecciona la imágen de golang para este caso, así mismo se establece la versión y el sistema operativo junto con el builder para el multi-stage
FROM golang:1.19-alpine AS builder

#Se establece el directorio de trabajo dentro del contenedor
WORKDIR /backend

#Se instala las dependencias necesarias para el proyecto
RUN apk update && apk add --no-cache git curl 

#Se copian los archivos necesarios para el proyecto
COPY go.mod ./
COPY go.sum ./

#Se copia todo el directorio actual al directorio de trabajo dentro del contenedor
COPY . .

#Se aplica el comando go mod para validar las dependencias
RUN go mod tidy

#Se establece el entorno de compilación
RUN go build -o main main.go

#Se selecciona la imágen de alpine para este caso, así mismo se establece la versión y el sistema operativo para empaquetar el contenedor anterior
FROM alpine:latest

#Se establece el directorio de trabajo dentro del contenedor
WORKDIR /backend

#Se copian los archivos necesarios para el proyecto pero con referencia al contenedor anterior
COPY --from=builder /backend .

#Se establece el puerto de escucha
EXPOSE 3000

#Se establece el comando para ejecutar el contenedor
CMD ["./main"]