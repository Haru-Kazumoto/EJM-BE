package repository

import (
	"EJM/dto"
	"EJM/pkg/models"
	"EJM/utils"
	"errors"
	"strings"

	"gorm.io/gorm"
)

type BinKartuRepository interface {
	FindBinKartu(pagination *models.Paginate, search string,value string) ([]models.BinKartu, *models.Paginate, error)
	FindBinKartuById(id uint) (models.BinKartu, error)
	FindByPrefixNo(prefixNo uint) error
	CreateBinKartu(binKartuDto *dto.CreateNewBinKartu) (models.BinKartu, error)
	UpdateBinKartu(id uint, binKartuDto *dto.UpdateBinKartu) error
	DeleteBinKartu(id uint) error
}

type BinKartu struct {
	db *gorm.DB
}

func NewBinKartuRepository(db *gorm.DB) *BinKartu {
	return &BinKartu{db: db}
}

func (binKartuObject *BinKartu) BinKartuModel() (tx *gorm.DB) {
	return binKartuObject.db.Model(&models.BinKartu{})
}

//find bin kartu paginate
func (binKartuObject *BinKartu) FindBinKartu(pagination *models.Paginate, search string,value string) ([]models.BinKartu, *models.Paginate, error) {
	var binKartu []models.BinKartu
	data := binKartuObject.BinKartuModel().
		Count(&pagination.Total)

	if search != "" {
		data.Where("lower(bin_kartus.bank_name) like ? OR bin_kartus.prefix_no = ?", "%"+strings.ToLower(search)+"%", search).Count(&pagination.Total)
	}

	// if usingActive {
	// 	data.Where("bin_kartus.is_active", true).Count(&pagination.Total)
	// }

	if value != "" {
		data.Order("bin_kartus.id = " + value + " desc")
	}

	// cari data
	data.Scopes(pagination.Pagination()).Debug().
		Find(&binKartu)

	// checking errors
	if err := data.Error; err != nil {
		return []models.BinKartu{}, pagination, err
	}

	return binKartu, pagination, nil
}

//find by id
func (binKartuObject *BinKartu) FindBinKartuById(id uint) (models.BinKartu, error) {
	findId := models.BinKartu{
		BaseModel: models.BaseModel{
			ID: id,
		},
	}

	binKartuModel  := binKartuObject.BinKartuModel()
	err := binKartuModel.First(&findId, "id = ?", id).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			// ID not found in the database
			return models.BinKartu{}, utils.ErrBinKartuNotFound
		}
		return models.BinKartu{}, err
	}

	return findId, nil
}


func (binKartuObject *BinKartu) FindByPrefixNo(prefixNo uint) error {
	jenisTransaksi := models.BinKartu{}

	if err := binKartuObject.BinKartuModel().
		First(&jenisTransaksi, "prefixNo = ?", prefixNo).Error; err == nil {
		return utils.ErrPrefixNoAlreadyExists
	}

	return nil
}

// create bin kartu
func (binKartuObject *BinKartu) CreateBinKartu(binKartuDto *dto.CreateNewBinKartu) (models.BinKartu, error) {
	binKartuModel := models.BinKartu{
		PrefixNo: binKartuDto.PrefixNo,
		BankName: binKartuDto.BankName,
	}

	err := binKartuObject.db.Debug().Create(&binKartuModel).Error

	if err != nil {
		return binKartuModel, err
	}

	return binKartuModel, nil
}

// update bin kartu
func (binKartuObject *BinKartu) UpdateBinKartu(id uint, binKartuDto *dto.UpdateBinKartu) error {
	binKartuModel := binKartuObject.BinKartuModel().Where("id = ?", id).Updates(models.BinKartu{
		PrefixNo: binKartuDto.PrefixNo,
		BankName: binKartuDto.BankName,
	})

	if err := binKartuModel.Error; err != nil {
		return err
	}

	return nil
}

//delete
func (binKartuObject *BinKartu) DeleteBinKartu(id uint) error {
	deleteBinKartu := binKartuObject.BinKartuModel().Where("id = ?", id).Delete(&models.BinKartu{})

	if err := deleteBinKartu.Error
	err != nil {
		return err
	}

	return nil
}