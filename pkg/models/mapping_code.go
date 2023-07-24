package models

type MappingCode struct {
	BaseModel
	// ActiveModel
	Code       string `json:"code" form:"code" validate:"required" gorm:"uniqueIndex"`
	Definition string `json:"definition" form:"definition"  validate:"required"`
	Status     StatusEnum `json:"status" form:"status" gorm:"type:status_enum"`
	Priority   int `json:"priority" form:"priority" validate:"required"`
	Active   ActiveEnum `json:"active" form:"active" validate:"required" gorm:"type:isactive_enum"`

}

