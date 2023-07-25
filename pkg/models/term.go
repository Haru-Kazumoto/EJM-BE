package models

type Term struct {
	BaseModel
	ATMKey string `json:"atm_key" form:"atm_key" gorm:"uniqueIndex" validate:"required"` 
	TermID string `json:"term_id" form:"term_id" gorm:"uniqueIndex" validate:"required"` 
	TermType string `json:"term_type" form:"term_type" validate:"required"` 
	Merk string `json:"merk" form:"merk" validate:"required"` 
	ManagedBy string `json:"managed_by" form:"managed_by" validate:"required"` 
	Jarkom string `json:"jarkom" form:"jarkom" validate:"required"` 
	Region string `json:"region" form:"region" validate:"required"` 
	Area string `json:"area" form:"area" validate:"required"` 
	CardRetainedTes string `json:"cardretained" form:"cardretained" validate:"required"` 
}