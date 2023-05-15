package models

type Staff struct {
	ID          int    `json:"id"`
	TruckID     int    `json:"truck_id"`
	CompanyID   int    `json:"company_id"`
	StaffName   string `json:"staff_name"`
	PhoneNumber string `json:"phone_number"`
	IsManager   bool   `json:"is_manager"`
	StaffStatus bool   `json:"staff_status"`
}
