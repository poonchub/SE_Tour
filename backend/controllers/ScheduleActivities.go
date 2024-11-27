package controllers

import (
	"net/http"
	"toursystem/config"
	"toursystem/entity"

	"github.com/gin-gonic/gin"
)

// GET /schedule-activity/:id
func GetScheduleActivityByTourScheduleID(c *gin.Context) {
	var scheduleActivity []entity.ScheduleActivities
    id := c.Param("id")

    db := config.DB()

    if err := db.Preload("Activity").Find(&scheduleActivity, "tour_schedule_id = ?", id).Error; err != nil {
        c.JSON(http.StatusNotFound, gin.H{"error": "tour package not found"})
        return
    }

    c.JSON(http.StatusOK, scheduleActivity)
}