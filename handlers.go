package main

import (
	"fmt"
	"net/http"

	"github.com/488Ques/aws-demo/controllers"
	"github.com/labstack/echo/v4"
)

func HomeHandler(c echo.Context) error {
	return c.Render(http.StatusOK, "home.html", nil)
}

func LoginForm(c echo.Context) error {
	return c.Render(http.StatusOK, "login.html", nil)
}

func InventoryHandler(c echo.Context) error {
	return c.Render(http.StatusOK, "inventory.html", nil)
}

func LoginUser(c echo.Context) error {
	username := c.FormValue("username")
	password := c.FormValue("password")

	id, err := controllers.Authenticate(username, password)
	if err != nil {
		return c.String(http.StatusBadRequest, err.Error())
	}

	return c.String(http.StatusOK, fmt.Sprintf("User ID is %d\n", id))
}
