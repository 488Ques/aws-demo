package main

import (
	"github.com/488Ques/aws-demo/controllers"
	"github.com/labstack/echo/v4"
)

func Routes(e *echo.Echo) {
	// UI routes
	e.GET("/", HomeHandler)
	e.GET("/login", LoginHandler)
	e.GET("/inventory", InventoryHandler)

	// User routes
	userRoute := e.Group("/user")
	userRoute.GET("/", controllers.AllUsers)
	userRoute.POST("/", controllers.CreateUser)
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

	// TODO: Delete this in production
	// Example routes
	bookRoute := e.Group("/book")
	bookRoute.POST("/", controllers.CreateBook)
	bookRoute.GET("/:id", controllers.GetBook)
	bookRoute.PUT("/:id", controllers.UpdateBook)
	bookRoute.DELETE("/:id", controllers.DeleteBook)
}
