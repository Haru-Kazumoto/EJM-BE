package dto

import "EJM/pkg/models"

type BasePagination struct {
	Page        string `json:"page" query:"page" validate:"required"`
	PageSize    string `json:"page_size" query:"page_size" validate:"required"`
	Search      string `json:"search" query:"search"`
	UsingActive bool   `json:"using_active" query:"using_active" form:"using_active"`
	Value       string `json:"value" query:"value"`
}

type BasePaginationNew struct {
	Page        string            `json:"page" query:"page" validate:"required"`
	PageSize    string            `json:"page_size" query:"page_size" validate:"required"`
	Search      string            `json:"search" query:"search"`
	Active models.ActiveEnum `json:"active" query:"active" form:"active"`
	Value       string            `json:"value" query:"value"`
}

// --------------------User
type CreateNewUser struct {
	Name     string `json:"name" validate:"required"`
	Username string `json:"username" validate:"required"`
	Password string `json:"password" validate:"required"`
	RoleId   uint   `json:"roleId" validate:"required,min=1"`
	Active     ActiveEnum `json:"active" form:"active" validate:"required"`
}

type LoginByPin struct {
	Pin string `json:"pin" validate:"required"`
}

type DeleteUserBulk struct {
	Ids []uint `json:"ids" validate:"required,min=1"`
}

type UpdateUser struct {
	ID uint `params:"id" validate:"required"`
	CreateNewUser
}

type ChangePassword struct {
	NewPassword string `json:"newPassword" validate:"required"`
}

type ToggleActive struct {
	ID      uint  `param:"id" validate:"required" swaggerignore:"true"`
	Payload *bool `json:"payload"`
}

type GetUsers struct {
	BasePagination
	Users []CreateNewUserResponse `json:"users"`
}

type CreateNewUserResponse struct {
	ID              uint   `json:"id"`
	Name       string     `json:"name"`
	Username   string     `json:"username"`
	RoleId     uint       `json:"roleId"`
	Active     ActiveEnum `json:"active"`
	RoleName   string     `json:"roleName"` // Field untuk menampilkan role name pada response
}

type GetNotification struct {
	BasePagination
	UserId uint `json:"user_id" form:"user_id" validate:"required"`
	Type   uint `json:"type" validate:"required"`
}

type PayloadCreateMenu struct {
	Data []CreateMenu `json:"data" validate:"required,min=1"`
}

type PayloadCreateAction struct {
	Data []CreateAction `json:"data" validate:"required,min=1"`
}

// mapping keyword list
type CreateMappingkeywordlist struct {
	Keyword string `json:"MappingKeywordList" form:"Name" validate:"required"`
}

type GetMappingkeywordlist struct {
	BasePagination
}

type UpdateMappingkeywordlist struct {
	ID uint `params:"id" validate:"required"`
	CreateMappingkeywordlist
}

type DeleteMappingkeywordlist struct {
	Ids []uint `json:"ids" validate:"required,min=1"`
}

// mapping code
type CreateNewMappingCode struct {
	Code       string     `json:"code" form:"code" validate:"required"`
	Definition string     `json:"definition" form:"definition" validate:"required"`
	Status     StatusEnum `json:"status" validate:"required"`
	Priority   int        `json:"priority" form:"priority" validate:"required"`
	Active     ActiveEnum `json:"active" form:"active" validate:"required"`
}
type StatusEnum string

const (
	Success StatusEnum = "success"
	Fail    StatusEnum = "fail"
)

type GetMappingCodes struct {
	BasePagination
}

type ActiveEnum string

const (
	Active   ActiveEnum = "active"
	InActive ActiveEnum = "inActive"
)

type DeleteMappingCodeBulk struct {
	Ids []uint `json:"ids" validate:"required,min=1"`
}

type UpdateMappingCode struct {
	ID uint `params:"id" validate:"required"`
	CreateNewMappingCode
}

// menu
type CreateMenu struct {
	Name     string `json:"name" validate:"required"`
	Path     string `json:"path" validate:"required"`
	ParentId *uint  `json:"parent_id"`
}

type CreateAction struct {
	Name    string `json:"name" validate:"required"`
	Path    string `json:"path" validate:"required"`
	Method  string `json:"method" validate:"required"`
	Default *bool  `json:"default"`
}

type GetMenu struct {
	Name     string `json:"name" validate:"required"`
	ParentId int    `json:"parent_id" validate:"required"`
	Path     string `json:"path" validate:"required"`
}

type DeleteMenu struct {
	Id uint `param:"id" validate:"required" swaggerignore:"true"`
}

