package models

type MappingKeywordList struct {
	BaseModel
	MappingKeywordList string `json:"MappingKeywordList" gorm:"not null"`
}
