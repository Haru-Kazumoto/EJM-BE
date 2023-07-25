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
	"gorm.io/gorm"

	"github.com/labstack/echo/v4"
)

type MappingCodeController struct {
	db                 *gorm.DB
	server             server.IServerInterface
	mappingCodeService *services.MappingCodeService
}

func NewMappingCodeController(srv *server.Server) *MappingCodeController {
	db := srv.DB
	return &MappingCodeController{
		server: srv,
		db:     db,
		mappingCodeService: services.NewMappingCodeService(&services.MappingCodeService{
			MappingCodeRepository: repository.NewMappingCodeRepository(srv.DB),
		})}
}

// Daftar Mapping Code Baru
// @Summary API untuk Membuat Mapping Code Baru
// @Tags    Mapping Codes
// @Accept  json
// @Produce json
// @Param   mappingcode body     dto.CreateNewMappingCode true "Daftar Mapping Code Baru"
// @Success 201  {object} utils.Response
// @Failure 400  {object} middlewares.ResponseError
// @Failure 401  {object} middlewares.ResponseError
// @Failure 404  {object} middlewares.ResponseError
// @Failure 500  {object} middlewares.ResponseError
// @Router  /mappingCodes/create [post]
func (mappingCodeController *MappingCodeController) CreateMappingCode(c echo.Context) error {
	req := new(dto.CreateNewMappingCode)
	if err := c.Bind(req); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, "Gagal melakukan binding data")
	}

	if err := c.Validate(req); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, "Validasi gagal")
	}

	data, err := mappingCodeController.mappingCodeService.CreateMappingCode(req)
	if err != nil {
		if errors.Is(err, utils.ErrCodeAlreadyExists) {
			res := utils.Response{
				Data:       nil,
				Message:    "Mapping Code Already Exist",
				StatusCode: 404,
			}
			return res.ReturnSingleMessage(c)
		}

		return err
	}

	res := utils.Response{
		Data:       data,
		Message:    "Berhasil Insert data",
		StatusCode: 201,
	}

	return res.ReturnSingleMessage(c)
}

// Ambil Daftar Mapping Code
// @Summary API untuk Ambil semua daftar mapping code
// @Tags    Mapping Codes
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
// @Router  /mappingCodes [get]
func (mappingCodeController *MappingCodeController) FindMappingCodes(c echo.Context) error {
	req := new(dto.GetMappingCodes)

	if err := c.Bind(req); err != nil {
		return err
	}
	if err := c.Validate(req); err != nil {
		return err
	}

	data, meta, err := mappingCodeController.mappingCodeService.FindMappingCodes(req)
	if err != nil {
		return err
	}

	res := utils.ResponsePaginate{
		Key:  "mapping_codes",
		Meta: meta,
		Response: utils.Response{
			Data:       data,
			Message:    "Sukses mengambil data mapping codes",
			StatusCode: 200,
		},
	}

	return res.ReturnPaginates(c)

}

// Update Mapping Code
// @Summary API untuk Update Data Mapping Code
// @Tags    Mapping Codes
// @Accept  json
// @Produce json
// @Param   id   path     int               true "Code ID"
// @Param   mappingCode body     dto.CreateNewMappingCode true "Update Mapping Code Baru"
// @Success 201  {object} utils.Response
// @Failure 400  {object} middlewares.ResponseError
// @Failure 401  {object} middlewares.ResponseError
// @Failure 404  {object} middlewares.ResponseError
// @Failure 500  {object} middlewares.ResponseError
// @Router  /mappingCodes/{id} [put]
func (mappingCodeController *MappingCodeController) UpdateMappingCode(c echo.Context) error {
	req := new(dto.UpdateMappingCode)
	id, errConvert := strconv.Atoi(c.Param("id"))
	if errConvert != nil {
		return errConvert
	}
	req.ID = uint(id)

	if err := c.Bind(req); err != nil {
		return err
	}
	if err := c.Validate(req); err != nil {
		return err
	}

	// Cek apakah ID tersedia dalam database
	_, errFind := mappingCodeController.mappingCodeService.FindMappingCodeById(uint(id))
	if errFind != nil {
		if errors.Is(errFind, utils.ErrMappingCodeNotFound) {
			res := utils.Response{
				Data:       nil,
				Message:    "Data Not Found",
				StatusCode: 404,
			}
		
			return res.ReturnSingleMessage(c)
		}
		return errFind
	}
	
	err := mappingCodeController.mappingCodeService.UpdateMappingCode(uint(id), req)
	if err != nil {
		return err
	}

	res := utils.Response{
		Data:       nil,
		Message:    "Berhasil Update data",
		StatusCode: 201,
	}

	return res.ReturnSingleMessage(c)
}

