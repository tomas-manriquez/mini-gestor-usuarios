package main

import (
    "go-template/routes"
    "go-template/services"
    "github.com/gin-gonic/gin"
)

func main() {
    // Inicializar conexión Mongo
    services.InitMongo()

    r := gin.Default()

    // Registrar rutas
    routes.RegisterRoutes(r)

    r.Run(":8080")
}