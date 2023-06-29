package controllers

import (
	"biss/earthquake.be/initializers"
	"biss/earthquake.be/models"
	"time"

	"github.com/gin-gonic/gin"
)

func CreateHeartbeats(c *gin.Context) {
	//get data
	var body struct {
		DeviceID []int
	}

	c.Bind(&body)

	for deviceId := range body.DeviceID {

		device := models.Heartbeat{DeviceID: deviceId, UpdatedAt: time.Now()}
		initializers.DB.Create(&device)

		// var deviceModel models.Device
		// var model = initializers.DB.First(&deviceModel, deviceId)
		// if model == nil {
		// 	device := models.Heartbeat{DeviceID: deviceId, UpdatedAt: time.Now()}
		// 	initializers.DB.Create(&device)
		// } else {
		// 	initializers.DB.Model(&model).Updates(models.Device{
		// 		UpdatedAt: time.Now(),
		// 	})
		// }

	}

	c.Status(200)
}
