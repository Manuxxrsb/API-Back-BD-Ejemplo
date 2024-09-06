package handlers

import (
	"backend/api/models"
	"net/http"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func CreateUser(db *gorm.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		var user models.Usuario
		if err := c.ShouldBindJSON(&user); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		result := db.Create(&user) //comando para migrar el modelo a la bd
		if result.Error != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": result.Error.Error()})
			return
		}

		c.JSON(http.StatusCreated, user)
	}
}
