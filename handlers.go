package main

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

func HomeHandler(c echo.Context) error {
	return c.Render(http.StatusOK, "home.html", nil)
}

func LoginHandler(c echo.Context) error {
	return c.Render(http.StatusOK, "login.html", nil)
}

func InventoryHandler(c echo.Context) error {
	return c.Render(http.StatusOK, "inventory.html", nil)
}
