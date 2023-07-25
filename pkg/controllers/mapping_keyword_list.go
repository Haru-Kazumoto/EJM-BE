package controllers

import (
	"EJM/dto"
	"EJM/internal/server"
	"EJM/pkg/repository"
	"EJM/pkg/services"
	"EJM/utils"
	"errors"
	"net/http"
	"strconv"

	// "github.com/nicksnyder/go-i18n/v2/i18n"

	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

type MappingKeyowrdListController struct {
	db *gorm.DB
	server server.IServerInterface
	mappingKeywordListService *services.MappingKeywordListService
}

func NewMappingKeywordListcontroller(srv *server.Server) *MappingKeyowrdListController {
	db := srv.DB
	return &MappingKeyowrdListController{
		server: srv,
		db: db,
		mappingKeywordListService: services.NewMappingKeywordListService(&services.MappingKeywordListService{
			MappingKeywordListRepository: repository.NewMappingKeywordListRepository(srv.DB),
		})}
}

// Daftar Mapping Keyword List Baru
// @Summary API untuk Membuat Mapping Keyword List Baru
// @Tags    Keyword List
// @Accept  json
// @Produce json
// @Param   mappingkeywordlist body     dto.CreateMappingkeywordlist true "Daftar Mapping Keyword List Baru"
// @Success 201  {object} utils.Response
// @Failure 400  {object} middlewares.ResponseError
// @Failure 401  {object} middlewares.ResponseError
// @Failure 404  {object} middlewares.ResponseError
// @Failure 500  {object} middlewares.ResponseError
// @Router  /mappingKeywordList/create [post]
func (mappingKeyowrdListController *MappingKeyowrdListController) CreateMappingKeywordList(c echo.Context) error {
	req := new(dto.CreateMappingkeywordlist)

	if err := c.Bind(req); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, "Gagal melakukan Binding keyword")
	}
	if err := c.Validate(req); err != nil{
		return echo.NewHTTPError(http.StatusBadRequest, "Validasi gagal")
	}

	data, err := mappingKeyowrdListController.mappingKeywordListService.CreateMappingKeywordList(req)
	if err != nil {
		if errors.Is(err, utils.ErrMappingKeywordListAlreadyExists) {
			res := utils.Response{
				Data:       nil,
				Message:    "Mapping Keyword List Already Exist",
				StatusCode: 404,
			}
			return res.ReturnSingleMessage(c)
		}

	return err
	}

	res := utils.Response{
		Data: data,
		Message: "Berhasil membuat keyword",
		StatusCode: 201,
	}

	return res.ReturnSingleMessage(c)
}

// Ambil Daftar Mapping Mapping keyword list
// @Summary API untuk Ambil semua daftar mapping keyword list
// @Tags    Keyword List
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
// @Router  /mappingKeywordList [get]
func (mappingKeyowrdListController *MappingKeyowrdListController) FindMappingkeywordlist(c echo.Context) error {
	req := new(dto.GetMappingkeywordlist)

	if err := c.Bind(req); err != nil {
		return err
	}
	if err := c.Validate(req); err != nil {
		return err
	}

	data, meta, err := mappingKeyowrdListController.mappingKeywordListService.FindMappingkeywordlist(req)
	if err != nil {
		return err
	}

	res := utils.ResponsePaginate{
		Key: "mapping_keyword_list",
		Meta: meta,
		Response: utils.Response{
			Data: data,
			Message: "Sukses mengambil data mapping keyword list",
			StatusCode: 200,
		},
	}

	return res.ReturnPaginates(c)
}


// Update Keyword List
// @Summary API untuk Update Data Keyword List
// @Tags    Keyword Lists
// @Accept  json
// @Produce json
// @Param   id   path     int               true "ID"
// @Param   mappingkeywordlist body     dto.UpdateMappingkeywordlist true "Update Keyword List Baru"
// @Success 201  {object} utils.Response
// @Failure 400  {object} middlewares.ResponseError
// @Failure 401  {object} middlewares.ResponseError
// @Failure 404  {object} middlewares.ResponseError
// @Failure 500  {object} middlewares.ResponseError
// @Router  /mappingKeywordList/{id} [put]

