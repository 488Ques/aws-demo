package main

import (
	"github.com/488Ques/aws-demo/controllers"
	"github.com/labstack/echo/v4"
)

func Routes(e *echo.Echo) {
	// UI routes
	e.GET("/", HomeHandler)
	e.GET("/inventory", InventoryHandler)

	// e.GET("/whoami", func(c echo.Context) error {
	// 	sess, err := session.Get("session", c)
	// 	if err != nil {
	// 		return err
	// 	}

	// 	return c.JSON(http.StatusOK, sess.Values["name"])
	// })

	// User routes
	userRoute := e.Group("/user")
	userRoute.GET("/", controllers.AllUsers)
	userRoute.POST("/", controllers.CreateUser)
	userRoute.GET("/login", LoginForm)
	userRoute.POST("/login", LoginUser)
	userRoute.GET("/logout", LogoutUser)
	userRoute.GET("/:id", controllers.GetUser)
	userRoute.PUT("/:id", controllers.UpdateUser)
	userRoute.DELETE("/:id", controllers.DeleteUser)

	// Truck routes
	truckRoute := e.Group("/truck")
	truckRoute.GET("/", controllers.AllTrucks)
	truckRoute.POST("/", controllers.CreateTruck)
	truckRoute.GET("/:id", controllers.GetTruck)
	truckRoute.PUT("/:id", controllers.UpdateTruck)
	truckRoute.DELETE("/:id", controllers.DeleteTruck)

	// Staff routes
	staffRoute := e.Group("/staff")
	staffRoute.GET("/", controllers.AllStaffs)
	staffRoute.POST("/", controllers.CreateStaff)
	staffRoute.GET("/:id", controllers.GetStaff)
	staffRoute.PUT("/:id", controllers.UpdateStaff)
	staffRoute.DELETE("/:id", controllers.DeleteStaff)

	// Inventory routes
	inventoryRoute := e.Group("/inventory")
	inventoryRoute.GET("/", controllers.AllInventory)
	inventoryRoute.POST("/", controllers.CreateInventory)
	inventoryRoute.GET("/:id", controllers.GetInventory)
	inventoryRoute.PUT("/:id", controllers.UpdateInventory)
	inventoryRoute.DELETE("/:id", controllers.DeleteInventory)

	// Company routes
	companyRoute := e.Group("/company")
	companyRoute.GET("/", controllers.AllCompanies)
	companyRoute.POST("/", controllers.CreateCompany)
	companyRoute.GET("/:id", controllers.GetCompany)
	companyRoute.PUT("/:id", controllers.UpdateCompany)
	companyRoute.DELETE("/:id", controllers.DeleteCompany)

	// TODO: Delete this in production
	// Book routes
	bookRoute := e.Group("/book")
	bookRoute.POST("/", controllers.CreateBook)
	bookRoute.GET("/:id", controllers.GetBook)
	bookRoute.PUT("/:id", controllers.UpdateBook)
	bookRoute.DELETE("/:id", controllers.DeleteBook)
}
