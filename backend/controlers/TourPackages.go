package controllers

import (
	"net/http"
	"toursystem/config"
	"toursystem/entity"

	"github.com/gin-gonic/gin"
)

// GET /tour-packages
func ListTourPackages(c *gin.Context) {
	var tourPackages []entity.TourPackages

	db := config.DB()

	if err := db.Find(&tourPackages).Error; err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
        return
    }

	c.JSON(http.StatusOK, &tourPackages)
}
