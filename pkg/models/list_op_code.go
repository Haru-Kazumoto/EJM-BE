package models

type ListOpCode struct {
	BaseModel
	OPCode string `json:"opCode" form:"opCode" gorm:"uniqueIndex" validate:"required"`
	ModelMesin string `json:"modelMesin" form:"modelMesin" validate:"required"`
	TipeTransaksiID uint `json:"tipeTransaksiid" form:"tipeTransaksiid" gorm:"uniqueIndex" validate:"required"`
	TipeTransaksi     MJenisTransaksi   `gorm:"foreignKey:TipeTransaksiID" json:"-"` //One to Many

}
