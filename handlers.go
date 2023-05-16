package main

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/488Ques/aws-demo/controllers"
	"github.com/488Ques/aws-demo/models"
	"github.com/488Ques/aws-demo/twilio_helper"
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
	sess, _ := session.Get("session", c)
	staffID := sess.Values["staff_id"].(int)
	products := controllers.GetStaffInventory(staffID)

	return c.Render(http.StatusOK, "inventory.html", &templateData{Inventory: &products})
}

func AddProductForm(c echo.Context) error {
	return c.Render(http.StatusOK, "addProduct.html", nil)
}

func AddProduct(c echo.Context) error {
	productName := c.FormValue("product_name")
	productQuantity := c.FormValue("product_quantity")
	minimumQuantity := c.FormValue("minimum_quantity")
	truckID := c.FormValue("truck_id")
	companyID := c.FormValue("company_id")

	quantity, _ := strconv.Atoi(productQuantity)
	minimum, _ := strconv.Atoi(minimumQuantity)
	truckid, _ := strconv.Atoi(truckID)
	companyid, _ := strconv.Atoi(companyID)

	product := &models.Inventory{
		ProductName:     productName,
		ProductQuantity: quantity,
		MinimumQuantity: minimum,
		TruckID:         truckid,
		CompanyID:       companyid,
		InventoryStatus: true,
	}

	_, err := controllers.AddProduct(product)
	if err != nil {
		return c.String(http.StatusInternalServerError, err.Error())
	}

	return c.Redirect(http.StatusSeeOther, "/inventory")
}

func EditProductForm(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return c.String(http.StatusInternalServerError, err.Error())
	}
	product, err := controllers.GetProduct(id)
	if err != nil {
		return c.String(http.StatusNotFound, err.Error())
	}

	return c.Render(http.StatusOK, "editProduct.html", &templateData{Product: product})
}

func EditProduct(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return c.String(http.StatusInternalServerError, err.Error())
	}
	productName := c.FormValue("product_name")
	productQuantity := c.FormValue("product_quantity")
	minimumQuantity := c.FormValue("minimum_quantity")
	truckID := c.FormValue("truck_id")
	companyID := c.FormValue("company_id")

	quantity, _ := strconv.Atoi(productQuantity)
	minimum, _ := strconv.Atoi(minimumQuantity)
	truckid, _ := strconv.Atoi(truckID)
	companyid, _ := strconv.Atoi(companyID)

	product, err := controllers.GetProduct(id)
	if err != nil {
		return c.String(http.StatusNotFound, err.Error())
	}

	product.ProductName = productName
	product.ProductQuantity = quantity
	product.MinimumQuantity = minimum
	product.TruckID = truckid
	product.CompanyID = companyid

	err = controllers.UpdateProduct(product)
	if err != nil {
		return c.String(http.StatusInternalServerError, err.Error())
	}

	if product.ProductQuantity < product.MinimumQuantity {
		twilio_helper.CreateMessage(fmt.Sprintf("Product %s in truck ID %s is going to run out of stock", productName, truckID))
	}

	return c.Redirect(http.StatusSeeOther, "/inventory")
}

func DeleteProduct(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return c.String(http.StatusInternalServerError, err.Error())
	}
	err = controllers.DeleteProduct(id)
	if err != nil {
		return c.String(http.StatusNotFound, err.Error())
	}

	return c.Redirect(http.StatusSeeOther, "/inventory")
}

func LoginUser(c echo.Context) error {
	username := c.FormValue("username")
	password := c.FormValue("password")

	user, err := controllers.Authenticate(username, password)
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
	sess.Values["user_id"] = user.ID
	sess.Values["staff_id"] = user.StaffID
	sess.Values["username"] = user.Username
	sess.Values["company_id"] = user.CompanyID
	sess.Save(c.Request(), c.Response())

	return c.Redirect(http.StatusSeeOther, "/")
}

func LogoutUser(c echo.Context) error {
	sess, err := session.Get("session", c)
	if err != nil {
		return err
	}
	sess.Options = &sessions.Options{
		Path:     "/",
		MaxAge:   -1,
		HttpOnly: true,
	}
	sess.Save(c.Request(), c.Response())

	return c.Redirect(http.StatusSeeOther, "/")
}
