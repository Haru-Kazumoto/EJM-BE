package services

import (
	"EJM/dto"
	"EJM/pkg/models"
	"EJM/pkg/repository"
	"EJM/utils"
	"errors"

	"gorm.io/gorm"
)

type BinKartuService struct {
	*gorm.DB
	BinKartuRepository repository.BinKartuRepository
}

func NewBinKartuService(service *BinKartuService) *BinKartuService {
	return service
}

// new bin kartu
func (binKartu *BinKartuService) CreateBinKartu(binKartuDto *dto.CreateNewBinKartu) (models.BinKartu, error) {
	binKartuRepo := binKartu.BinKartuRepository

	// Cek apakah prefix no sudah ada di database
	PrefixNoExist := binKartuRepo.FindByPrefixNo(binKartuDto.PrefixNo)

	if PrefixNoExist != nil {
		return models.BinKartu{}, PrefixNoExist
	}

	data, err := binKartuRepo.CreateBinKartu(binKartuDto)
	if err != nil {
		return models.BinKartu{}, err
	}

	return data, nil
}


// find all binkartu [tested]
func (binKartu *BinKartuService) FindBinKartu(binKartuDto *dto.GetBinKartu) ([]models.BinKartu, *models.Paginate, error) {
	pagination := models.Paginate{
		Page:     binKartuDto.Page,
		PageSize: binKartuDto.PageSize,
	}

	var binKartuRepo  repository.BinKartuRepository = binKartu.BinKartuRepository

	data, meta, err := binKartuRepo.FindBinKartu(&pagination, binKartuDto.Search, binKartuDto.Value) 
	if err != nil {
		return []models.BinKartu{}, meta, err
	}

	return data, meta, nil
}

// update bin Kartu [tested]
func (binKartu *BinKartuService) UpdateBinKartu(id uint, binKartuDto *dto.UpdateBinKartu) error {
	var binKartuRepo repository.BinKartuRepository
	binKartuRepo = binKartu.BinKartuRepository

	_, err := binKartuRepo.FindBinKartuById(binKartuDto.ID)

	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			// return utils.ErrJenisTransaksiNotExists
		} else {
			return err
		}
	}

	_, errFind := binKartuRepo.FindBinKartuById(id)
	if errFind != nil { // kalo bin gak ada
		// return utils.ErrJenisTransaksiNotFound
	}

	return binKartuRepo.UpdateBinKartu(id, binKartuDto)
}


// delete bin [tested]
func (binKartu *BinKartuService) DeleteBinKartu(id uint) error {
	var binKartuRepo repository.BinKartuRepository = binKartu.BinKartuRepository

	// cari id dulu
	_, err := binKartuRepo.FindBinKartuById(id)
	if err != nil {
		return utils.ErrBinKartuNotFound
	}

	// delete
	return binKartuRepo.DeleteBinKartu(id)
}

// Find bin By Id
func (binKartu *BinKartuService) FindBinKartuById(id uint) (models.BinKartu, error) {
	var binKartuRepo repository.BinKartuRepository = binKartu.BinKartuRepository

	data, err := binKartuRepo.FindBinKartuById(id)

	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			// ID not found in the database
			return models.BinKartu{}, utils.ErrBinKartuNotFound
		}
		return models.BinKartu{}, err
	}

	return data, nil
}