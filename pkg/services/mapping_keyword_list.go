package services

import (
	"EJM/dto"
	"EJM/pkg/models"
	"EJM/pkg/repository"
	"EJM/utils"
	"errors"

	"gorm.io/gorm"
)

type MappingKeywordListService struct{
	*gorm.DB
	MappingKeywordListRepository repository.MappingKeywordListRepository
}

type IMappingKeywordListService interface {
	FindMappingkeywordlist(mappingKeywordList *dto.GetMappingkeywordlist) ([]models.MappingKeywordList, *models.Paginate, error)
	FindMappingkeywordlistById(id uint)(models.MappingKeywordList, error)
	CreateMappingkeywordlist(mappingKeywordList *dto.CreateMappingkeywordlist) (models.MappingKeywordList, error)
	UpdateMappingkeywordlist(id uint, mappingKeywordList *dto.UpdateMappingkeywordlist) error
	DeleteMappingkeywordlist(id uint) error
}

func NewMappingKeywordListService(constructor *MappingKeywordListService) *MappingKeywordListService{
	return constructor
}

//new mapping keyword list
func (mappingKeywordList *MappingKeywordListService) CreateMappingKeywordList(mappingKeywordListDto *dto.CreateMappingkeywordlist) (models.MappingKeywordList, error){
	mappingKeywordLists := mappingKeywordList.MappingKeywordListRepository

	KeywordIsExist := mappingKeywordLists.FindMappingKeywordListByKeyword(mappingKeywordListDto.Keyword)

	if KeywordIsExist != nil {
		return models.MappingKeywordList{}, KeywordIsExist
	}
	data, err := mappingKeywordLists.CreateMappingKeywordList(mappingKeywordListDto)
	if err != nil {
		return models.MappingKeywordList{}, err
	}

	return data, nil
}

// find mapping keyword list
func (mappingKeywordList *MappingKeywordListService) FindMappingkeywordlist(mappingKeywordLists *dto.GetMappingkeywordlist) ([]models.MappingKeywordList, *models.Paginate, error) {
	pagination := models.Paginate{
		Page: mappingKeywordLists.Page,
		PageSize: mappingKeywordLists.PageSize,
	}

	var mappingKewordListRepo repository.MappingKeywordListRepository = mappingKeywordList.MappingKeywordListRepository

	data, meta, err := mappingKewordListRepo.FindMappingkeywordlist(&pagination, mappingKeywordLists.Search, mappingKeywordLists.Value,)
	if err != nil {
		return []models.MappingKeywordList{}, meta, err
	}

	return data, meta, nil
}

// update mapping keyword list
func (mappingKeywordList *MappingKeywordListService) UpdateMappingkeywordlist(id uint, mapping_keyword_list *dto.UpdateMappingkeywordlist) error {
	var mappingKeywordListRepo repository.MappingKeywordListRepository
	mappingKeywordListRepo = mappingKeywordList.MappingKeywordListRepository

	_, err := mappingKeywordListRepo.FindMappingkeywordlistById(mapping_keyword_list.ID)

	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {

		} else {
			return err
		}
	}

	_, errFindMappingKeywordList := mappingKeywordListRepo.FindMappingkeywordlistById(id)
	if errFindMappingKeywordList != nil {
		// return utils.ErrUserNotFound
	}

	return mappingKeywordListRepo.UpdateMappingkeywordlist(id, mapping_keyword_list)
}

// delete mapping keyword list
func (mappingKeywordList *MappingKeywordListService) DeleteMappingkeywordlist(id uint) error {
	var mappingKewordListRepo repository.MappingKeywordListRepository
	mappingKewordListRepo = mappingKeywordList.MappingKeywordListRepository

	// find id
	_, err := mappingKewordListRepo.FindMappingkeywordlistById(id)
	if err != nil {
		return utils.ErrUserNotFound
	}

	//delete
	return mappingKewordListRepo.DeleteMappingkeywordlist(id)
}

// find mapping keyword list by id
func (mappingKeywordList *MappingKeywordListService) FindMappingkeywordlistById(id uint) (models.MappingKeywordList, error){
	var mapppingKeywordListRepo repository.MappingKeywordListRepository
	mapppingKeywordListRepo = mappingKeywordList.MappingKeywordListRepository

	data, err := mapppingKeywordListRepo.FindMappingkeywordlistById(id)

	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			// ID not found in the database
			return models.MappingKeywordList{}, utils.ErrMappingKeywordListNotFound
		}
		return models.MappingKeywordList{}, err
	}
	return data, nil

}