package controllers

import (
	"EJM/dto"
	"EJM/internal/server"
	"EJM/pkg/models"
	"EJM/pkg/repository"
	"EJM/pkg/services"
	"EJM/utils"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
	"github.com/nicksnyder/go-i18n/v2/i18n"
	"gorm.io/gorm"
)

type RoleController struct {
	server      *server.Server
	db          *gorm.DB
	roleService *services.RoleService
}

type IRoleService interface {
	CreateRole(role *dto.CreateRole) (models.Role, error)
	FindRoles(getRoles *dto.GetRoles) ([]models.Role, *models.Paginate, error)
	FindRoleById(roleId uint) (models.Role, error)
	DeleteRole(roleId uint) error
	UpdateRole(id uint,role *dto.UpdateRole, ) error
	GetAccessRole(roleId uint) ([]models.RolesMenus, error)
	SetAccessRole(accessRole *dto.SetAccessRole) error
	DeleteAccessRole(roleId uint, menuId uint) error
	UpdateAccessRole(accessRole *dto.SetAccessRole) error
}

type Role interface {
	IRoleService
}

func NewRoleController(srv *server.Server) *RoleController {
	db := srv.DB
	return &RoleController{
		server: srv,
		db:     db,
		roleService: services.NewRoleService(&services.RoleService{
			RoleRepository: repository.NewRoleRepository(srv.DB),
			// ActionRepository: repository.NewActionRepository(srv.DB),
			// CasbinRepository: repository.NewCasbinRepository(srv.DB),
			MenuRepository: repository.NewMenuRepository(srv.DB),
		}),
	}
}

// Update Role Berdasarkan ID
// @Summary API untuk Update Role Berdasarkan ID
// @Tags    Role
// @Accept  json
// @Produce json
// @Param   id    path     string         true "Role ID"
// @Param   roles body     dto.CreateRole true "Update Role"
// @Success 201   {object} utils.Response
// @Failure 400   {object} middlewares.ResponseError
// @Failure 401   {object} middlewares.ResponseError
// @Failure 404   {object} middlewares.ResponseError
// @Failure 500   {object} middlewares.ResponseError
// @Router  /roles/{id} [PATCH]
func (roleConttroller *RoleController) UpdateRole(c echo.Context) error {
	request := new(dto.UpdateRole)
	id, errConv := strconv.Atoi(c.Param("id"))
	if errConv != nil {
		return errConv
	}
	request.ID = uint(id)

	// midleware
	if err := c.Bind(request); err != nil {
		return err
	}
	if err := c.Validate(request); err != nil {
		return err
	}

	// userID := utils.User(c).ID

	// transaction
	// roleRepoDB := roleConttroller.roleService.RoleRepository
	// casbinRepoDB := roleConttroller.roleService.CasbinRepository
	// casbinRepoDB.Begin(roleConttroller.db)
	// roleRepoDB.Begin(roleConttroller.db)

	// services
	// var roleService IRoleService
	// roleService = roleConttroller.roleService

	errUpdate := roleConttroller.roleService.UpdateRole(uint(id), request)
	// errUpdate := roleService.UpdateRole(uint(id), request)
	if errUpdate != nil {
		// casbinRepoDB.Rollback()
		// roleRepoDB.Rollback()
		return errUpdate
	}

	// errAkses := roleService.SetAccessRole(&dto.SetAccessRole{
	// 	RoleId:  request.ID,
	// 	Actions: request.Actions,
	// })
	// if errAkses != nil {
	// 	roleRepoDB.Rollback()
	// 	casbinRepoDB.Rollback()
	// 	return errAkses
	// }

	// casbinRepoDB.Commit()
	// roleRepoDB.Commit()

	res := utils.Response{
		Data:       nil,
		Message:    "Berhasil Update data",
		StatusCode: 201,
	}

	return res.ReturnSingleMessage(c)
}

