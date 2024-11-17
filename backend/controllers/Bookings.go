package controllers

import (
	"net/http"
	"toursystem/config"
	"toursystem/entity"

	"github.com/gin-gonic/gin"
)

// GET /bookings
func ListBookings(c *gin.Context) {
	var bookings []entity.Bookings

	db := config.DB()

	db.Preload("Customer").Preload("TourSchedule").Preload("BookingStatus").Preload("Promotion").Find(&bookings)

	c.JSON(http.StatusOK, &bookings)
}

// GET /booking/:id
func GetBookingByID(c *gin.Context) {
	ID := c.Param("id")
	var booking entity.Bookings

	db := config.DB()
	results := db.Preload("Customer").Preload("TourSchedule").Preload("BookingStatus").Preload("Promotion").First(&booking, ID)
	if results.Error != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": results.Error.Error()})
		return
	}
	if booking.ID == 0 {
		c.JSON(http.StatusNoContent, gin.H{})
		return
	}
	c.JSON(http.StatusOK, booking)
}

// GET /bookings/:id
func GetBookingByCustomerID(c *gin.Context) {
	ID := c.Param("id")
	var bookings []entity.Bookings

	db := config.DB()
	results := db.Preload("Customer").Preload("TourSchedule").Preload("BookingStatus").Preload("Promotion").Find(&bookings, "customer_id=?", ID)
	if results.Error != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": results.Error.Error()})
		return
	}
	c.JSON(http.StatusOK, bookings)
}

// POST /booking
func CreateBooking(c *gin.Context) {
	var booking entity.Bookings

	if err := c.ShouldBindJSON(&booking); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	db := config.DB()

	var customer entity.Customers
	db.First(&customer, booking.CustomerID)
	if customer.ID == 0 {
		c.JSON(http.StatusNotFound, gin.H{"error": "customer not found"})
		return
	}

	var tourSchedule entity.TourSchedules
	db.First(&tourSchedule, booking.TourScheduleID)
	if customer.ID == 0 {
		c.JSON(http.StatusNotFound, gin.H{"error": "tour packege not found"})
		return
	}

	var promotion entity.Promotions
	db.First(&promotion, booking.PromotionID)
	if customer.ID == 0 {
		c.JSON(http.StatusNotFound, gin.H{"error": "promotion not found"})
		return
	}

	bk := entity.Bookings{
		TotalPrice: 0,
		CustomerID: booking.CustomerID,
		TourScheduleID: booking.TourScheduleID,
		BookingStatusID: booking.BookingStatusID,
		PromotionID: booking.PromotionID,
	}

	if err := db.Preload("Customer").Create(&bk).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"message": "Created success", "data": bk})
}

// PATCH /booking
func UpdateBooking(c *gin.Context) {
	ID := c.Param("id")

	var booking entity.Bookings

	db := config.DB()
	result := db.First(&booking, ID)
	if result.Error != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "id not found"})
		return
	}

	if err := c.ShouldBindJSON(&booking); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Bad request, unable to map payload"})
		return
	}

	result = db.Save(&booking)
	if result.Error != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Bad request"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Updated successful"})
}