package main

import (
	"net/http"

	"github.com/labstack/echo-contrib/session"
	"github.com/labstack/echo/v4"
)

func Routes(e *echo.Echo) {
	// UI routes
	e.GET("/", HomeHandler)
	e.GET("/inventory", InventoryHandler)

	e.GET("/whoami", func(c echo.Context) error {
		sess, err := session.Get("session", c)
		if err != nil {
			return c.String(http.StatusUnauthorized, err.Error())
		}
		userID := sess.Values["user_id"].(int)
		staffID := sess.Values["staff_id"]
		username := sess.Values["username"]
		companyID := sess.Values["company_id"]

		data := map[string]any{
			"user_id":    userID,
			"staff_id":   staffID,
			"username":   username,
			"company_id": companyID,
		}
		return c.JSON(http.StatusOK, data)
	})

	// User routes
	userRoute := e.Group("/user")
	userRoute.GET("/login", LoginForm)
	userRoute.POST("/login", LoginUser)
	userRoute.GET("/logout", LogoutUser)

	// Truck routes
	// truckRoute := e.Group("/truck")

	// Staff routes
	// staffRoute := e.Group("/staff")

	// Inventory routes
	inventoryRoute := e.Group("/inventory")
	inventoryRoute.GET("/add", AddProductForm)
	inventoryRoute.POST("/add", AddProduct)
	inventoryRoute.GET("/:id", EditProductForm)
	inventoryRoute.POST("/:id", EditProduct)
	inventoryRoute.DELETE("/:id", DeleteProduct)

	// Company routes
	// companyRoute := e.Group("/company")
}
