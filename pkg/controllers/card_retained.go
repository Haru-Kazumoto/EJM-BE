package controllers

import (
	"EJM/dto"
	"EJM/internal/server"
	"EJM/pkg/repository"
	"EJM/pkg/services"
	"EJM/utils"
	"strconv"

	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

type CardRetainedController struct {
	db                     *gorm.DB
	server                 server.IServerInterface
	cardRetainedService *services.CardRetainedService
}

func NewCardRetainedController(srv *server.Server) *CardRetainedController {
	db := srv.DB
	return &CardRetainedController{
		server: srv,
		db:     db,
		cardRetainedService: services.NewCardRetainedService(&services.CardRetainedService{
			CardRetainedRepository: repository.NewCardRetainedRepository(srv.DB),
		})}
}

// Get Card Retained
// @Summary API untuk Get Card Retained
// @Tags    Card Retained
// @Accept  json
// @Produce json
// @Param   term_id     query    string false "ID Terminal"
// @Param   atm_key     query    string false "ATM Key"
// @Param   term_type   query    string false "Jenis Mesin"
// @Param   merk        query    string false "Merk Mesin"
// @Param   managed_by  query    string false "Jenis Pengelola"
// @Param   jarkom      query    string false "Jenis Koneksi"
// @Param   region      query    string false "Region"
// @Param   area        query    string false "Area"
// @Success 200       {object} utils.ResponsePaginate
// @Failure 400       {object} middlewares.ResponseError
// @Failure 401       {object} middlewares.ResponseError
// @Failure 404       {object} middlewares.ResponseError
// @Failure 500       {object} middlewares.ResponseError
// @Router  /cardRetained [get]
func (cardRetainedController *CardRetainedController) GetCardRetained(c echo.Context) error {
	req := new(dto.GetCardRetained)

	if err := c.Bind(req); err != nil {
		return err
	}

	data, err := cardRetainedController.cardRetainedService.GetCardRetained(req)
	if err != nil {
		return err
	}

	// Menghitung total cardretained untuk mesin dengan Merk yang sesuai dengan filter
	totalCardRetainedByMerk := make(map[string]int)
	// Menghitung total cardretained untuk setiap area dengan region yang sesuai dengan filter
	totalCardRetainedByRegion := make(map[string]int)
	// Menghitung total cardretained untuk setiap managed_by dengan region yang sesuai dengan filter
	totalCardRetainedByManagedBy := make(map[string]int)

	for _, item := range data {
		// Cek apakah Region yang dimasukkan sesuai dengan data Region pada item
		if req.Region != "" && item.Region == req.Region {
			cardRetained, err := strconv.Atoi(item.CardRetainedTes)
			if err != nil {
				// Jika terjadi kesalahan konversi, beri nilai default 0
				cardRetained = 0
			}

			totalCardRetainedByRegion[item.Area] += cardRetained
			totalCardRetainedByMerk[item.Merk] += cardRetained
			totalCardRetainedByManagedBy[item.ManagedBy] += cardRetained
		}
	}

	// Membuat slice dari struct untuk mengisi data cardRetainedByRegion
	cardRetainedByRegionData := []struct {
		Area            string `json:"area"`
		TotalCardRetained int    `json:"totalCardRetained"`
	}{}

	for area, totalCardRetained := range totalCardRetainedByRegion {
		areaData := struct {
			Area            string `json:"area"`
			TotalCardRetained int    `json:"totalCardRetained"`
		}{
			Area:            area,
			TotalCardRetained: totalCardRetained,
		}
		cardRetainedByRegionData = append(cardRetainedByRegionData, areaData)
	}

	// Membuat slice dari struct untuk mengisi data cardRetainedByMerk
	cardRetainedByMerkData := []struct {
		Merk            string `json:"merk"`
		TotalCardRetained int    `json:"totalCardRetained"`
	}{}

	for merk, totalCardRetained := range totalCardRetainedByMerk {
		merkData := struct {
			Merk            string `json:"merk"`
			TotalCardRetained int    `json:"totalCardRetained"`
		}{
			Merk:            merk,
			TotalCardRetained: totalCardRetained,
		}
		cardRetainedByMerkData = append(cardRetainedByMerkData, merkData)
	}

	// Membuat slice dari struct untuk mengisi data cardRetainedByManagedBy
	cardRetainedByManagedByData := []struct {
		ManagedBy        string `json:"managed_by"`
		TotalCardRetained int    `json:"totalCardRetained"`
	}{}

	for managedBy, totalCardRetained := range totalCardRetainedByManagedBy {
		managedByData := struct {
			ManagedBy        string `json:"managed_by"`
			TotalCardRetained int    `json:"totalCardRetained"`
		}{
			ManagedBy:        managedBy,
			TotalCardRetained: totalCardRetained,
		}
		cardRetainedByManagedByData = append(cardRetainedByManagedByData, managedByData)
	}

	// Membuat respons JSON dengan data yang telah diolah
	response := struct {
		CardRetainedByMerk  []struct {
			Merk            string `json:"merk"`
			TotalCardRetained int    `json:"totalCardRetained"`
		} `json:"cardRetainedByMerk"`
		CardRetainedByRegion []struct {
			Area            string `json:"area"`
			TotalCardRetained int    `json:"totalCardRetained"`
		} `json:"cardRetainedByRegion"`
		CardRetainedByManagedBy []struct {
			ManagedBy        string `json:"managed_by"`
			TotalCardRetained int    `json:"totalCardRetained"`
		} `json:"cardRetainedByManagedBy"`
	}{
		CardRetainedByMerk:  cardRetainedByMerkData,
		CardRetainedByRegion: cardRetainedByRegionData,
		CardRetainedByManagedBy: cardRetainedByManagedByData,
	}

	res := utils.Response{
		Data:       response,
		Message:    "Berhasil Get Data Card Retained",
		StatusCode: 200,
	}

	return res.ReturnSingleMessage(c)
}
