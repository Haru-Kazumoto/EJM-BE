package controllers

import (
	"EJM/dto"
	"EJM/internal/server"
	"EJM/pkg/repository"
	"EJM/pkg/services"
	"EJM/utils"
	"github.com/nicksnyder/go-i18n/v2/i18n"
	"net/http"
	"time"

	"github.com/labstack/echo/v4"
)

type AuthController struct {
	server      *server.Server
	authService services.IAuthService
}

func NewAuthController(srv *server.Server) *AuthController {
	return &AuthController{
		server: srv,
		authService: services.NewAuthService(&services.AuthService{
			RegisterRepository: repository.NewRegisterRepository(srv.DB),
			RoleRepository:     repository.NewRoleRepository(srv.DB),
			JWTWhitelist: repository.JWTWhitelistImplementation(srv.DB),
			Config:       srv.Config,
		}),
	}
}

// Data Session
// @Summary Data Sesi
// @Tags    Auth
// @Accept  json
// @Produce json
// @Param   Accept-Language header   string false "Bahasa" Enums(en,id)
// @Success 201             {object} utils.Response{data=models.User}
// @Failure 400             {object} middlewares.ResponseError
// @Failure 401             {object} middlewares.ResponseError
// @Failure 404             {object} middlewares.ResponseError
// @Failure 500             {object} middlewares.ResponseError
// @Router  /session [get]
func (authController *AuthController) Session(c echo.Context) error {
	user := utils.User(c)
	data, err := authController.authService.Session(user.ID, utils.GetAccessToken(c))
	if err != nil {
		return err
	}

	res := utils.Response{
		Data: data,
		Translating: &i18n.Message{
			ID:    "session",
			Other: "Success get session data",
		},
		StatusCode: http.StatusOK,
	}

	return res.ReturnSingleMessage(c)

}

// Refresh Token
// @Summary Refresh JWT Token Yang Expired
// @Tags    Auth
// @Accept  json
// @Produce json
// @Param   Accept-Language header   string           false "Bahasa" Enums(en,id)
// @Param   login           body     dto.RefreshToken true  "Refresh Token"
// @Success 201             {object} utils.Response{data=services.LoginResponse}
// @Failure 400             {object} middlewares.ResponseError
// @Failure 401             {object} middlewares.ResponseError
// @Failure 404             {object} middlewares.ResponseError
// @Failure 500             {object} middlewares.ResponseError
// @Router  /auth/refresh-token [post]
func (authController *AuthController) RefreshToken(c echo.Context) error {
	req := new(dto.RefreshToken)
	if err := c.Bind(req); err != nil {
		return err
	}
	if err := c.Validate(req); err != nil {
		return err
	}

	loginService, err := authController.authService.RefreshToken(req.Token)
	if err != nil {
		return err
	}

	res := utils.Response{
		Data:       loginService,
		Message:    "Berhasil Membuat Token Baru",
		StatusCode: 200,
	}

	c.SetCookie(&http.Cookie{
		Name:    "Token",
		Value:   loginService.Token,
		Expires: time.Now().Add(time.Duration(authController.server.Config.Auth.JWTExpired) * time.Minute),
	})
	c.SetCookie(&http.Cookie{
		Name:    "RefreshToken",
		Value:   loginService.RefreshToken,
		Expires: time.Now().Add(time.Duration((authController.server.Config.Auth.JWTExpired)+60) * time.Minute),
	})
	return res.ReturnSingleMessage(c)
}

// Login By Link
// @Summary Autentikasi User untuk Resource Private Server
// @Tags    Auth
// @Accept  json
// @Produce json
// @Param   Accept-Language header   string         false "Bahasa" Enums(en,id)
// @Param   login           body     dto.LoginByPin true  "Login"
// @Success 201             {object} utils.Response{data=services.LoginResponse}
// @Failure 400             {object} middlewares.ResponseError
// @Failure 401             {object} middlewares.ResponseError
// @Failure 404             {object} middlewares.ResponseError
// @Failure 500             {object} middlewares.ResponseError
// @Router  /auth/oauth/login [post]
func (authController *AuthController) OauthLogin(c echo.Context) error {
	req := new(dto.LoginByPin)

	if err := c.Bind(req); err != nil {
		return err
	}
	if err := c.Validate(req); err != nil {
		return err
	}

	loginService, err := authController.authService.LoginByPin(req)
	if err != nil {
		return err
	}

	res := utils.Response{
		Data:       loginService,
		Message:    "Berhasil Login",
		StatusCode: 200,
	}

	c.SetCookie(&http.Cookie{
		Name:    "Token",
		Value:   loginService.Token,
		Path:    "/",
		Expires: time.Now().Add(time.Duration(authController.server.Config.Auth.JWTExpired) * time.Minute),
	})
	c.SetCookie(&http.Cookie{
		Name:    "RefreshToken",
		Value:   loginService.RefreshToken,
		Path:    "/",
		Expires: time.Now().Add(time.Duration((authController.server.Config.Auth.JWTExpired)+60) * time.Minute),
	})

	authController.server.Casbin.LoadPolicy()

	return res.ReturnSingleMessage(c)
}

// Login
// @Summary Autentikasi User untuk Resource Private Server
// @Tags    Auth
// @Accept  json
// @Produce json
// @Param   Accept-Language header   string    false "Bahasa" Enums(en,id)
// @Param   login           body     dto.Login true  "Login"
// @Success 201             {object} utils.Response{data=services.LoginResponse}
// @Failure 400             {object} middlewares.ResponseError
// @Failure 401             {object} middlewares.ResponseError
// @Failure 404             {object} middlewares.ResponseError
// @Failure 500             {object} middlewares.ResponseError
// @Router  /auth/login [post]
func (authController *AuthController) LoginUser(c echo.Context) error {
	req := new(dto.Login)

	if err := c.Bind(req); err != nil {
		return err
	}
	if err := c.Validate(req); err != nil {
		return err
	}

	loginService, err := authController.authService.LoginUser(req)
	if err != nil {
		return err
	}

	res := utils.Response{
		Data:       loginService,
		Message:    "Berhasil Login",
		StatusCode: 200,
	}

	c.SetCookie(&http.Cookie{
		Name:    "Token",
		Value:   loginService.Token,
		Path:    "/",
		Expires: time.Now().Add(time.Duration(authController.server.Config.Auth.JWTExpired) * time.Minute),
	})
	c.SetCookie(&http.Cookie{
		Name:    "RefreshToken",
		Value:   loginService.RefreshToken,
		Path:    "/",
		Expires: time.Now().Add(time.Duration((authController.server.Config.Auth.JWTExpired)+60) * time.Minute),
	})

	authController.server.Casbin.LoadPolicy()

	return res.ReturnSingleMessage(c)
}

func (authController *AuthController) Logout(c echo.Context) error {
	c.SetCookie(&http.Cookie{
		Name:    "Token",
		Value:   "",
		Path:    "/",
		Expires: time.Now(),
	})

	c.SetCookie(&http.Cookie{
		Name:    "RefreshToken",
		Value:   "",
		Path:    "/",
		Expires: time.Now(),
	})

	var authService = authController.authService
	errLogout := authService.Logout(utils.GetAccessToken(c))
	if errLogout != nil {
		return errLogout
	}

	res := utils.Response{
		Message:    "Berhasil Logout",
		StatusCode: 200,
	}

	return res.ReturnSingleMessage(c)
}
