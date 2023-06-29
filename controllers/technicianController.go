package controllers

import (
	"biss/earthquake.be/initializers"
	"biss/earthquake.be/models"

	"github.com/gin-gonic/gin"
)

func CreateTechnician(c *gin.Context) {
	//get data
	var body struct {
		Name string
	}

	c.Bind(&body)

	//create a technician
	technician := models.Technician{Name: body.Name}
	result := initializers.DB.Create(&technician)

	if result.Error != nil {
		c.Status(400)
		return
	}

	c.JSON(200, gin.H{
		"technician": technician,
	})
}
