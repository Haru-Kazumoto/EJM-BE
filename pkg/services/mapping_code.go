package services

import (
	// "EJM/config"
	"EJM/dto"
	"EJM/utils"
	"errors"

	// "EJM/internal/logs"
	"EJM/pkg/models"
	"EJM/pkg/repository"

	// "EJM/utils"
	// "errors"
	"gorm.io/gorm"
)

type MappingCodeService struct {
	*gorm.DB
	MappingCodeRepository repository.MappingCodeRepository
}

func NewMappingCodeService(service *MappingCodeService) *MappingCodeService {
	return service
}




// type IMappingCodeService interface {
// 	FindMappingCodes(mappingCodes *dto.GetMappingCodes) ([]models.MappingCode, *models.Paginate, error)
// 	FindMappingCodeById(id uint) (models.MappingCode, error)
// 	// FindMappingCodeByDefinition(definition string) error
// 	CreateMappingCode(mappingCode *dto.CreateNewMappingCode) (models.MappingCode, error)
// 	UpdateMappingCode(id uint, mappingCode *dto.UpdateMappingCode) error
// 	DeleteMappingCode(id uint) error
// }




// 

// new mapping code 
func (mappingCode *MappingCodeService) CreateMappingCode(mappingCodeDto *dto.CreateNewMappingCode) (models.MappingCode, error) {
	mappingCodes := mappingCode.MappingCodeRepository

	// Cek apakah code sudah ada di database
	CodeIsExist := mappingCodes.FindMappingCodeByCode(mappingCodeDto.Code)

	if CodeIsExist != nil {
		return models.MappingCode{}, CodeIsExist
	}

	data, err := mappingCodes.CreateMappingCode(mappingCodeDto)
	if err != nil {
		return models.MappingCode{}, err
	}

	return data, nil
}


// find all usres [tested]
func (mappingCode *MappingCodeService) FindMappingCodes(mappingCodes *dto.GetMappingCodes) ([]models.MappingCode, *models.Paginate, error) {
	pagination := models.Paginate{
		Page:     mappingCodes.Page,
		PageSize: mappingCodes.PageSize,
	}

	var mappingCodeRepo repository.MappingCodeRepository = mappingCode.MappingCodeRepository

	data, meta, err := mappingCodeRepo.FindMappingCodes(&pagination,mappingCodes.Search,mappingCodes.UsingActive,  mappingCodes.Value) 
	if err != nil {
		return []models.MappingCode{}, meta, err
	}

	return data, meta, nil
}

// update code [tested]
func ( mappingCode *MappingCodeService) UpdateMappingCode(id uint, mapping_code *dto.UpdateMappingCode) error {
	var mappingCodeRepo repository.MappingCodeRepository
	mappingCodeRepo = mappingCode.MappingCodeRepository

	_, err := mappingCodeRepo.FindMappingCodeById(mapping_code.ID)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			// return utils.ErrRoleNotExists
		} else {
			return err
		}
	}

	_, errFindMappingCode := mappingCodeRepo.FindMappingCodeById(id)
	if errFindMappingCode != nil { 
		// return utils.ErrUserNotFound
	}

	return mappingCodeRepo.UpdateMappingCode(id, mapping_code)
}


// delete mappingCode [tested]
func (mappingCode *MappingCodeService) DeleteMappingCode(id uint) error {
	var mappingCodeRepo repository.MappingCodeRepository = mappingCode.MappingCodeRepository

	// cari id dulu
	_, err := mappingCodeRepo.FindMappingCodeById(id)
	if err != nil {
		return utils.ErrMappingCodeNotFound
	}

	// delete
	return mappingCodeRepo.DeleteMappingCode(id)
}

// Find mapping code By Id
func (mappingCode *MappingCodeService) FindMappingCodeById(id uint) (models.MappingCode, error) {
	var mappingCodeRepo repository.MappingCodeRepository = mappingCode.MappingCodeRepository

	data, err := mappingCodeRepo.FindMappingCodeById(id)

	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			// ID not found in the database
			return models.MappingCode{}, utils.ErrMappingCodeNotFound
		}
		return models.MappingCode{}, err
	}

	return data, nil
}

