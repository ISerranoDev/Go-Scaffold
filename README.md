# Go-Scaffold API RESTful

Este proyecto implementa una API RESTful utilizando Go (Golang) que interactúa con una base de datos Postgres. La API soporta autenticación JWT y configuración de CORS.

## Requisitos

- Go 1.25
- POSTGRES 16

## Características

- **Base de datos POSTGRES 16**: Conexión a base de datos con configuración optimizada para conexiones seguras.
- **Autenticación JWT**: Manejo de autenticación usando tokens JWT.
- **CORS**: Configuración de orígenes confiables para solicitudes de recursos entre dominios.

## Cómo ejecutarlo

Para ejecutar el servidor, deberemos ejecutar primero la seed para instanciar un primer usuario y rol en nuestra aplicación.

```
go run /cmd/seed.go
```

Posteriormente, podremos ejecutar el servidor con el siguiente comando:

```
go run /cmd/main.go
```


