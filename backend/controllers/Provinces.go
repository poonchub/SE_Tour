package controllers

import (
	"net/http"
	"toursystem/config"
	"toursystem/entity"

	"github.com/gin-gonic/gin"
)

// GET /province
func ListProvinces(c *gin.Context) {
	var province []entity.Provinces

	db := config.DB()

	if err := db.Find(&province).Error; err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
        return
    }

	c.JSON(http.StatusOK, &province)
}