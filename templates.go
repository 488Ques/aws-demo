package main

import (
	"github.com/488Ques/aws-demo/models"
	"github.com/labstack/echo-contrib/session"
	"github.com/labstack/echo/v4"
)

type templateData struct {
	IsAuthenticated bool
	Inventory       *[]models.Inventory
	Product         *models.Inventory
}

func addDefaultData(td *templateData, c echo.Context) error {
	isAuth, err := isAuthenticated(c)
	if err != nil {
		return err
	}
	td.IsAuthenticated = isAuth
	return nil
}

func isAuthenticated(c echo.Context) (bool, error) {
	sess, err := session.Get("session", c)
	if err != nil {
		return false, err
	}

	if sess.Values["user_id"] == nil {
		return false, nil
	}
	return true, nil
}