type DeleteAction struct {
	GetDetailAction
}

type GetDetailMenu struct {
	DeleteMenu
}

type GetDetailAction struct {
	Id uint `param:"id" validate:"required" swaggerignore:"true"`
}

type GetActionByMenu struct {
	MenuId uint `param:"menu_id"`
}

type DeleteMenuBulk struct {
	Ids []uint `json:"ids" validate:"required,min=1"`
}

type DeleteActionBulk struct {
	DeleteMenuBulk
}

type UpdateMenu struct {
	GetDetailMenu
	APIS []uint `json:"apis"`
	CreateMenu
}

type UpdateAction struct {
	GetDetailAction
	CreateAction
}

// --------------------Roles
type CreateRole struct {
	Name          string        `json:"name" validate:"required"`
	Description string        `validate:"required" json:"description"` //Main field
	// Actions       []MenusAction `json:"actions" validate:"required"`
	// ObjectActions []struct {
	// 	ID       int    `json:"id" validate:"required"`
	// 	Name     string `json:"name" validate:"required"`
	// 	ParentID uint   `json:"parentId"`
	// 	Path     string `json:"path" validate:"required"`
	// 	State    int    `json:"state"`
	// 	Type     string `json:"type" validate:"required"`
	// } `json:"object_actions" validate:"required,dive,required"`
}

type DeleteRoleBulk struct {
	RoleIds []uint `json:"role_ids" validate:"required,min=1"`
}

type UpdateRole struct {
	ID uint `params:"id" form:"id" validate:"required"`
	CreateRole
}

type CreateRoleMenu struct {
	ID    uint   `param:"id" form:"id" validate:"required"`
	Menus []uint `json:"list_menu" validate:"required"`
}

type UpdateRoleMenu struct {
	ID    uint   `param:"id" form:"id" validate:"required"`
	Menus []uint `json:"list_menu" validate:"required"`
}

type GetRoles struct {
	BasePagination
}

// --------------------Auth
type RefreshToken struct {
	Token string `json:"refresh_token" validate:"required"`
}

type MenusAction struct {
	MenuId uint `json:"menuId" validate:"required"`
	APIId  []uint
}

type SetAccessRole struct {
	RoleId  uint          `json:"roleId" validate:"required"`
	Actions []MenusAction `json:"actions" validate:"required"`
}

type Login struct {
	Username string `json:"username" validate:"required"`
	Password string `json:"password" validate:"required"`
}

type ObjectStructures struct {
	Name     string  `json:"name" validate:"required"`
	Code     *string `json:"structureCode"`
	ParentID string  `json:"parentId"`
	ID       string  `json:"id" validate:"required"`
}

//Jenis Transaksi
type CreateNewJenisTransaksi struct {
	TransactionType string `json:"transactionType" form:"transactionType" validate:"required"`
	TransactionGroup string `json:"transactionGroup" form:"transactionGroup" validate:"required"`
}

type UpdateJenisTransaksi struct {
	ID uint `params:"id" validate:"required"`
	CreateNewJenisTransaksi
}

type GetJenisTransaksi struct {
	BasePagination
}

// list op code
type CreateListOpCode struct {
	OPCode string `json:"opCode" form:"opCode" validate:"required"`
	ModelMesin string `json:"modelMesin" form:"modelMesin" validate:"required"`
	TipeTransaksiID uint `json:"tipeTransaksiid" form:"tipeTransaksiid" validate:"required"`
}

type GetListOpCode struct {
	BasePagination
	ListOpCodes []GetListOpCodeResponse `json:"listOpCodes"`

}
type GetListOpCodeResponse struct {
	ID              uint   `json:"id"`
	OPCode           string `json:"opCode"`
	ModelMesin       string `json:"modelMesin"`
	TipeTransaksiID  uint   `json:"tipeTransaksiid"`
	TransactionType  string `json:"transactionType"`
	TransactionGroup string `json:"transactionGroup"`
}

type DeleteListOpCode struct {
	ID uint `params:"id" validate:"required"`
}

type UpdateListOpCode struct {
	ID uint `params:"id" validate:"required"`
	CreateListOpCode
}

type CreateNewBinKartu struct {
	PrefixNo uint `json:"prefix_no" form:"prefix_no" gorm:"uniqueIndex" validate:"required"`
	BankName string `json:"bank_name" form:"bank_name" validate:"required"`
}

type UpdateBinKartu struct{
	ID uint `params:"id" validate:"required"`
	CreateNewBinKartu
}

type GetBinKartu struct {
	BasePagination
}