// Find Role Berdasarkan ID
// @Summary API untuk Get Role Berdasarkan ID
// @Tags    Role
// @Accept  json
// @Produce json
// @Param   id  path     string true "Role ID"
// @Success 201 {object} utils.Response
// @Failure 400 {object} middlewares.ResponseError
// @Failure 401 {object} middlewares.ResponseError
// @Failure 404 {object} middlewares.ResponseError
// @Failure 500 {object} middlewares.ResponseError
// @Router  /roles/{id} [GET]
func (roleController *RoleController) FindRoleById(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))

	if err != nil {
		return err
	}

	// services
	var roleServices IRoleService
	roleServices = roleController.roleService

	data, err := roleServices.FindRoleById(uint(id))

	if err != nil {
		// return utils.ErrRoleNotExists
	}

	res := utils.Response{
		Data:       data,
		Message:    "Berhasil Get Data Role",
		StatusCode: 200,
	}

	return res.ReturnSingleMessage(c)
}

// Delete Role Berdasarkan ID
// @Summary API untuk Delete Role Berdasarkan ID
// @Tags    Role
// @Accept  json
// @Produce json
// @Param   roleId path     string true "Role Id"
// @Success 201    {object} utils.Response
// @Failure 400    {object} middlewares.ResponseError
// @Failure 401    {object} middlewares.ResponseError
// @Failure 404    {object} middlewares.ResponseError
// @Failure 500    {object} middlewares.ResponseError
// @Router  /roles/delete/{roleId} [delete]
func (roleController *RoleController) DeleteRole(c echo.Context) error {
	data := c.Param("roleId")

	roleId, _ := strconv.Atoi(data)

	// transaction
	// roleRepoDB := roleController.roleService.RoleRepository
	// casbinRepoDB := roleController.roleService.CasbinRepository

	// casbinRepoDB.Begin(roleController.db)
	// roleRepoDB.Begin(roleController.db)

	// services
	var roleServices IRoleService
	roleServices = roleController.roleService

	errDeleteRole := roleServices.DeleteRole(uint(roleId))

	if errDeleteRole != nil {
		// casbinRepoDB.Rollback()
		// roleRepoDB.Rollback()
		// return errDeleteRole
	}

	// casbinRepoDB.Commit()
	// roleRepoDB.Commit()

	res := utils.Response{
		Translating: &i18n.Message{
			ID:    "role.delete",
			Other: "Berhasil Menghapus Role",
		},
		StatusCode: 200,
	}

	return res.ReturnSingleMessage(c)
}

// Daftar Role Baru
// @Summary API untuk Daftar Role Baru
// @Tags    Role
// @Accept  json
// @Produce json
// @Param   user body     dto.CreateRole true "Daftar Role Baru"
// @Success 201  {object} utils.Response{data=models.Role}
// @Failure 400  {object} middlewares.ResponseError
// @Failure 401  {object} middlewares.ResponseError
// @Failure 404  {object} middlewares.ResponseError
// @Failure 500  {object} middlewares.ResponseError
// @Router  /roles [post]
func (roleController *RoleController) CreateRole(c echo.Context) error {
	request := new(dto.CreateRole)

	// midleware
	if err := c.Bind(request); err != nil {
		return err
	}
	if err := c.Validate(request); err != nil {
		return err
	}

	// userLoginID := utils.User(c).ID
	// roleRepoDB := roleController.roleService.RoleRepository

	// service
	var roleService IRoleService
	roleService = roleController.roleService

	data, err := roleService.CreateRole(request)
	if err != nil {
		// roleRepoDB.Rollback()
		return err
	}

	// errAkses := roleService.SetAccessRole(&dto.SetAccessRole{
	// 	RoleId:  data.ID,
		// Actions: request.Actions,
	// })
	// if errAkses != nil {
		// roleRepoDB.Rollback()
	// 	return errAkses
	// }

	// roleRepoDB.Commit()

	res := utils.Response{
		Data:       data,
		Message:    "Berhasil Insert data",
		StatusCode: 201,
	}

	return res.ReturnSingleMessage(c)
}

