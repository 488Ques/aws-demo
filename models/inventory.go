package models

type Inventory struct {
	ID              int    `json:"id"`
	TruckID         int    `json:"truck_id"`
	ProductName     string `json:"product_name"`
	ProductQuantity int    `json:"product_quantity"`
	MinimumQuantity int    `json:"minimum_quantity"`
	CompanyID       int    `json:"company_id"`
	InventoryStatus bool   `json:"inventory_status"`
}

func (Inventory) TableName() string {
	return "inventory"
}
