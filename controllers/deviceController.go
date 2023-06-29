package controllers

import (
	"biss/earthquake.be/initializers"
	"biss/earthquake.be/models"
	"time"

	"github.com/gin-gonic/gin"
)

func CreateDevice(c *gin.Context) {
	//get data
	var body struct {
		Latitude     float64
		Longitude    float64
		TechnicianID int
	}

	c.Bind(&body)

	//create a user
	device := models.Device{Latitude: body.Latitude, Longitude: body.Longitude, TechnicianID: body.TechnicianID, CreatedAt: time.Now(), IsAlertIssued: false}
	result := initializers.DB.Create(&device)

	if result.Error != nil {
		c.Status(400)
		return
	}

	c.JSON(200, gin.H{
		"device": device,
	})
}

func Alert(c *gin.Context) {

	//get id from url
	id := c.Param("id")

	var device models.Device
	initializers.DB.First(&device, id)

	initializers.DB.Model(&device).Updates(models.Device{
		IsAlertIssued: true,
		UpdatedAt:     time.Now(),
	})

	c.JSON(200, gin.H{
		"device": device,
	})
}

func GetDevices(c *gin.Context) {

	latitude := c.Param("latitude")
	longitude := c.Param("longitude")

	var devices []models.Device
	initializers.DB.Where("latitude  = ?", latitude).Where("longitude  = ?", longitude).Find(&devices)

	c.JSON(200, gin.H{
		"devices": devices,
	})
}

func GetDevice(c *gin.Context) {

	//get id from url
	id := c.Param("id")

	var device models.Device
	initializers.DB.First(&device, id)

	c.JSON(200, gin.H{
		"device": device,
	})
}

func GetAlert(c *gin.Context) {

	var devices []models.Device
	initializers.DB.Where("is_alert_issued  = ?", "true").Find(&devices)

	c.JSON(200, gin.H{
		"devices": devices,
	})
}
