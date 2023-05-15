package controllers

import (
	"net/http"

	"github.com/488Ques/aws-demo/config"
	"github.com/488Ques/aws-demo/models"
	"github.com/labstack/echo/v4"
)

func AllStaffs(c echo.Context) error {
	var staffs []*models.Staff
	db := config.DB()

	if err := db.Where("staff_status <> ?", "0").Find(&staffs).Error; err != nil {
		data := buildErrorJSON(err)
		return c.JSON(http.StatusInternalServerError, data)
	}

	response := buildResponseJSON(staffs)
	return c.JSON(http.StatusOK, response)
}

func CreateStaff(c echo.Context) error {
	s := new(models.Staff)
	db := config.DB()

	if err := c.Bind(s); err != nil {
		data := buildErrorJSON(err)
		return c.JSON(http.StatusInternalServerError, data)
	}

	staff := &models.Staff{
		StaffName:   s.StaffName,
		PhoneNumber: s.PhoneNumber,
		IsManager:   s.IsManager,
		StaffStatus: true,
	}

	if err := db.Create(&staff).Error; err != nil {
		data := buildErrorJSON(err)
		return c.JSON(http.StatusInternalServerError, data)
	}

	response := buildResponseJSON(staff)
	return c.JSON(http.StatusOK, response)
}

func UpdateStaff(c echo.Context) error {
	id := c.Param("id")
	s := new(models.Staff)
	db := config.DB()

	// Binding data
	if err := c.Bind(s); err != nil {
		data := buildErrorJSON(err)
		return c.JSON(http.StatusInternalServerError, data)
	}

	existingStaff := new(models.Staff)
	if err := db.First(&existingStaff, id).Error; err != nil {
		data := buildErrorJSON(err)
		return c.JSON(http.StatusNotFound, data)
	}

	existingStaff.StaffName = s.StaffName
	existingStaff.PhoneNumber = s.PhoneNumber
	existingStaff.IsManager = s.IsManager
	if err := db.Save(&existingStaff).Error; err != nil {
		data := buildErrorJSON(err)
		return c.JSON(http.StatusInternalServerError, data)
	}

	response := buildResponseJSON(existingStaff)

	return c.JSON(http.StatusOK, response)
}

func GetStaff(c echo.Context) error {
	id := c.Param("id")
	staff := new(models.Staff)
	db := config.DB()

	if err := db.First(&staff, id).Error; err != nil {
		data := buildErrorJSON(err)
		return c.JSON(http.StatusNotFound, data)
	}

	response := buildResponseJSON(staff)

	return c.JSON(http.StatusOK, response)
}

func DeleteStaff(c echo.Context) error {
	id := c.Param("id")
	staff := new(models.Staff)
	db := config.DB()

	if err := db.First(&staff, id).Error; err != nil {
		data := buildErrorJSON(err)
		return c.JSON(http.StatusNotFound, data)
	}

	staff.StaffStatus = false
	if err := db.Save(staff).Error; err != nil {
		data := buildErrorJSON(err)
		return c.JSON(http.StatusInternalServerError, data)
	}

	response := map[string]interface{}{
		"message": "Staff ID " + id + " has been deleted",
	}
	return c.JSON(http.StatusOK, response)
}