// Delete Mapping Code
// @Summary API untuk Delete Mapping Code
// @Tags    Mapping Codes
// @Accept  json
// @Produce json
// @Param   id  path     int true "Code ID"
// @Success 201 {object} utils.Response
// @Failure 400 {object} middlewares.ResponseError
// @Failure 401 {object} middlewares.ResponseError
// @Failure 404 {object} middlewares.ResponseError
// @Failure 500 {object} middlewares.ResponseError
// @Router  /mappingCodes/{id} [delete]
func (mappingCodeController *MappingCodeController) DeleteMappingCode(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))

	// mappingCodeRepo := mappingCodeController.mappingCodeService.RegisterRepository

	// mappingCodeRepo.Begin(mappingCodeController.db)

	err = mappingCodeController.mappingCodeService.DeleteMappingCode(uint(id))
	if err != nil {
		if errors.Is(err, utils.ErrMappingCodeNotFound) {
			res := utils.Response{
				Data:       nil,
				Message:    "Data Not Found",
				StatusCode: 404,
			}
		
			return res.ReturnSingleMessage(c)
		}
		return err
	}

	// mappingCodeRepo.Commit()

	res := utils.Response{
		Data:       nil,
		Message:    "Berhasil Hapus Data",
		StatusCode: 201,
	}

	return res.ReturnSingleMessage(c)
}

// Delete Mapping Code Bulk
// @Summary API untuk Delete Mapping Code Bulk
// @Tags    Mapping Codes
// @Accept  json
// @Produce json
// @Param   mapping_code_ids body     dto.DeleteMappingCodeBulk true "Delete Mapping Code Bulk"
// @Success 201      {object} utils.Response
// @Failure 400      {object} middlewares.ResponseError
// @Failure 401      {object} middlewares.ResponseError
// @Failure 404      {object} middlewares.ResponseError
// @Failure 500      {object} middlewares.ResponseError
// @Router  /mappingCodes [delete]
func (mappingCodeController *MappingCodeController) DeleteMappingCodeBulk(c echo.Context) error {
	req := new(dto.DeleteMappingCodeBulk)
	if err := c.Bind(req); err != nil {
		return err
	}
	if err := c.Validate(req); err != nil {
		return err
	}

	// mappingCodeRepo := mappingCodeController.mappingCodeService.RegisterRepository

	// mappingCodeRepo.Begin(mappingCodeController.db)

	for _, id := range req.Ids {
		err := mappingCodeController.mappingCodeService.DeleteMappingCode(uint(id))
		if err != nil {
			if errors.Is(err, utils.ErrMappingCodeNotFound) {
				res := utils.Response{
					Data:       nil,
					Message:    "Data Not Found",
					StatusCode: 404,
				}
			
				return res.ReturnSingleMessage(c)
			}
				return err
		}
	}

	// mappingCodeRepo.Commit()

	res := utils.Response{
		Data:       nil,
		Message:    "Berhasil Hapus Data",
		StatusCode: 201,
	}

	return res.ReturnSingleMessage(c)
}

// Find Mapping Code By ID
// @Summary API untuk Mapping Code By ID
// @Tags    Mapping Codes
// @Accept  json
// @Produce json
// @Param   id  path     int true "Code ID"
// @Success 201 {object} utils.Response
// @Failure 400 {object} middlewares.ResponseError
// @Failure 401 {object} middlewares.ResponseError
// @Failure 404 {object} middlewares.ResponseError
// @Failure 500 {object} middlewares.ResponseError
// @Router  /mappingCodes/{id} [GET]

func (mappingCodeController *MappingCodeController) FindMappingCodeById(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))

	if err != nil {
		return err
	}

	data, err := mappingCodeController.mappingCodeService.FindMappingCodeById(uint(id))
	if err != nil {
		if errors.Is(err, utils.ErrMappingCodeNotFound) {
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
		Message:    "Berhasil Get Data Mapping Code",
		StatusCode: 200,
	}

	return res.ReturnSingleMessage(c)
}
