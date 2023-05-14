package models

type Truck struct {
	ID          int    `json:"id"`
	CompanyID   *int   `json:"company_id"`
	TruckName   string `json:"truck_name"`
	Location    string `json:"location"`
	TruckStatus bool   `json:"truck_status"`
}
