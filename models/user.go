package models

type User struct {
	ID         int    `json:"id"`
	StaffID    *int   `json:"staff_id"`
	Username   string `json:"username"`
	Password   string `json:"password"`
	CompanyID  *int   `json:"company_id"`
	UserStatus bool   `json:"user_status"`
}
