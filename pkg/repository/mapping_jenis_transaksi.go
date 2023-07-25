package repository

import (
	"EJM/dto"
	"EJM/pkg/models"
	"EJM/utils"
	"errors"
	"strings"

	"gorm.io/gorm"
)

type MJenisTransaksiRepository interface {
	FindJenisTransaksi(pagination *models.Paginate, search string, value string) ([]models.MJenisTransaksi, *models.Paginate, error)
	FindJenisTransaksiById(id uint) (models.MJenisTransaksi, error)
	FindByTransactionType(transactionType string) error
	CreateJenisTransaksi(mJenisTransaksi *dto.CreateNewJenisTransaksi) (models.MJenisTransaksi, error)
	UpdateJenisTransaksi(id uint, mJenisTransaksi *dto.UpdateJenisTransaksi) error
	DeleteJenisTransaksi(id uint) error
}

type MJenisTransaksi struct {
	db *gorm.DB
}

func NewMJenisTransaksiRepository(db *gorm.DB) *MJenisTransaksi {
	return &MJenisTransaksi{db: db}
}

func (mJenisTransaksiObject *MJenisTransaksi) MJenisTransaksiModel() (tx *gorm.DB) {
	return mJenisTransaksiObject.db.Model(&models.MJenisTransaksi{})
}

//find jenis transaksi paginate
func (mJenisTransaksiObject *MJenisTransaksi) FindJenisTransaksi(pagination *models.Paginate, search string, value string) ([]models.MJenisTransaksi, *models.Paginate, error) {
	var jenisTransaksi []models.MJenisTransaksi
	data := mJenisTransaksiObject.MJenisTransaksiModel().
		Count(&pagination.Total)

	if search != "" {
		data.Where("lower(m_jenis_transaksis.transaction_type) like ? OR lower(m_jenis_transaksis.transaction_group) like ?", "%"+strings.ToLower(search)+"%","%"+strings.ToLower(search)+"%").Count(&pagination.Total)
	}



	if value != "" {
		data.Order("m_jenis_transaksis.id = " + value + " desc")
	}

	// cari data
	data.Scopes(pagination.Pagination()).Debug().
		Find(&jenisTransaksi)

	// checking errors
	if err := data.Error; err != nil {
		return []models.MJenisTransaksi{}, pagination, err
	}

	return jenisTransaksi, pagination, nil
}

//find by id
func (mJenisTransaksiObject *MJenisTransaksi) FindJenisTransaksiById(id uint) (models.MJenisTransaksi, error) {
	findId := models.MJenisTransaksi{
		BaseModel: models.BaseModel{
			ID: id,
		},
	}
	// Definisikan mJenisTransaksiModel() di dalam fungsi FindMappingCodeById
	mJenisTransaksiModel  := mJenisTransaksiObject.MJenisTransaksiModel()
	err := mJenisTransaksiModel.First(&findId, "id = ?", id).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			// ID not found in the database
			return models.MJenisTransaksi{}, utils.ErrMappingCodeNotFound
		}
		return models.MJenisTransaksi{}, err
	}

	return findId, nil
}

func (mJenisTransaksiObject *MJenisTransaksi) FindByTransactionType(transactionType string) error {
	jenisTransaksi := models.MJenisTransaksi{}

	if err := mJenisTransaksiObject.MJenisTransaksiModel().
		First(&jenisTransaksi, "transactionType = ?", transactionType).Error; err == nil {
		return utils.ErrTransaksiAlreadyExists
	}

	return nil
}

// create jenis transaksi
func (mJenisTransaksiObject *MJenisTransaksi) CreateJenisTransaksi(jenisTransaksiDto *dto.CreateNewJenisTransaksi) (models.MJenisTransaksi, error) {
	mJenisTransaksiModel := models.MJenisTransaksi{
		TransactionType: jenisTransaksiDto.TransactionType,
		TransactionGroup: jenisTransaksiDto.TransactionGroup,
	}

	err := mJenisTransaksiObject.db.Debug().Create(&mJenisTransaksiModel).Error

	if err != nil {
		return mJenisTransaksiModel, err
	}

	return mJenisTransaksiModel, nil
}

// update mapping code
func (mJenisTransaksiObject *MJenisTransaksi) UpdateJenisTransaksi(id uint, jenisTransaksiDto *dto.UpdateJenisTransaksi) error {
	mJenisTransaksiModel := mJenisTransaksiObject.MJenisTransaksiModel().Where("id = ?", id).Updates(models.MJenisTransaksi{
		TransactionType: jenisTransaksiDto.TransactionType,
		TransactionGroup: jenisTransaksiDto.TransactionGroup,
	})

	if err := mJenisTransaksiModel.Error; err != nil {
		return err
	}

	return nil
}

//delete
func (mJenisTransaksiObject *MJenisTransaksi) DeleteJenisTransaksi(id uint) error {
	deleteJenisTransaksi := mJenisTransaksiObject.MJenisTransaksiModel().Where("id = ?", id).Delete(&models.MJenisTransaksi{})

	if err := deleteJenisTransaksi.Error
	err != nil {
		return err
	}

	return nil
}