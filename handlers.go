package main

import (
	"fmt"
	"net/http"

	"github.com/488Ques/aws-demo/controllers"
	"github.com/gorilla/sessions"
	"github.com/labstack/echo-contrib/session"
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
		return c.String(http.StatusUnauthorized, err.Error())
	}

	// Save user ID to session
	sess, err := session.Get("session", c)
	if err != nil {
		return err
	}
	sess.Options = &sessions.Options{
		Path:     "/",
		MaxAge:   86400, // = 1 day
		HttpOnly: true,
	}
	sess.Values["authUserID"] = id
	sess.Save(c.Request(), c.Response())

	return c.String(http.StatusOK, fmt.Sprintf("User ID is %d\n", id))
}
