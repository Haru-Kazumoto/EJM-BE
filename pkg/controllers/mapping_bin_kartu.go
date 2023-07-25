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

type BinKartuController struct {
	db                     *gorm.DB
	server                 server.IServerInterface
	binKartuService *services.BinKartuService
}

func NewBinKartuController(srv *server.Server) *BinKartuController {
	db := srv.DB
	return &BinKartuController{
		server: srv,
		db:     db,
		binKartuService: services.NewBinKartuService(&services.BinKartuService{
			BinKartuRepository: repository.NewBinKartuRepository(srv.DB),
		})}
}

// Daftar Bin Kartu Baru
// @Summary API untuk Membuat Bin Kartu Baru
// @Tags    Bin Kartu
// @Accept  json
// @Produce json
// @Param   binKartu body     dto.CreateNewBinKartu true "Daftar Bin Kartu Baru"
// @Success 201  {object} utils.Response
// @Failure 400  {object} middlewares.ResponseError
// @Failure 401  {object} middlewares.ResponseError
// @Failure 404  {object} middlewares.ResponseError
// @Failure 500  {object} middlewares.ResponseError
// @Router  /binKartu/create [post]
func (binKartuController *BinKartuController) CreateBinKartu(c echo.Context) error {
	req := new(dto.CreateNewBinKartu)
	if err := c.Bind(req); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, "Gagal melakukan binding data")
	}

	if err := c.Validate(req); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, "Validasi gagal")
	}

	data, err := binKartuController.binKartuService.CreateBinKartu(req)
	if err != nil {
		if errors.Is(err, utils.ErrPrefixNoAlreadyExists) {
			res := utils.Response{
				Data:       nil,
				Message:    "Prefix No Already Exist",
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

// Ambil Daftar Bin Kartu
// @Summary API untuk Ambil semua daftar Bin Kartu
// @Tags    Bin Kartu
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
// @Router  /binKartu [get]
func (binKartuController *BinKartuController) FindBinKartu(c echo.Context) error {
	req := new(dto.GetBinKartu)

	if err := c.Bind(req); err != nil {
		return err
	}
	if err := c.Validate(req); err != nil {
		return err
	}

	data, meta, err := binKartuController.binKartuService.FindBinKartu(req)
	if err != nil {
		return err
	}

	res := utils.ResponsePaginate{
		Key:  "bin_kartu",
		Meta: meta,
		Response: utils.Response{
			Data:       data,
			Message:    "Sukses mengambil data mapping bin kartu",
			StatusCode: 200,
		},
	}

	return res.ReturnPaginates(c)
}

// Update Bin Kartu
// @Summary API untuk update bin kartu
// @Tags    Bin Kartu
// @Accept  json
// @Produce json
// @Param   id   path     int               true "ID"
// @Param   binKartu body     dto.CreateNewBinKartu     true "Update Bin Kartu"
// @Success 201  {object} utils.Response
// @Failure 400  {object} middlewares.ResponseError
// @Failure 401  {object} middlewares.ResponseError
// @Failure 404  {object} middlewares.ResponseError
// @Failure 500  {object} middlewares.ResponseError
// @Router  /binKartu/{id} [put]
func (binKartuController *BinKartuController) UpdateBinKartu(c echo.Context) error {
	req := new(dto.UpdateBinKartu)
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
	_, errFind := binKartuController.binKartuService.FindBinKartuById(uint(id))
	if errFind != nil {
		if errors.Is(errFind, utils.ErrBinKartuNotFound) {
			res := utils.Response{
				Data:       nil,
				Message:    "Data Not Found",
				StatusCode: 404,
			}

			return res.ReturnSingleMessage(c)
		}
		return errFind
	}

	err := binKartuController.binKartuService.UpdateBinKartu(uint(id), req)
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

// Delete Bin Kartu
// @Summary API untuk Delete Bin Kartu
// @Tags    Bin Kartu
// @Accept  json
// @Produce json
// @Param   id  path     int true "ID"
// @Success 201 {object} utils.Response
// @Failure 400 {object} middlewares.ResponseError
// @Failure 401 {object} middlewares.ResponseError
// @Failure 404 {object} middlewares.ResponseError
// @Failure 500 {object} middlewares.ResponseError
// @Router  /binKartu/{id} [delete]
func (binKartuController *BinKartuController) DeleteBinKartu(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))

	// binKartuRepo := binKartuController.binKartuService.BinKartuRepository

	// binKartuRepo.Begin(binKartuController.db)

	err = binKartuController.binKartuService.DeleteBinKartu(uint(id))
	if err != nil {
		if errors.Is(err, utils.ErrBinKartuNotFound) {
			res := utils.Response{
				Data:       nil,
				Message:    "Data Not Found",
				StatusCode: 404,
			}

			return res.ReturnSingleMessage(c)
		}
		return err
	}

	// binKartuRepo.Commit()

	res := utils.Response{
		Data:       nil,
		Message:    "Berhasil Hapus Data",
		StatusCode: 201,
	}

	return res.ReturnSingleMessage(c)
}

// Find Bin Kartu Berdasarkan ID
// @Summary API untuk Find Bin Kartu By ID
// @Tags    Bin Kartu
// @Accept  json
// @Produce json
// @Param   id  path     string true "ID"
// @Success 201 {object} utils.Response
// @Failure 400 {object} middlewares.ResponseError
// @Failure 401 {object} middlewares.ResponseError
// @Failure 404 {object} middlewares.ResponseError
// @Failure 500 {object} middlewares.ResponseError
// @Router  /binKartu/{id} [GET]
func (binKartuController *BinKartuController) FindBinKartuById(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))

	if err != nil {
		return err
	}

	data, err := binKartuController.binKartuService.FindBinKartuById(uint(id))
	if err != nil {
		if errors.Is(err, utils.ErrBinKartuNotFound) {
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
		Message:    "Berhasil Get Data Mapping Bin Kartu",
		StatusCode: 200,
	}

	return res.ReturnSingleMessage(c)
}