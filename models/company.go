package models

type Company struct {
	ID            int    `json:"id"`
	CompanyName   string `json:"company_name"`
	Address       string `json:"address"`
	Email         string `json:"email"`
	CompanyStatus bool   `json:"company_status"`
}
