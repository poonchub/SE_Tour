package controllers

import (
	"net/http"
	"toursystem/config"
	"toursystem/entity"

	"github.com/gin-gonic/gin"
)

// GET /tour-schedule/:id
func GetTourScheduleByID(c *gin.Context) {
	var tourSchedule entity.TourSchedules
    id := c.Param("id")

    db := config.DB()

    if err := db.First(&tourSchedule, "id = ?", id).Error; err != nil {
        c.JSON(http.StatusNotFound, gin.H{"error": "tour package not found"})
        return
    }

    c.JSON(http.StatusOK, tourSchedule)
}

// PATCH /tour-schedule/:id
func UpdateTourSchedule(c *gin.Context) {
	ID := c.Param("id")

	var tourSchedule entity.TourSchedules

	db := config.DB()
	result := db.First(&tourSchedule, ID)
	if result.Error != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "id not found"})
		return
	}

	if err := c.ShouldBindJSON(&tourSchedule); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Bad request, unable to map payload"})
		return
	}

	result = db.Save(&tourSchedule)
	if result.Error != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Bad request"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Updated successful"})
}