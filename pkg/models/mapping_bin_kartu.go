package models

type BinKartu struct {
	BaseModel
	PrefixNo uint `json:"prefix_no" form:"prefix_no" gorm:"uniqueIndex" validate:"required"`
	BankName string `json:"bank_name" form:"bank_name" validate:"required"`
}