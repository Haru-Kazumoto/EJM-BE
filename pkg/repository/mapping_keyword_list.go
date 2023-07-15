package repository

import (
	"EJM/dto"
	"EJM/pkg/models"
	"strings"

	"gorm.io/gorm"
)

type MappingKeywordListRepository interface {
	CreateMappingkeywordlist(mappingKeywordList *dto.CreateMappingkeywordlist) (models.MappingKeywordList, error)             				   // [register new user]
	FindMappingkeywordlist(pagination *models.Paginate, search string, value string) ([]models.MappingKeywordList, *models.Paginate, error)
	FindMappingkeywordlistById(id uint) (models.MappingKeywordList, error) // [delete user by id]
	UpdateMappingkeywordlist(id uint, user *dto.UpdateMappingkeywordlist) error                         						  //[update user by id]
	DeleteMappingkeywordlist(id uint) error                         
}

type MappingKeywordList struct {
	db *gorm.DB
	db2 *gorm.DB
}

func NewMappingKeywordLIstRepository(db *gorm.DB) *MappingKeywordList {
	return &MappingKeywordList{db: db, db2: db}
}

func (mappingKeywordListObject *MappingKeywordList) Begin(tx *gorm.DB) {
	mappingKeywordListObject.db = tx.Begin()
}

func (mappingKeywordListObject *MappingKeywordList) Rollback() {
	mappingKeywordListObject.db.Rollback()

		// after commit transaction, we have to rollback transaction
		mappingKeywordListObject.db = mappingKeywordListObject.db2
} 

func (mappingKeywordListObject *MappingKeywordList) MappingKeywordListModel() (tx *gorm.DB) {
	return mappingKeywordListObject.db.Model(&models.MappingKeywordList{})
}

// find all mapping keyword list paginated

func (mappingKeywordListObject *MappingKeywordList) FindMappingkeywordlist(pagination *models.Paginate, search, value string) ([]models.MappingKeywordList, *models.Paginate, error) {
	var mappingKeywordLists []models.MappingKeywordList
	data := mappingKeywordListObject.MappingKeywordListModel().
		Count(&pagination.Total)

		if search != "" {
			data.Where("lower(mappingKeywordLisit.code) like ?", "%"+strings.ToLower(search)+"%").Count(&pagination.Total)
		}

		if value != "" {
			data.Order("mappingKeywordList.id = " + value + " descc")
		}

		// search data
		data.Scopes(pagination.Pagination()).Debug().
			Find(&mappingKeywordLists)

		// checking errors
		if err := data.Error; err != nil {
			return []models.MappingKeywordList{}, pagination, err
		}

		return mappingKeywordLists, pagination, nil
}


// find by id
func (mappingKeywordListObject *MappingKeywordList) FindMappingkeywordlistById(id uint) (models.MappingKeywordList, error) {
	findId := models.MappingKeywordList{
		BaseModel: models.BaseModel{
			ID: id,
		},
	}
	if err := mappingKeywordListObject.MappingKeywordListModel().First(&findId, "id =?", id).Error; err != nil {
		return models.MappingKeywordList{}, err
	}

	return findId, nil
}

// create mapping keyword list
func (mappingKeywordList *MappingKeywordList) CreateMappingKeywordList(mapping_keyword_list *dto.CreateMappingkeywordlist) (models.MappingKeywordList, error) {
	mappingKeywordListModel := models.MappingKeywordList{
		MappingKeywordList: mapping_keyword_list.Keyword,
	}

	err := mappingKeywordList.db.Debug().Create(&mappingKeywordListModel).Error

	if err != nil {
		return mappingKeywordListModel, nil
	}

	return mappingKeywordListModel, nil
}

// update mapping keyword list
func (mappingKeywordListObject *MappingKeywordList) UpdateMappingKeywordList(id uint, mappingKeywordList *dto.UpdateMappingkeywordlist) error {
	update := mappingKeywordListObject.MappingKeywordListModel().Where("mappingKeywordList.id = ?", id).Updates(models.MappingKeywordList{
		MappingKeywordList: mappingKeywordList.Keyword,
	})

	if err := update.Error; err != nil {
		return err
	}

	return nil 
}

// delete mapping keyword list
func (mappingKeywordListObject *MappingKeywordList) DeleteMappingKeywordList(id uint) error {
	deleteMappingKeywordList := mappingKeywordListObject.MappingKeywordListModel().Where("mappingKeywordList.id = ?", id).Delete(&models.MappingKeywordList{})

	if err := deleteMappingKeywordList.Error; err != nil {
		return err
	}

	return nil
}