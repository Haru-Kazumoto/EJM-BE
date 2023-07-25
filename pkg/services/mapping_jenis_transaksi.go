package services

import (
	"EJM/dto"
	"EJM/pkg/models"
	"EJM/pkg/repository"
	"EJM/utils"
	"errors"

	"gorm.io/gorm"
)

type MJenisTransaksiService struct {
	*gorm.DB
	MJenisTransaksiRepository repository.MJenisTransaksiRepository
}

func NewMJenisTransaksiService(service *MJenisTransaksiService) *MJenisTransaksiService {
	return service
}

// new jenis transaksi
func (mJenisTransaksi *MJenisTransaksiService) CreateJenisTransaksi(jenisTransaksiDto *dto.CreateNewJenisTransaksi) (models.MJenisTransaksi, error) {
	mJenisTransaksiRepo := mJenisTransaksi.MJenisTransaksiRepository

	// Cek apakah transaction type sudah ada di database
	TransactionTypeExist := mJenisTransaksiRepo.FindByTransactionType(jenisTransaksiDto.TransactionType)

	if TransactionTypeExist != nil {
		return models.MJenisTransaksi{}, TransactionTypeExist
	}

	data, err := mJenisTransaksiRepo.CreateJenisTransaksi(jenisTransaksiDto)
	if err != nil {
		return models.MJenisTransaksi{}, err
	}

	return data, nil
}


// find all usres [tested]
func (mJenisTransaksi *MJenisTransaksiService) FindJenisTransaksi(jenisTransaksi *dto.GetJenisTransaksi) ([]models.MJenisTransaksi, *models.Paginate, error) {
	pagination := models.Paginate{
		Page:     jenisTransaksi.Page,
		PageSize: jenisTransaksi.PageSize,
	}

	var mJenisTransaksiRepo  repository.MJenisTransaksiRepository = mJenisTransaksi.MJenisTransaksiRepository

	data, meta, err := mJenisTransaksiRepo.FindJenisTransaksi(&pagination, jenisTransaksi.Search, jenisTransaksi.Value) 
	if err != nil {
		return []models.MJenisTransaksi{}, meta, err
	}

	return data, meta, nil
}

// update user [tested]
func (mJenisTransaksi *MJenisTransaksiService) UpdateJenisTransaksi(id uint, jenisTransaksiDto *dto.UpdateJenisTransaksi) error {
	var mJenisTransaksiRepo repository.MJenisTransaksiRepository
	mJenisTransaksiRepo = mJenisTransaksi.MJenisTransaksiRepository

	_, err := mJenisTransaksiRepo.FindJenisTransaksiById(jenisTransaksiDto.ID)

	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			// return utils.ErrJenisTransaksiNotExists
		} else {
			return err
		}
	}

	_, errFind := mJenisTransaksiRepo.FindJenisTransaksiById(id)
	if errFind != nil { // kalo user gak ada
		// return utils.ErrJenisTransaksiNotFound
	}

	return mJenisTransaksiRepo.UpdateJenisTransaksi(id, jenisTransaksiDto)
}


// delete mappingCode [tested]
func (mJenisTransaksi *MJenisTransaksiService) DeleteJenisTransaksi(id uint) error {
	var mJenisTransaksiRepo repository.MJenisTransaksiRepository = mJenisTransaksi.MJenisTransaksiRepository

	// cari id dulu
	_, err := mJenisTransaksiRepo.FindJenisTransaksiById(id)
	if err != nil {
		return utils.ErrMappingCodeNotFound
	}

	// delete
	return mJenisTransaksiRepo.DeleteJenisTransaksi(id)
}

// Find mapping code By Id
func (mJenisTransaksi *MJenisTransaksiService) FindJenisTransaksiById(id uint) (models.MJenisTransaksi, error) {
	var mJenisTransaksiRepo repository.MJenisTransaksiRepository = mJenisTransaksi.MJenisTransaksiRepository

	data, err := mJenisTransaksiRepo.FindJenisTransaksiById(id)

	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			// ID not found in the database
			return models.MJenisTransaksi{}, utils.ErrJenisTransaksiNotFound
		}
		return models.MJenisTransaksi{}, err
	}

	return data, nil
}

