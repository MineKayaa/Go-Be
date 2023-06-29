package main

import (
	"biss/earthquake.be/initializers"
	"biss/earthquake.be/models"
)

func init() {
	initializers.LoadEnvVariables()
	initializers.ConnectToDb()
}

func main() {
	initializers.DB.AutoMigrate(&models.Technician{}, &models.Device{}, &models.Heartbeat{})
}
