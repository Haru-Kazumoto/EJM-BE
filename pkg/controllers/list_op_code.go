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

	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

type ListOpCodeController struct {
	db *gorm.DB
	server server.IServerInterface
	listOpCodeService *services.ListOpCodeService
}

func NewListOpCodeController(srv *server.Server) *ListOpCodeController  {
	db := srv.DB
	return &ListOpCodeController{
		server: srv,
		db: db,
		listOpCodeService: services.NewListOpCodeService(&services.ListOpCodeService{
			ListOpCodeRepository: repository.NewListOpCodeRepository(srv.DB),
		})}	
}

// Daftar List Op Code Baru
// @Summary API untuk Membuat List Op Code Baru
// @Tags    List Op Code
// @Accept  json
// @Produce json
// @Param   listopcode body     dto.CreateListOpCode true "Daftar List Op Code Baru"
// @Success 201  {object} utils.Response
// @Failure 400  {object} middlewares.ResponseError
// @Failure 401  {object} middlewares.ResponseError
// @Failure 404  {object} middlewares.ResponseError
// @Failure 500  {object} middlewares.ResponseError
// @Router  /listOpCode/create [post]
func (listOpCodeController *ListOpCodeController) CreateListOpCode(c echo.Context) error  {
	req := new(dto.CreateListOpCode)

	if err := c.Bind(req); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, "Gagal melakukan binding code")
	}
	if err := c.Validate(req); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, "Validasi gagal")
	}
	

	data, err := listOpCodeController.listOpCodeService.CreateListOpCode(req)
	if err != nil {
		if errors.Is(err, utils.ErrOpCodeAlreadyExists) {
			res := utils.Response{
				Data:       nil,
				Message:    "OP Code Already Exist",
				StatusCode: 404,
			}
			return res.ReturnSingleMessage(c)
		}

		return err
	}

	res := utils.Response{
		Data: data,
		Message: "Berhasil membuat code",
		StatusCode: 201,
	}

	return res.ReturnSingleMessage(c)
}

// Ambil Daftar List Op Code
// @Summary API untuk Ambil semua daftar list op code
// @Tags    List Op code
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
// @Router  /listOpCode [get]
func (listOpCodeController *ListOpCodeController) FindListOpCode(c echo.Context) error  {
	req := new(dto.GetListOpCode)

	if err := c.Bind(req); err != nil {
		return err
	}
	if err := c.Validate(req); err != nil {
		return err
	}

	data, meta, err := listOpCodeController.listOpCodeService.FindListOpCode(req)
	if err != nil {
		return err
	}

	res := utils.ResponsePaginate {
		Key: "list_op_code",
		Meta: meta,
		Response: utils.Response{
			Data: data,
			Message: "Sukses mengambil data list op code",
			StatusCode: 200,
		},
	}

	return res.ReturnPaginates(c)
}


// Update List Op Code
// @Summary API untuk Update Data List OpCode
// @Tags    List Op Code
// @Accept  json
// @Produce json
// @Param   id   path     int               true "ID"
// @Param   listopcode body     dto.CreateListOpCode true "Update List Op Code Baru"
// @Success 201  {object} utils.Response
// @Failure 400  {object} middlewares.ResponseError
// @Failure 401  {object} middlewares.ResponseError
// @Failure 404  {object} middlewares.ResponseError
// @Failure 500  {object} middlewares.ResponseError
// @Router  /listOpCode/{id} [put]
func (listOpCodeController *ListOpCodeController) UpdateListOpCode(c echo.Context) error  {
	req := new(dto.UpdateListOpCode)
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

	// check id in database 
	_, errFind := listOpCodeController.listOpCodeService.FindListOpCodeById(uint(id))
	if errFind != nil {
		if errors.Is(errFind, utils.ErrListOpCodeNotFound) {
			res := utils.Response{
				Data: nil,
				Message: "Data not found",
				StatusCode: 404,
			}

			return res.ReturnSingleMessage(c)
		}

		return errFind
	}

	err := listOpCodeController.listOpCodeService.UpdateListOpCode(uint(id), req)
	if err != nil {
		return err
	}

	res := utils.Response{
		Data: nil,
		Message: "Berhasil Update Data",
		StatusCode: 201,
	}

	return res.ReturnSingleMessage(c)
}


// Delete List Op Code
// @Summary API untuk Delete List Op Code
// @Tags    List Op Code
// @Accept  json
// @Produce json
// @Param   id  path     int true "List Op ID"
// @Success 201 {object} utils.Response
// @Failure 400 {object} middlewares.ResponseError
// @Failure 401 {object} middlewares.ResponseError
// @Failure 404 {object} middlewares.ResponseError
// @Failure 500 {object} middlewares.ResponseError
// @Router  /listOpCode/{id} [delete]
func (listOpCodeController *ListOpCodeController) DeleteListOpCode(c echo.Context) error  {
	id, err := strconv.Atoi(c.Param("id"))
	
	err = listOpCodeController.listOpCodeService.DeleteListOpCode(uint(id))
	if err != nil {
		if errors.Is(err, utils.ErrListOpCodeNotFound) {
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
		Data: nil,
		Message: "Berhasil hapus data",
		StatusCode: 201,
	}

	return res.ReturnSingleMessage(c)
} 

// Find List Op Code Berdasarkan ID
// @Summary API untuk List Op Code By ID
// @Tags    List Op Code
// @Accept  json
// @Produce json
// @Param   id  path     string true "List Op ID"
// @Success 201 {object} utils.Response
// @Failure 400 {object} middlewares.ResponseError
// @Failure 401 {object} middlewares.ResponseError
// @Failure 404 {object} middlewares.ResponseError
// @Failure 500 {object} middlewares.ResponseError
// @Router  /listOpCode/{id} [GET]

func (listOpCodeController *ListOpCodeController) FindListOpCodeById(c echo.Context) error  {
	id, err := strconv.Atoi(c.Param("id"))

	if err != nil {
		return err
	}

	data, err := listOpCodeController.listOpCodeService.FindListOpCodeById(uint(id))
	if err != nil {
		if errors.Is(err, utils.ErrListOpCodeNotFound) {
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
		Data: data,
		Message: "Berhasil get data list op code",
		StatusCode: 200,
	}

	return res.ReturnSingleMessage(c)
}