// Ambil Daftar Role
// @Summary API untuk Ambil daftar Role
// @Tags    Role
// @Accept  json
// @Produce json
// @Param   page      query    string true  "Halaman"
// @Param   page_size query    string true  "Jumlah Data Per halaman"
// @Param   search    query    string false "Mencari Data"
// @Success 200       {object} utils.ResponsePaginate
// @Failure 400       {object} middlewares.ResponseError
// @Failure 401       {object} middlewares.ResponseError
// @Failure 404       {object} middlewares.ResponseError
// @Failure 500       {object} middlewares.ResponseError
// @Router  /roles [get]
func (roleController *RoleController) GetRoles(c echo.Context) error {
	dtos := new(dto.GetRoles)

	// midleware
	if err := c.Bind(dtos); err != nil {
		return err
	}
	if err := c.Validate(dtos); err != nil {
		return err
	}

	// service
	var (
		roleService IRoleService
	)

	roleService = roleController.roleService
	data, total, err := roleService.FindRoles(dtos)
	if err != nil {
		return err
	}

	res := &utils.ResponsePaginate{
		Key:  "roles",
		Meta: total,
		Response: utils.Response{
			Data:       data,
			Message:    "Sukses mengambil data role",
			StatusCode: 200,
		},
	}

	return res.ReturnPaginates(c)
}

// Delete Akses Role
// @Summary API Delete Akses Role
// @Tags    Role
// @Accept  json
// @Produce json
// @Param   roleId query    string true "role id"
// @Param   menuId query    string true "menu id"
// @Success 200    {object} utils.ResponsePaginate
// @Failure 400    {object} middlewares.ResponseError
// @Failure 401    {object} middlewares.ResponseError
// @Failure 404    {object} middlewares.ResponseError
// @Failure 500    {object} middlewares.ResponseError
// @Router  /roles/delete-access-role [DELETE]
func (roleController *RoleController) DeleteAccessRole(c echo.Context) error {
	idRole, errRoleId := strconv.Atoi(c.QueryParam("roleId"))
	menuId, errRoleMenuId := strconv.Atoi(c.QueryParam("menuId"))

	if errRoleId != nil {
		return errRoleId
	}

	if errRoleMenuId != nil {
		return errRoleMenuId
	}

	var roleMenuServices IRoleService
	roleMenuServices = roleController.roleService

	err := roleMenuServices.DeleteAccessRole(uint(idRole), uint(menuId))

	if err != nil {
		return err
	}

	err = roleController.server.Casbin.LoadPolicy()

	res := utils.Response{
		Message:    "Sukses Delete Access Role",
		StatusCode: http.StatusOK,
	}

	return res.ReturnSingleMessage(c)
}

// Set Access Role
// @Summary Set Akses Role
// @Tags    Auth
// @Accept  json
// @Produce json
// @Param   login body     dto.SetAccessRole true "Akses Role"
// @Success 201   {object} utils.Response
// @Failure 400   {object} middlewares.ResponseError
// @Failure 401   {object} middlewares.ResponseError
// @Failure 404   {object} middlewares.ResponseError
// @Failure 500   {object} middlewares.ResponseError
// @Router  /roles/set-access-role [post]
func (roleController *RoleController) SetAccessRole(c echo.Context) error {
	req := new(dto.SetAccessRole)
	if err := c.Bind(req); err != nil {
		return err
	}
	if err := c.Validate(req); err != nil {
		return err
	}

	roleController.server.DB = roleController.server.DB.Begin()

	err := roleController.roleService.SetAccessRole(req)
	if err != nil {
		roleController.server.DB.Rollback()
		return err
	}

	err = roleController.server.Casbin.LoadPolicy()

	if err != nil {
		return err
	}

	roleController.server.DB.Commit()

	res := utils.Response{
		Data:       nil,
		Message:    "Sukses Set Access Role",
		StatusCode: http.StatusOK,
	}

	return res.ReturnSingleMessage(c)
}

