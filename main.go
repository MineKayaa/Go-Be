package main

import (
	"biss/earthquake.be/controllers"
	"biss/earthquake.be/initializers"

	"github.com/gin-gonic/gin"
)

func init() {
	initializers.LoadEnvVariables()
	initializers.ConnectToDb()
}

func main() {
	r := gin.Default()

	r.POST("/create-technician", controllers.CreateTechnician)
	r.POST("/register", controllers.CreateDevice)
	r.POST("/heartbeats", controllers.CreateHeartbeats)
	r.PUT("/alert/:id", controllers.Alert)

	r.GET("/get-devices", controllers.GetDevices)
	r.GET("/get-device-info", controllers.GetDevice)
	r.GET("/get-alert", controllers.GetAlert)

	r.Run()
}
