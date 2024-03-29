package controllers

import (
	"net/http"

	"github.com/488Ques/aws-demo/config"
	"github.com/488Ques/aws-demo/models"
	"github.com/labstack/echo/v4"
)

func AllInventory(c echo.Context) error {
	var products []*models.Inventory
	db := config.DB()

	if err := db.Where("inventory_status <> ?", "0").Find(&products).Error; err != nil {
		data := buildErrorJSON(err)
		return c.JSON(http.StatusInternalServerError, data)
	}

	response := buildResponseJSON(products)
	return c.JSON(http.StatusOK, response)
}

func CreateInventory(c echo.Context) error {
	p := new(models.Inventory)
	db := config.DB()

	if err := c.Bind(p); err != nil {
		data := buildErrorJSON(err)
		return c.JSON(http.StatusInternalServerError, data)
	}

	product := &models.Inventory{
		ProductName:     p.ProductName,
		ProductQuantity: p.ProductQuantity,
		MinimumQuantity: p.MinimumQuantity,
		InventoryStatus: true,
	}

	if err := db.Create(&product).Error; err != nil {
		data := buildErrorJSON(err)
		return c.JSON(http.StatusInternalServerError, data)
	}

	response := buildResponseJSON(product)
	return c.JSON(http.StatusOK, response)
}

func UpdateInventory(c echo.Context) error {
	id := c.Param("id")
	p := new(models.Inventory)
	db := config.DB()

	// Binding data
	if err := c.Bind(p); err != nil {
		data := buildErrorJSON(err)
		return c.JSON(http.StatusInternalServerError, data)
	}

	existingProduct := new(models.Inventory)
	if err := db.First(&existingProduct, id).Error; err != nil {
		data := buildErrorJSON(err)
		return c.JSON(http.StatusNotFound, data)
	}

	existingProduct.ProductName = p.ProductName
	existingProduct.ProductQuantity = p.ProductQuantity
	existingProduct.MinimumQuantity = p.MinimumQuantity
	existingProduct.InventoryStatus = p.InventoryStatus
	if err := db.Save(&existingProduct).Error; err != nil {
		data := buildErrorJSON(err)
		return c.JSON(http.StatusInternalServerError, data)
	}

	response := buildResponseJSON(existingProduct)

	return c.JSON(http.StatusOK, response)
}

func GetInventory(c echo.Context) error {
	id := c.Param("id")
	product := new(models.Inventory)
	db := config.DB()

	if err := db.First(&product, id).Error; err != nil {
		data := buildErrorJSON(err)
		return c.JSON(http.StatusNotFound, data)
	}

	response := buildResponseJSON(product)

	return c.JSON(http.StatusOK, response)
}

func DeleteInventory(c echo.Context) error {
	id := c.Param("id")
	product := new(models.Inventory)
	db := config.DB()

	if err := db.First(&product, id).Error; err != nil {
		data := buildErrorJSON(err)
		return c.JSON(http.StatusNotFound, data)
	}

	product.InventoryStatus = false
	if err := db.Save(product).Error; err != nil {
		data := buildErrorJSON(err)
		return c.JSON(http.StatusInternalServerError, data)
	}

	response := map[string]interface{}{
		"message": "Product ID " + id + " has been deleted",
	}
	return c.JSON(http.StatusOK, response)
}

func GetProduct(id int) (*models.Inventory, error) {
	product := new(models.Inventory)
	db := config.DB()

	if err := db.First(&product, id).Error; err != nil {
		return nil, err
	}

	return product, nil
}

func UpdateProduct(p *models.Inventory) error {
	db := config.DB()
	if err := db.Save(p).Error; err != nil {
		return err
	}
	return nil
}

func DeleteProduct(id int) error {
	product, err := GetProduct(id)
	if err != nil {
		return err
	}
	db := config.DB()

	product.InventoryStatus = false
	if err := db.Save(product).Error; err != nil {
		return err
	}

	return nil
}

func AddProduct(p *models.Inventory) (int, error) {
	db := config.DB()

	if err := db.Create(p).Error; err != nil {
		return 0, err
	}

	return p.ID, nil
}

func GetStaffInventory(staffID int) []models.Inventory {
	db := config.DB()
	products := []models.Inventory{}

	db.Model(&models.Inventory{}).Joins("inner join staffs on staffs.truck_id = inventory.truck_id").Where("staffs.id = ? AND inventory_status <> 0", staffID).Find(&products)

	return products
}
