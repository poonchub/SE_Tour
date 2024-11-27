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

	if err := db.Preload("Province").Preload("TourPrices").Preload("TourImages").Preload("TourDescriptions").Preload("TourSchedules").Find(&tourPackages).Error; err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
        return
    }

	c.JSON(http.StatusOK, &tourPackages)
}

// GET /tour-package/:id
func GetTourPackageByID(c *gin.Context) {
	var tourpackage entity.TourPackages
    id := c.Param("id")

    db := config.DB()

    if err := db.Preload("Province").Preload("TourPrices.PersonType").Preload("TourImages").Preload("TourDescriptions").Preload("TourSchedules.TourScheduleStatus").Preload("TourSchedules.ScheduleActivities.Activity").First(&tourpackage, "id = ?", id).Error; err != nil {
        c.JSON(http.StatusNotFound, gin.H{"error": "tour package not found"})
        return
    }

    c.JSON(http.StatusOK, tourpackage)
}