package controllers

import (
	"net/http"

	"github.com/488Ques/aws-demo/config"
	"github.com/488Ques/aws-demo/models"
	"github.com/labstack/echo/v4"
)

func AllTrucks(c echo.Context) error {
	var trucks []*models.Truck
	db := config.DB()

	if err := db.Where("truck_status <> ?", "0").Find(&trucks).Error; err != nil {
		data := buildErrorJSON(err)
		return c.JSON(http.StatusInternalServerError, data)
	}

	response := buildResponseJSON(trucks)
	return c.JSON(http.StatusOK, response)
}

func CreateTruck(c echo.Context) error {
	t := new(models.Truck)
	db := config.DB()

	if err := c.Bind(t); err != nil {
		data := buildErrorJSON(err)
		return c.JSON(http.StatusInternalServerError, data)
	}

	truck := &models.Truck{
		TruckName:   t.TruckName,
		Location:    t.Location,
		TruckStatus: true,
	}

	if err := db.Create(&truck).Error; err != nil {
		data := buildErrorJSON(err)
		return c.JSON(http.StatusInternalServerError, data)
	}

	response := buildResponseJSON(truck)
	return c.JSON(http.StatusOK, response)
}

func UpdateTruck(c echo.Context) error {
	id := c.Param("id")
	t := new(models.Truck)
	db := config.DB()

	// Binding data
	if err := c.Bind(t); err != nil {
		data := buildErrorJSON(err)
		return c.JSON(http.StatusInternalServerError, data)
	}

	existingTruck := new(models.Truck)
	if err := db.First(&existingTruck, id).Error; err != nil {
		data := buildErrorJSON(err)
		return c.JSON(http.StatusNotFound, data)
	}

	existingTruck.TruckName = t.TruckName
	existingTruck.Location = t.Location
	if err := db.Save(&existingTruck).Error; err != nil {
		data := buildErrorJSON(err)
		return c.JSON(http.StatusInternalServerError, data)
	}

	response := buildResponseJSON(existingTruck)

	return c.JSON(http.StatusOK, response)
}

func GetTruck(c echo.Context) error {
	id := c.Param("id")
	truck := new(models.Truck)
	db := config.DB()

	if err := db.First(&truck, id).Error; err != nil {
		data := buildErrorJSON(err)
		return c.JSON(http.StatusNotFound, data)
	}

	response := buildResponseJSON(truck)

	return c.JSON(http.StatusOK, response)
}

func DeleteTruck(c echo.Context) error {
	id := c.Param("id")
	truck := new(models.Truck)
	db := config.DB()

	if err := db.First(&truck, id).Error; err != nil {
		data := buildErrorJSON(err)
		return c.JSON(http.StatusNotFound, data)
	}

	truck.TruckStatus = false
	if err := db.Save(truck).Error; err != nil {
		data := buildErrorJSON(err)
		return c.JSON(http.StatusInternalServerError, data)
	}

	response := map[string]interface{}{
		"message": "Truck ID " + id + " has been deleted",
	}
	return c.JSON(http.StatusOK, response)
}
