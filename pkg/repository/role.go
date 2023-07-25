package repository

import (
	"EJM/dto"
	"EJM/pkg/models"
	"EJM/utils"

	"errors"
	"strings"

	"gorm.io/gorm"
)

type RoleRepository interface {
	TransactionRepository
	CreateRole(role *dto.CreateRole) (models.Role, error)
	FindRoleById(id uint) (models.Role, error)
	FindRoles(pagination *models.Paginate, search string, usingActive bool, value string) ([]models.Role, *models.Paginate, error)
	FindRoleByName(name string) error
	DeleteRole(roleId uint) error
	UpdateRole(id uint,role *dto.UpdateRole) error
	SetAccessRole(roleId uint, menus []uint) error
	UpdateAccessRole(roleId uint, menus []uint) error
	DeleteAccessRole(roleId uint, menuId uint) error
	DeleteMenusByRole(roleId uint) error
	GetAccessRole(roleId uint) ([]models.RolesMenus, error)
	FindByName(name string) (models.Role, error)
}

type Role struct {
	db  *gorm.DB
	db2 *gorm.DB
}

func NewRoleRepository(db *gorm.DB) *Role {
	return &Role{db: db, db2: db}
}

func (roleObject *Role) Begin(tx *gorm.DB) {
	roleObject.db = tx.Begin()
}

func (roleObject *Role) Rollback() {
	roleObject.db.Rollback()

	// after commit transaction, we have to rollback transaction
	roleObject.db = roleObject.db2
}

func (roleObject *Role) Commit() {
	roleObject.db.Commit()

	// after commit transaction, we have to rollback transaction
	roleObject.db = roleObject.db2
}

func (roleObject *Role) RoleModel() (tx *gorm.DB) {
	return roleObject.db.Model(&models.Role{})
}

func (roleObject *Role) RolesMenusModel() (tx *gorm.DB) {
	return roleObject.db.Model(&models.RolesMenus{})
}

// create role
func (roleObject *Role) CreateRole(role *dto.CreateRole) (models.Role, error) {
	// var objectMenus []interface{}

	// for _, action := range role.ObjectActions {
	// 	objectMenus = append(objectMenus, action)
	// }

	roleModel := models.Role{
		Name: role.Name,
		Description: role.Description,
		
	}
	err := roleObject.db.Create(&roleModel).Error

	if err != nil {
		return models.Role{}, err
	}
	return roleModel, nil

}

// Find Role By Name
func (roleObject *Role) FindRoleByName(name string) error {

	checkrole := models.Role{
		Name: name,
	}

	if err := roleObject.RoleModel().First(&checkrole, "name = ?", name).Error; err == nil {
		return utils.ErrRoleAlreadyExists
	}

	return nil
}

// Find By Name
func (roleObject *Role) FindByName(name string) (models.Role, error) {
	findRole := models.Role{
		Name: name,
	}
	if err := roleObject.RoleModel().First(&findRole, "roles.name LIKE ?", "%"+name+"%").Error; err != nil {
		return models.Role{}, err
	}

	return findRole, nil
}

// update role
func (roleObject *Role) UpdateRole(id uint,role *dto.UpdateRole) error {
	// var objectMenus []interface{}
	// for _, action := range role.ObjectActions {
	// 	objectMenus = append(objectMenus, action)
	// }

	update := roleObject.RoleModel().Where("roles.id = ?", id).Updates(models.Role{
		Name: role.Name,
		Description: role.Description,
	})

	if err := update.Error; err != nil {
		return err
	}

	return nil
}

// find role
func (roleObject *Role) FindRoleById(id uint) (models.Role, error) {
	checkRole := models.Role{
		BaseModel: models.BaseModel{
			ID: id,
		},
	}
	if err := roleObject.RoleModel().First(&checkRole, "id = ?", id).Error; err != nil {
		return models.Role{}, err
	}

	return checkRole, nil
}

// delete role
func (roleObject *Role) DeleteRole(roleId uint) error {
	deleteUser := roleObject.RoleModel().Where("roles.id = ?", roleId).Delete(&models.Role{})

	if err := deleteUser.Error; err != nil {
		return err
	}

	return nil
}

// find all role
func (roleObject *Role) FindRoles(pagination *models.Paginate, search string, usingActive bool, value string) ([]models.Role, *models.Paginate, error) {
	var roles []models.Role
	data := roleObject.RoleModel().Count(&pagination.Total)

	if search != "" {
		data.Where("lower(roles.name) like ? OR  lower(roles.description) like ?", "%"+strings.ToLower(search)+"%", "%"+strings.ToLower(search)+"%").Count(&pagination.Total)
	}

	if usingActive {
		data.Where("roles.is_active", true).Count(&pagination.Total)
	}

	// if value != "" {
	// 	data.Order("roles.id = " + value + " desc")
	// }

	// cari data
	data.Scopes(pagination.Pagination()).Debug().Find(&roles)

	// checking errors
	if err := data.Error; err != nil {
		return []models.Role{}, pagination, err
	}

	return roles, pagination, nil
}

// set role menu
func (roleObject *Role) SetAccessRole(roleId uint, menus []uint) error {
	if len(menus) < 1 {
		return errors.New("Menus tidak boleh kosong")
	}

	for _, menu := range menus {
		roleObject.RolesMenusModel().Create(&models.RolesMenus{
			MenuId: menu,
			RoleId: roleId,
		})
	}

	return nil
}

// update role menu
func (roleObject *Role) UpdateAccessRole(roleId uint, menus []uint) error {
	if len(menus) < 1 {
		return errors.New("Menus tidak boleh kosong")
	}

	deleteMenu := roleObject.RolesMenusModel().Where("role_id = ?", roleId).Delete(models.RolesMenus{})

	if err := deleteMenu.Error; err != nil {
		return err
	}

	for _, menu := range menus {
		roleObject.RolesMenusModel().Create(&models.RolesMenus{
			MenuId: menu,
			RoleId: roleId,
		})
	}

	return nil
}

// delete role menu
func (roleObject *Role) DeleteAccessRole(roleId uint, menuId uint) error {
	deleteRole := roleObject.RolesMenusModel().Where("role_id = ? AND menu_id = ?", roleId, menuId).Delete(&models.RolesMenus{})

	if err := deleteRole.Error; err != nil {
		return err
	}
	return nil
}

// delete role menu by role
func (roleObject *Role) DeleteMenusByRole(roleId uint) error {
	deleteRole := roleObject.RolesMenusModel().Where("role_id = ?", roleId).Delete(&models.RolesMenus{})

	if err := deleteRole.Error; err != nil {
		return err
	}
	return nil
}

// get role menu
func (roleObject *Role) GetAccessRole(roleId uint) ([]models.RolesMenus, error) {
	var roleMenus []models.RolesMenus

	roleMenuData := roleObject.RolesMenusModel().Where("roles_menus.role_id = ?", roleId).Find(&roleMenus)

	if err := roleMenuData.Error; err != nil {
		return []models.RolesMenus{}, err
	}

	return roleMenus, nil
}