func (mappingKeyowrdListController *MappingKeyowrdListController) UpdateMappingkeywordlist(c echo.Context) error  {
	req := new(dto.UpdateMappingkeywordlist)
	id, errConvert := strconv.Atoi(c.Param("id"))
	if errConvert != nil {
		return errConvert
	}	 
	req.ID = uint(id)

	if err := c.Bind(req); err != nil{
		return err
	}

	if err := c.Validate(req); err != nil {
		return err
	}

	// check id available in database?
	_, errFind := mappingKeyowrdListController.mappingKeywordListService.FindMappingkeywordlistById(uint(id))
	if errFind != nil {
		if errors.Is(errFind, utils.ErrMappingKeywordListAlreadyExists) {
			res := utils.Response{
				Data: nil,
				Message: "Data not found",
				StatusCode: 404,
			}

			return res.ReturnSingleMessage(c)
		}
		return errFind
	}

	err := mappingKeyowrdListController.mappingKeywordListService.UpdateMappingkeywordlist(uint(id), req)
	if err != nil {
		return err
	}

	res := utils.Response{
		Data: nil,
		Message: "Berhasil update keyword",
		StatusCode: 201,
	}

	return res.ReturnSingleMessage(c)
}

// Delete Mapping Keyword List
// @Summary API untuk Delete Mapping Keywprd List
// @Tags    Keyword List
// @Accept  json
// @Produce json
// @Param   id  path     int true "Keyword List ID"
// @Success 201 {object} utils.Response
// @Failure 400 {object} middlewares.ResponseError
// @Failure 401 {object} middlewares.ResponseError
// @Failure 404 {object} middlewares.ResponseError
// @Failure 500 {object} middlewares.ResponseError
// @Router  /mappingKeywordList/{id} [delete]
func (mappingKeywordListController *MappingKeyowrdListController) DeleteMappingkeywordlist(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))

	err = mappingKeywordListController.mappingKeywordListService.DeleteMappingkeywordlist(uint(id))
	if err != nil {
		if errors.Is(err, utils.ErrMappingKeywordListNotFound) {
			res := utils.Response{
				Data: nil,
				Message: "Data not found",
				StatusCode: 404,
			}

			return res.ReturnSingleMessage(c)
		}
		return err
	}

	res := utils.Response{
		Data:       nil,
		Message:    "Berhasil Hapus Data",
		StatusCode: 201,
	}

	return res.ReturnSingleMessage(c)
}


// Find Keyword List By ID
// @Summary API untuk Keyword List By ID
// @Tags    Keyword List
// @Accept  json
// @Produce json
// @Param   id  path     int true "ID"
// @Success 201 {object} utils.Response
// @Failure 400 {object} middlewares.ResponseError
// @Failure 401 {object} middlewares.ResponseError
// @Failure 404 {object} middlewares.ResponseError
// @Failure 500 {object} middlewares.ResponseError
// @Router  /mappingKeywordList/{id} [GET]
func (mappingKeywordListController *MappingKeyowrdListController) FindMappingkeywordlistById(c echo.Context) error  {
	id, err := strconv.Atoi(c.Param("id"))

	if err != nil {
		return err
	}

	data, err := mappingKeywordListController.mappingKeywordListService.FindMappingkeywordlistById(uint(id))
	if err != nil {
		if errors.Is(err, utils.ErrMappingKeywordListNotFound) {
			res := utils.Response{
				Data:       nil,
				Message:    "Data Not Found",
				StatusCode: 404,
			}
		
			return res.ReturnSingleMessage(c)
		}
		return err
	}

	res := utils.Response{
		Data:       data,
		Message:    "Berhasil Get Data Mapping Keyword List",
		StatusCode: 200,
	}

	return res.ReturnSingleMessage(c)
}