// Update Access Role
// @Summary Update Akses Role
// @Tags    Auth
// @Accept  json
// @Produce json
// @Param   login body     dto.SetAccessRole true "Akses Role"
// @Success 201   {object} utils.Response
// @Failure 400   {object} middlewares.ResponseError
// @Failure 401   {object} middlewares.ResponseError
// @Failure 404   {object} middlewares.ResponseError
// @Failure 500   {object} middlewares.ResponseError
// @Router  /roles/update-access-role [PATCH]
func (roleController *RoleController) UpdateAccessRole(c echo.Context) error {
	req := new(dto.SetAccessRole)
	if err := c.Bind(req); err != nil {
		return err
	}
	if err := c.Validate(req); err != nil {
		return err
	}

	roleController.server.DB = roleController.server.DB.Begin()

	err := roleController.roleService.SetAccessRole(req)
	if err != nil {
		roleController.server.DB.Rollback()
		return err
	}

	err = roleController.server.Casbin.LoadPolicy()

	if err != nil {
		return err
	}

	roleController.server.DB.Commit()

	res := utils.Response{
		Data:       nil,
		Message:    "Sukses Update Access Role",
		StatusCode: http.StatusOK,
	}

	return res.ReturnSingleMessage(c)
}

// Ambil Daftar Role Menu
// @Summary API untuk Ambil daftar Role
// @Tags    Role
// @Accept  json
// @Produce json
// @Param   page      query    string true  "Halaman"
// @Param   page_size query    string true  "Jumlah Data Per halaman"
// @Param   search    query    string false "Mencari Data"
// @Success 200       {object} utils.ResponsePaginate
// @Failure 400       {object} middlewares.ResponseError
// @Failure 401       {object} middlewares.ResponseError
// @Failure 404       {object} middlewares.ResponseError
// @Failure 500       {object} middlewares.ResponseError
// @Router  /accessroles [get]
func (roleController *RoleController) GetAccessRole(c echo.Context) error {
	id, err := strconv.Atoi(c.QueryParam("roleId"))

	if err != nil {
		return err
	}

	var roleMenuService IRoleService
	roleMenuService = roleController.roleService

	data, err := roleMenuService.GetAccessRole(uint(id))

	if err != nil {
		return err
	}

	res := utils.Response{
		Data:       data,
		Message:    "Berhasil Get Role Menu",
		StatusCode: 200,
	}

	return res.ReturnSingleMessage(c)
}

// Hapus Role Banyak
// @Summary API untuk Hapus Role banyak
// @Tags    Role
// @Accept  json
// @Produce json
// @Param   roleIds body     dto.DeleteRoleBulk true "Hapus Role banyak"
// @Success 200     {object} utils.Response
// @Failure 400     {object} middlewares.ResponseError
// @Failure 401     {object} middlewares.ResponseError
// @Failure 404     {object} middlewares.ResponseError
// @Failure 500     {object} middlewares.ResponseError
// @Router  /roles/delete-bulk [delete]
func (roleConttroller *RoleController) DeleteRoleBulk(c echo.Context) error {
	dto := new(dto.DeleteRoleBulk)
	if err := c.Bind(dto); err != nil {
		return err
	}
	if err := c.Validate(dto); err != nil {
		return err
	}

	// casbinRepoDB := roleConttroller.roleService.CasbinRepository
	// roleRepoDB := roleConttroller.roleService.RoleRepository
	// casbinRepoDB.Begin(roleConttroller.db)
	// roleRepoDB.Begin(roleConttroller.db)

	var roleServices IRoleService
	roleServices = roleConttroller.roleService

	for _, id := range dto.RoleIds {
		errDeleteRole := roleServices.DeleteRole(id)

		if errDeleteRole != nil {
			// casbinRepoDB.Rollback()
			// roleRepoDB.Rollback()
			return errDeleteRole
		}
	}

	// casbinRepoDB.Commit()
	// roleRepoDB.Commit()

	responses := utils.Response{
		Data: nil,
		Translating: &i18n.Message{
			ID:    "role.delete",
			Other: "Sukses Delete Role",
		},
		StatusCode: 201,
	}

	return responses.ReturnSingleMessage(c)
}
