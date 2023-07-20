package repository

import (
	"EJM/dto"
	"EJM/pkg/models"
	"EJM/utils"
	"errors"
	"strings"

	"gorm.io/gorm"
)

type ListOpCodeRepository interface{
	FindListOpCode(pagination *models.Paginate, search string, value string) ([]models.ListOpCode, *models.Paginate, error)
	FindListOpCodeById(id uint) (models.ListOpCode, error)
	CreateListOpCode(listOpCode *dto.CreateListOpCode) (models.ListOpCode, error)
	UpdateListOpCode(id uint, listOpCode *dto.UpdateListOpCode) error
	DeleteListOpCode(id uint) error
	GetTransactionDetailsByID(tipeTransaksiID uint) (string, error)
}

type ListOpCode struct {
	db *gorm.DB
	db2 *gorm.DB
}

func NewListOpCodeRepository(db *gorm.DB) *ListOpCode {
	return &ListOpCode{db: db, db2: db}
}

func (listOpCodeObject *ListOpCode) Begin(tx *gorm.DB) {
	listOpCodeObject.db = tx.Begin()
}

func (listOpCodeObject *ListOpCode) Rollback()  {
	listOpCodeObject.db.Rollback()

	listOpCodeObject.db.Commit()
}

func (listOpCode *ListOpCode) ListOpCodeModel() (tx *gorm.DB) {
	return listOpCode.db.Model(&models.ListOpCode{})
}

// find all list op code paginated
func (listOpCodeObject *ListOpCode) FindListOpCode(pagination *models.Paginate, search, value string) ([]models.ListOpCode, *models.Paginate, error) {
	var listOpCodes []models.ListOpCode
	data := listOpCodeObject.ListOpCodeModel().
		Count(&pagination.Total)

	if search != "" {
		data.Where("lower(list_op_codes.code) like ?", "%"+strings.ToLower(search)+"%").Count(&pagination.Total)
	}

	if value != "" {
		data.Order("list_op_codes.id = " + value + "desc")
	}

	// search data
	data.Scopes(pagination.Pagination()).Preload("TipeTransaksi", func(tx *gorm.DB) *gorm.DB {
		return tx.Select("id", "transaction_type", "transaction_group")
	}).Find(&listOpCodes)
	// checking error
	if err := data.Error; err != nil {
		return []models.ListOpCode{}, pagination, err
	}

	return listOpCodes, pagination, nil
}

// Di dalam repository.ListOpCodeRepository
// type TransactionDetails struct {
// 	TransactionType   string
// 	TransactionGroup  string
// }

func (listOpCode *ListOpCode) GetTransactionDetailsByID(tipeTransaksiID uint) (string, error) {
	var details string
	err := listOpCode.db.Model(&models.MJenisTransaksi{}).
		Select("transaction_group").
		Where("id = ?", tipeTransaksiID).
		Scan(&details).Error
	if err != nil {
		return "", err
	}

	return details, nil
}

// find by id
func (listOpCodeObject *ListOpCode) FindListOpCodeById(id uint) (models.ListOpCode, error) {
	findId := models.ListOpCode{
		BaseModel: models.BaseModel{
			ID: id,
		},
	}
	// you should definition ListOpCode() in function FindListOpCode
	listOpCodeModel := listOpCodeObject.db.Model(&models.ListOpCode{})

	err := listOpCodeModel.First(&findId, "id = ?", id).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return models.ListOpCode{}, utils.ErrListOpCodeNotFound
		}
		return models.ListOpCode{}, err
	}

	return findId, nil
}

// create list op code
func (listOpCode *ListOpCode) CreateListOpCode(list_op_code *dto.CreateListOpCode) (models.ListOpCode, error)  {
	listOpCodeModel := models.ListOpCode{
		OPCode: list_op_code.OPCode,
		ModelMesin: list_op_code.ModelMesin,
		TipeTransaksiID: list_op_code.TipeTransaksiID,
	}

	err := listOpCode.db.Debug().Create(&listOpCodeModel).Error

	if err != nil {
		return listOpCodeModel, nil
	}

	return listOpCodeModel, nil
}
// update list op code
func (listOpCodeObject *ListOpCode) UpdateListOpCode(id uint, listOpCode *dto.UpdateListOpCode) error  {
	update := listOpCodeObject.ListOpCodeModel().Where("id = ?", id).Updates(models.ListOpCode{
		OPCode: listOpCode.OPCode,
		ModelMesin: listOpCode.ModelMesin,
		TipeTransaksiID: listOpCode.TipeTransaksiID,
	})

	if err := update.Error; err != nil {
		return err
	}

	return nil
}

// delete list op code
func (listOpCodeObject *ListOpCode) DeleteListOpCode(id uint) error  {
	deleteListOpCode := listOpCodeObject.ListOpCodeModel().Where("id = ?", id).Delete(&models.ListOpCode{})

	if err := deleteListOpCode.Error
	err != nil {
		return err
	}

	return nil
}