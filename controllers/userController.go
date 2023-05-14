package controllers

import (
	"net/http"

	"github.com/488Ques/aws-demo/config"
	"github.com/488Ques/aws-demo/models"
	"github.com/labstack/echo/v4"
)

func AllUsers(c echo.Context) error {
	var users []*models.User
	db := config.DB()

	if err := db.Where("user_status <> ?", "0").Find(&users).Error; err != nil {
		data := buildErrorJSON(err)
		return c.JSON(http.StatusInternalServerError, data)
	}

	response := buildResponseJSON(users)
	return c.JSON(http.StatusOK, response)
}

func CreateUser(c echo.Context) error {
	u := new(models.User)
	db := config.DB()

	if err := c.Bind(u); err != nil {
		data := buildErrorJSON(err)
		return c.JSON(http.StatusInternalServerError, data)
	}

	user := &models.User{
		Username:   u.Username,
		Password:   u.Password,
		UserStatus: true,
	}

	if err := db.Create(&user).Error; err != nil {
		data := buildErrorJSON(err)
		return c.JSON(http.StatusInternalServerError, data)
	}

	response := buildResponseJSON(user)
	return c.JSON(http.StatusOK, response)
}

func UpdateUser(c echo.Context) error {
	id := c.Param("id")
	u := new(models.User)
	db := config.DB()

	// Binding data
	if err := c.Bind(u); err != nil {
		data := buildErrorJSON(err)
		return c.JSON(http.StatusInternalServerError, data)
	}

	existingUser := new(models.User)
	if err := db.First(&existingUser, id).Error; err != nil {
		data := buildErrorJSON(err)
		return c.JSON(http.StatusNotFound, data)
	}

	existingUser.Username = u.Username
	existingUser.Password = u.Password
	if err := db.Save(&existingUser).Error; err != nil {
		data := buildErrorJSON(err)
		return c.JSON(http.StatusInternalServerError, data)
	}

	response := buildResponseJSON(existingUser)

	return c.JSON(http.StatusOK, response)
}

func GetUser(c echo.Context) error {
	id := c.Param("id")
	user := new(models.User)
	db := config.DB()

	if err := db.First(&user, id).Error; err != nil {
		data := buildErrorJSON(err)
		return c.JSON(http.StatusNotFound, data)
	}

	response := buildResponseJSON(user)

	return c.JSON(http.StatusOK, response)
}

func DeleteUser(c echo.Context) error {
	id := c.Param("id")
	user := new(models.User)
	db := config.DB()

	if err := db.First(&user, id).Error; err != nil {
		data := buildErrorJSON(err)
		return c.JSON(http.StatusNotFound, data)
	}

	user.UserStatus = false
	if err := db.Save(user).Error; err != nil {
		data := buildErrorJSON(err)
		return c.JSON(http.StatusInternalServerError, data)
	}

	response := map[string]interface{}{
		"message": "user ID " + id + " has been deleted",
	}
	return c.JSON(http.StatusOK, response)
}
