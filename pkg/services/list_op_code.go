package services

import (
	"EJM/dto"
	"EJM/pkg/models"
	"EJM/pkg/repository"
	"EJM/utils"
	"errors"

	"gorm.io/gorm"
)

type ListOpCodeService struct {
	*gorm.DB
	ListOpCodeRepository repository.ListOpCodeRepository
}

func NewListOpCodeService(service *ListOpCodeService) *ListOpCodeService {
	return service
}

// new list op code
func (listOpCode *ListOpCodeService) CreateListOpCode(listOpCodeDto *dto.CreateListOpCode) (models.ListOpCode, error) {
	listOpCodes :=  listOpCode.ListOpCodeRepository
	data, err := listOpCodes.CreateListOpCode(listOpCodeDto)
	if err != nil {
		return models.ListOpCode{}, err
	}

	return data, nil
}

//find list op code
func (listOpCode *ListOpCodeService) FindListOpCode(listOpCodes *dto.GetListOpCode) ([]models.ListOpCode, *models.Paginate, error) {
	pagination := models.Paginate{
		Page:     listOpCodes.Page,
		PageSize: listOpCodes.PageSize,
	}

	var listOpCodeRepo repository.ListOpCodeRepository = listOpCode.ListOpCodeRepository

	data, meta, err := listOpCodeRepo.FindListOpCode(&pagination, listOpCodes.Search, listOpCodes.Value)
	if err != nil {
		return []models.ListOpCode{}, meta, err
	}

	return data, meta, nil
}


// update list op code
func (listOpcode *ListOpCodeService) UpdateListOpCode(id uint, list_op_code *dto.UpdateListOpCode) error  {
	var listOpCodeRepo repository.ListOpCodeRepository
	listOpCodeRepo = listOpcode.ListOpCodeRepository

	_, err := listOpCodeRepo.FindListOpCodeById(list_op_code.ID)

	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound)  {

		} else {
			return err
		}
	}

	_, errFindListOpCode := listOpCodeRepo.FindListOpCodeById(id)
	if errFindListOpCode != nil{

	}

	return listOpCodeRepo.UpdateListOpCode(id, list_op_code)

}

// delete list op code
func (listOpCode *ListOpCodeService) DeleteListOpCode(id uint) error  {
	var listOpCodeRepo repository.ListOpCodeRepository
	listOpCodeRepo = listOpCode.ListOpCodeRepository


	// find id
	_, err := listOpCodeRepo.FindListOpCodeById(id)
	if err != nil {
		return utils.ErrUserNotFound
	}

	//delete
	return listOpCodeRepo.DeleteListOpCode(id)
}


// find list op code by id
func (listOpCode *ListOpCodeService) FindListOpCodeById(id uint) (models.ListOpCode, error)  {
	var listOpCodeRepo repository.ListOpCodeRepository
	listOpCodeRepo = listOpCode.ListOpCodeRepository

	data, err := listOpCodeRepo.FindListOpCodeById(id)

	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return models.ListOpCode{}, utils.ErrListOpCodeNotFound
		}
		return models.ListOpCode{}, err
	}
	return data, nil

}
