# Usar la imagen base de golang 1.21-alpine
FROM --platform=linux/amd64 golang:1.25.5-alpine
#FROM golang:1.23.1-alpine

# Establecer el directorio de trabajo dentro del contenedor
WORKDIR /app

# Instalar dependencias necesarias (bash, git, etc.)
RUN apk add --no-cache bash git curl

# Copiar el código de la aplicación (puedes modificar según tu estructura)
COPY . /app

# Establecer las variables de entorno necesarias
ENV ENV=development

# Exponer el puerto en el que correrá la aplicación
EXPOSE 4000

# Entrar en bash al iniciar el contenedor
CMD ["sh"]
