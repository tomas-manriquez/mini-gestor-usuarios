# mini-gestor-usuarios


Este proyecto consiste en una aplicación web que permite gestionar usuarios guardados en una base de datos.

# Pasos para ejecutar el proyecto

## 0. Requisitos

- Tener instalado MongoDB (la base de datos de la aplicación usa el nombre 'citiaps', y la colección 'users'), y corriendo en el puerto 27017
- Tener instalado Golang (el proyecto se desarrolló con go 1.24.5)
- Tener instalado Nuxt.js v4.0 (npm version 11.4.2, node version v24.3.0)
- Tener los siguientes puertos liberados: 8080, 3000

## 1. Carga de datos

1. Ir al directorio ```/go-template-structured\ (1)/populate``` en una terminal
2. Ejecutar el comando ```go version``` para asegurar que Go está instalado
3. Ejecutar el comando ```go run populate.go```. Debería entregar un output como el siguiente:
```plaintext
Connected to MongoDB
Inserted user IDs:
687ac5846a21b1cd4651f491
687ac5846a21b1cd4651f492
687ac5846a21b1cd4651f493
```
## 2. Ejecución de backend

1. Ir al directorio ```/go-template-structured\ (1)``` en una terminal
2. Ejecutar el comando ```go run . ```. Debería entregar un output como el siguiente:
```plaintext
Connected to MongoDB
[GIN-debug] [WARNING] Creating an Engine instance with the Logger and Recovery middleware already attached.

[GIN-debug] [WARNING] Running in "debug" mode. Switch to "release" mode in production.
 - using env:   export GIN_MODE=release
 - using code:  gin.SetMode(gin.ReleaseMode)

[GIN-debug] GET    /ping                     --> go-template/controllers.PingController (4 handlers)
[GIN-debug] POST   /users                    --> go-template/controllers.CreateUser (4 handlers)
[GIN-debug] GET    /users                    --> go-template/controllers.GetAllUsers (4 handlers)
[GIN-debug] GET    /users/:id                --> go-template/controllers.GetUserById (4 handlers)
[GIN-debug] [WARNING] You trusted all proxies, this is NOT safe. We recommend you to set a value.
Please check https://pkg.go.dev/github.com/gin-gonic/gin#readme-don-t-trust-all-proxies for details.
[GIN-debug] Listening and serving HTTP on :8080
```

## 3. Ejecución de frontend

1. Ir al directorio ```/frontend```
2. Ejecutar el comando ```npm install```
3. Ejecutar el comando  ```npm run dev```
4. Abrir una pestaña en un browser a localhost:3000
