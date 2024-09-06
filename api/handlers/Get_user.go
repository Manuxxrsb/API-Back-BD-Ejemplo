package handlers

import (
	"backend/api/models"
	"net/http"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func GetUser(db *gorm.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		id := c.Param("id")
		var user models.Usuario
		if err := db.Preload("Monedero.Monedas").First(&user, id).Error; err != nil {
			c.JSON(http.StatusNotFound, gin.H{"error": "User not found"})
			return
		}
		c.JSON(http.StatusOK, user)
	}
}

func GetAllUsers(db *gorm.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		var users []models.Usuario
		if err := db.Preload("Monedero.Monedas").Find(&users).Error; err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		c.JSON(http.StatusOK, users)
	}
}
