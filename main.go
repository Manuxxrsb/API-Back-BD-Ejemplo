package main

import (
	"backend/api/handlers"
	"backend/api/utils"
	"log"

	"github.com/gin-gonic/gin"
)

func main() {
	// Conectar a la base de datos
	db, err := utils.OpenGormDB()
	if err != nil {
		log.Fatalf("Error connecting to the database: %v", err)
	}

	//db.AutoMigrate(&models.LogEntry{})		esta linea crea las tablas en la bd

	// Crear el router
	r := gin.Default()

	// Configurar las rutas para CRUD de usuarios
	r.POST("/users", handlers.CreateUser(db))
	r.GET("/users/:id", handlers.GetUser(db))
	r.GET("/users", handlers.GetAllUsers(db))
	r.PUT("/users/:id", handlers.UpdateUser(db))
	r.DELETE("/users/:id", handlers.DeleteUser(db))
	//r.POST("/log", handlers.CreateLogs(db))

	// Iniciar el servidor
	r.Run(":8084")

}
