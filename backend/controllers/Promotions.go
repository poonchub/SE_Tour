package controllers

import (
	"net/http"
	"toursystem/config"
	"toursystem/entity"

	"github.com/gin-gonic/gin"
)

// GET /promotion/:code
func GetPromotionByCode(c *gin.Context) {
	var promotion entity.Promotions
    id := c.Param("code")

    db := config.DB()

    if err := db.First(&promotion, "promotion_code = ?", id).Error; err != nil {
        c.JSON(http.StatusNotFound, gin.H{"error": "tour package not found"})
        return
    }

    c.JSON(http.StatusOK, promotion)
}