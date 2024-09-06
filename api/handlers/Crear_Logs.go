package handlers

import (
	"backend/api/models"
	"net/http"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func CreateLogs(db *gorm.DB) gin.HandlerFunc {

	return func(c *gin.Context) {

		var Log models.LogEntry                        // creo el modelo
		if err := c.ShouldBindJSON(&Log); err != nil { //manejo de errores
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		result := db.Create(&Log)
		if result.Error != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": result.Error.Error()}) //retorno un erorr si no se pudo migrar a la bd
			return
		}

		c.JSON(http.StatusCreated, Log) //retorno una respuesta
	}

}
