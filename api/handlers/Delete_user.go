package handlers

import (
	"backend/api/models"
	"net/http"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

// DeleteUser elimina un usuario de la base de datos
func DeleteUser(db *gorm.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		id := c.Param("id")
		var user models.Usuario
		result := db.First(&user, id)
		if result.Error != nil {
			c.JSON(http.StatusNotFound, gin.H{"error": "User no se encuentra en la base de datos"})
			return
		}

		db.Delete(&user)
		c.JSON(http.StatusNoContent, nil)
	}
}
