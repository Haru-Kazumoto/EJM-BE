package repository

import (
	"EJM/dto"
	"EJM/pkg/models"
	// "EJM/utils"
	// "errors"
	"strings"

	"gorm.io/gorm"
)

type MappingCodeRepository interface {
	// TransactionRepository
	FindMappingCodes(pagination *models.Paginate, search string, value string) ([]models.MappingCode, *models.Paginate, error)
	FindMappingCodeById(id uint) (models.MappingCode, error)
	FindMappingCodeByDefinition(definition string) error
	CreateMappingCode(mappingCode *dto.CreateNewMappingCode) (models.MappingCode, error)
	UpdateMappingCode(id uint, mappingCode *dto.UpdateMappingCode) error
	DeleteMappingCode(id uint) error
}

type MappingCode struct {
	db  *gorm.DB
	db2 *gorm.DB
}

func NewMappingCodeRepository(db *gorm.DB) *MappingCode {
	return &MappingCode{db: db, db2: db}
}

func (mappingCodeObject *MappingCode) Begin(tx *gorm.DB) {
	mappingCodeObject.db = tx.Begin()
}

func (mappingCodeObject *MappingCode) Rollback() {
	mappingCodeObject.db.Rollback()

	// after commit transaction, we have to rollback transaction
	mappingCodeObject.db = mappingCodeObject.db2
}

func (mappingCodeObject *MappingCode) Commit() {
	mappingCodeObject.db.Commit()

	// after commit transaction, we have to rollback transaction
	mappingCodeObject.db = mappingCodeObject.db2
}

func (mappingCodeObject *MappingCode) MappingCodeModel() (tx *gorm.DB) {
	return mappingCodeObject.db.Model(&models.MappingCode{})
}

// find all mapping codes paginated
func (mappingCodeObject *MappingCode) FindMappingCodes(pagination *models.Paginate, search, value string) ([]models.MappingCode, *models.Paginate, error) {
	var mappingCodes []models.MappingCode
	data := mappingCodeObject.MappingCodeModel().
			Count(&pagination.Total)

	if search != "" {
		data.Where("lower(mappingCodes.code) like ? ", "%"+strings.ToLower(search)+"%").Count(&pagination.Total)
	}

	if value != "" {
		data.Order("mappingCodes.id = " + value + " desc")
	}

	// cari data
	data.Scopes(pagination.Pagination()).Debug().
	Find(&mappingCodes)

	// checking errors
	if err := data.Error; err != nil {
		return []models.MappingCode{}, pagination, err
	}

	return mappingCodes, pagination, nil
}

// find by id
func (mappingCodeObject *MappingCode) FindMappingCodeById(id uint) (models.MappingCode, error) {
	findId := models.MappingCode{
		BaseModel: models.BaseModel{
			ID: id,
		},
	}
	if err := mappingCodeObject.MappingCodeModel().First(&findId, "id = ?", id).Error; err != nil {
		return models.MappingCode{}, err
	}

	return findId, nil
}

// find by definition
func (mappingCodeObject *MappingCode) FindMappingCodeByDefinition(definition string) error {
	mappingCode := models.MappingCode{}

	if err := mappingCodeObject.MappingCodeModel().
		First(&mappingCode, "definition = ?", definition).Error; err != nil {
		return err
	}

	return nil
}

// create mapping code
func (mappingCode *MappingCode) CreateMappingCode(mapping_code *dto.CreateNewMappingCode) (models.MappingCode, error) {
	// otp, _ := utils.GenerateOTP(12)
	mappingCodeModel := models.MappingCode{
		Code:             mapping_code.Code,
		Definition:           mapping_code.Definition,
		Status:          mapping_code.Status,
		Priority:          mapping_code.Priority,
		IsActive:       mapping_code.IsActive,
	}
	err := mappingCode.db.Debug().Create(&mappingCodeModel).Error

	if err != nil {
		return mappingCodeModel, err
	}
	return mappingCodeModel, nil
}



// update mapping code
func (mappingCodeObject *MappingCode) UpdateMappingCode(id uint, mappingCode *dto.UpdateMappingCode) error {
	// cariPin, _ := mappingCodeObject.FindById(id)
	// otp, _ := utils.GenerateOTP(12)
	// var pin *string
	// if cariPin.Pin == nil {
	// 	// pin = &otp
	// } else {
	// 	pin = cariPin.Pin
	// }

	update := mappingCodeObject.MappingCodeModel().Where("mappingCodes.id = ?", id).Updates(models.MappingCode{
		Code:        mappingCode.Code,
		Definition:           mappingCode.Definition,
		Status:          mappingCode.Status,
		Priority:          mappingCode.Priority,
		IsActive:       mappingCode.IsActive,
	})

	if err := update.Error; err != nil {
		return err
	}

	return nil
}

// delete mapping code
func (mappingCodeObject *MappingCode) DeleteMappingCode(id uint) error {
	deleteMappingCode := mappingCodeObject.MappingCodeModel().Where("mappingCodes.id = ?", id).Delete(&models.MappingCode{})

	if err := deleteMappingCode.Error; err != nil {
		return err
	}

	return nil
}



// Find By Name
// func (mappingCodeObject *MappingCode) FindByCode(code string) (models.MappingCode, error) {
// 	findMappingCode := models.MappingCode{
// 		Code: code,
// 	}
// 	if err := mappingCodeObject.MappingCodeModel().First(&findMappingCode, "mappingCode.code LIKE ?", "%"+code+"%").Error; err != nil {
// 		return models.MappingCode{}, err
// 	}

// 	return findMappingCode, nil
// }







