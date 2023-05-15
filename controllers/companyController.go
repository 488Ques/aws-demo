package controllers

import (
	"net/http"

	"github.com/488Ques/aws-demo/config"
	"github.com/488Ques/aws-demo/models"
	"github.com/labstack/echo/v4"
)

func AllCompanies(c echo.Context) error {
	var companies []*models.Company
	db := config.DB()

	if err := db.Where("company_status <> ?", "0").Find(&companies).Error; err != nil {
		data := buildErrorJSON(err)
		return c.JSON(http.StatusInternalServerError, data)
	}

	response := buildResponseJSON(companies)
	return c.JSON(http.StatusOK, response)
}

func CreateCompany(c echo.Context) error {
	comp := new(models.Company)
	db := config.DB()

	if err := c.Bind(comp); err != nil {
		data := buildErrorJSON(err)
		return c.JSON(http.StatusInternalServerError, data)
	}

	company := &models.Company{
		CompanyName:   comp.CompanyName,
		Address:       comp.Address,
		Email:         comp.Email,
		CompanyStatus: true,
	}

	if err := db.Create(&company).Error; err != nil {
		data := buildErrorJSON(err)
		return c.JSON(http.StatusInternalServerError, data)
	}

	response := buildResponseJSON(company)
	return c.JSON(http.StatusOK, response)
}

func UpdateCompany(c echo.Context) error {
	id := c.Param("id")
	comp := new(models.Company)
	db := config.DB()

	// Binding data
	if err := c.Bind(comp); err != nil {
		data := buildErrorJSON(err)
		return c.JSON(http.StatusInternalServerError, data)
	}

	existingCompany := new(models.Company)
	if err := db.First(&existingCompany, id).Error; err != nil {
		data := buildErrorJSON(err)
		return c.JSON(http.StatusNotFound, data)
	}

	existingCompany.CompanyName = comp.CompanyName
	existingCompany.Address = comp.Address
	existingCompany.Email = comp.Email
	existingCompany.CompanyStatus = comp.CompanyStatus
	if err := db.Save(&existingCompany).Error; err != nil {
		data := buildErrorJSON(err)
		return c.JSON(http.StatusInternalServerError, data)
	}

	response := buildResponseJSON(existingCompany)

	return c.JSON(http.StatusOK, response)
}

func GetCompany(c echo.Context) error {
	id := c.Param("id")
	company := new(models.Company)
	db := config.DB()

	if err := db.First(&company, id).Error; err != nil {
		data := buildErrorJSON(err)
		return c.JSON(http.StatusNotFound, data)
	}

	response := buildResponseJSON(company)

	return c.JSON(http.StatusOK, response)
}

func DeleteCompany(c echo.Context) error {
	id := c.Param("id")
	company := new(models.Company)
	db := config.DB()

	if err := db.First(&company, id).Error; err != nil {
		data := buildErrorJSON(err)
		return c.JSON(http.StatusNotFound, data)
	}

	company.CompanyStatus = false
	if err := db.Save(company).Error; err != nil {
		data := buildErrorJSON(err)
		return c.JSON(http.StatusInternalServerError, data)
	}

	response := map[string]interface{}{
		"message": "Company ID " + id + " has been deleted",
	}
	return c.JSON(http.StatusOK, response)
}
