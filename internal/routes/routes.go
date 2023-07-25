package routes

import (
	"EJM/config"
	"EJM/internal/middlewares"
	"EJM/internal/request_logger"
	"EJM/internal/server"
	"EJM/pkg/controllers"
	"EJM/utils"
	"encoding/json"
	"fmt"
	"os"
	"strconv"
	"time"

	"github.com/go-playground/validator/v10"
	casbin_mw "github.com/labstack/echo-contrib/casbin"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/sirupsen/logrus"
)

func InitializeRoute(server *server.Server, cfg *config.Config) {
	// controller
	userController := controllers.NewUserController(server)
	roleController := controllers.NewRoleController(server)
	authController := controllers.NewAuthController(server)
	menuController := controllers.NewMenuController(server)
	actionController := controllers.NewActionController(server)
	mappingCodeController := controllers.NewMappingCodeController(server)
	mappingKeywordListController := controllers.NewMappingKeywordListcontroller(server)
	jenisTransaksiController := controllers.NewMJenisTransaksiController(server)
	listOpCodeController := controllers.NewListOpCodeController(server)
	binKartuController := controllers.NewBinKartuController(server)
	cardRetainedController := controllers.NewCardRetainedController(server)


	// middleware
	server.Echo.Validator = &middlewares.CustomValidator{
		Validator: validator.New(),
	}

	// Logger
	request_logger.Logger().SetOutput(os.Stdout)
	request_logger.Logger().SetFormatter(&logrus.TextFormatter{
		TimestampFormat: time.RFC3339,
		ForceColors:     true,
	})
	server.Echo.Logger = request_logger.Logger()
	server.Echo.Use(middlewares.LoggerLogrus())

	server.Echo.HTTPErrorHandler = middlewares.HttpErrorHandler(server.Config)
	server.Echo.Use(middleware.RequestID())
	server.Echo.Use(middlewares.ContextDB(server.DB))
	server.Echo.Use(middleware.Gzip())
	server.Echo.Use(middleware.Recover())
	server.Echo.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins:     []string{"*"},
		AllowCredentials: true,
		ExposeHeaders:    []string{"Content-Disposition"},
		AllowHeaders: []string{
			echo.HeaderOrigin,
			echo.HeaderAuthorization,
			echo.HeaderContentType,
			"module",
			"Content-Range",
			"Accept-Language",
		},
	}))

	prefix := server.Echo.Group(fmt.Sprintf("/api/v%s", server.Config.App.Version))

	if server.Config.App.Env == "development" {
		prefix.GET("/test-queue", func(c echo.Context) error {
			data, _ := json.Marshal(map[string]interface{}{
				"project_id": 11,
				"user_id":    3,
				"volume_id":  nil,
			})

			err := server.Queue.VolumeAttachment.Publish(string(data))
			if err != nil {
				return c.String(400, err.Error())
			}

			return c.String(200, fmt.Sprintf("Job Queue Sent %s", string(data)))
		})
	}

	prefix.Static("/profiles", "assets/profiles")
	prefix.Static("/reports", "assets/reports")
	prefix.Static("/assets/volumes", "assets/volumes")

	/**
	Public Route
	*/
	authRoutes := prefix.Group("/auth")
	{
		authRoutes.POST("/login", authController.LoginUser)
		authRoutes.POST("/refresh-token", authController.RefreshToken)
		authRoutes.POST("/register", userController.RegisterUser)
		authRoutes.Use(middleware.JWTWithConfig(middleware.JWTConfig{
			SigningKey:  []byte(cfg.Auth.JWTKey),
			Claims:      &utils.UserJWTPayload{},
			TokenLookup: "header:Authorization,cookie:Token",
		}))
		{
			authRoutes.POST("/logout", authController.Logout)
		}
	}

	/**
	Protected Route
	*/
	// authorization middleware
	prefix.Use(middleware.JWTWithConfig(middleware.JWTConfig{
		SigningKey:  []byte(cfg.Auth.JWTKey),
		Claims:      &utils.UserJWTPayload{},
		TokenLookup: "header:Authorization,cookie:Token",
	}))
	{
		prefix.GET("/session", authController.Session)

		prefix.Use(casbin_mw.MiddlewareWithConfig(casbin_mw.Config{
			Enforcer: server.Casbin,
			UserGetter: func(c echo.Context) (string, error) {
				return strconv.Itoa(utils.User(c).RoleID), nil
			},
			ErrorHandler: func(c echo.Context, internal error, proposedStatus int) error {
				return utils.ErrUnauthorizedRole
			},
		}))

		// User
		userRoutes := prefix.Group("/users")
		{
			userRoutes.POST("/register", userController.RegisterUser)
			userRoutes.GET("", userController.FindUsers)
			userRoutes.PUT("/status/:id", userController.ToggleActiveNonActive)
			userRoutes.PUT("/:id", userController.UpdateUser)
			userRoutes.PUT("/changePw/:id", userController.ChangePassword)
			userRoutes.GET("/:id", userController.FindUserById)
			userRoutes.DELETE("/:id", userController.DeleteUser)
			userRoutes.DELETE("", userController.DeleteUserBulk)
		}

		// menus
		menuRoutes := prefix.Group("/menus")
		{
			menuRoutes.GET("", menuController.GetAll)
			menuRoutes.GET("/:id", menuController.FindDetailMenu)
			menuRoutes.PUT("/:id", menuController.UpdateMenu)
			menuRoutes.DELETE("/:id", menuController.DeleteMenu)
			menuRoutes.DELETE("/bulk", menuController.DeleteMenuBulk)
			menuRoutes.GET("/paginated", menuController.GetAllPaginated)
			menuRoutes.POST("/create-bulk", menuController.CreateMenuBulk)
		}

		// actions
		actionRoutes := prefix.Group("/actions")
		{
			actionRoutes.GET("/:id", actionController.FindDetailAction)
			actionRoutes.PUT("/:id", actionController.UpdateAction)
			actionRoutes.DELETE("/:id", actionController.DeleteAction)
			actionRoutes.DELETE("/bulk", actionController.DeleteActionBulk)
			actionRoutes.GET("/paginated", actionController.GetAllPaginated)
			actionRoutes.GET("/menu/:menu_id", actionController.FindActionByMenu)
			actionRoutes.POST("/create-bulk", actionController.CreateActionBulk)
		}

		// Role
		roleRoutes := prefix.Group("/roles")
		{
			roleRoutes.POST("", roleController.CreateRole)
			roleRoutes.GET("", roleController.GetRoles)
			roleRoutes.DELETE("/delete/:roleId", roleController.DeleteRole)
			roleRoutes.DELETE("/delete-bulk", roleController.DeleteRoleBulk)
			roleRoutes.PATCH("/:id", roleController.UpdateRole)
			roleRoutes.GET("/:id", roleController.FindRoleById)

			roleRoutes.GET("/accessroles", roleController.GetAccessRole)
			roleRoutes.POST("/set-access-role", roleController.SetAccessRole)
			roleRoutes.DELETE("/delete-access-role", roleController.DeleteAccessRole)
		}

		//Mapping Code
		mappingCodeRoutes := prefix.Group("/mappingCodes")
		{
			mappingCodeRoutes.GET("", mappingCodeController.FindMappingCodes)
			mappingCodeRoutes.GET("/:id", mappingCodeController.FindMappingCodeById)
			mappingCodeRoutes.POST("/create", mappingCodeController.CreateMappingCode)
			mappingCodeRoutes.PUT("/:id", mappingCodeController.UpdateMappingCode)
			mappingCodeRoutes.DELETE("/:id", mappingCodeController.DeleteMappingCode)
			mappingCodeRoutes.DELETE("", mappingCodeController.DeleteMappingCodeBulk)
		}

		// Mapping Keyword List
		mappingKeywordListRoutes := prefix.Group("/mappingKeywordList")
		{
			mappingKeywordListRoutes.GET("", mappingKeywordListController.FindMappingkeywordlist)
			mappingKeywordListRoutes.GET("/:id", mappingKeywordListController.FindMappingkeywordlistById)
			mappingKeywordListRoutes.POST("/create", mappingKeywordListController.CreateMappingKeywordList)
			mappingKeywordListRoutes.PUT("/:id", mappingKeywordListController.UpdateMappingkeywordlist)
			mappingKeywordListRoutes.DELETE("/:id", mappingKeywordListController.DeleteMappingkeywordlist)
		}

		//Jenis Transaksi
		jenisTransaksiRoutes := prefix.Group("/jenisTransaksi")
		{
			jenisTransaksiRoutes.GET("", jenisTransaksiController.FindJenisTransaksi)
			jenisTransaksiRoutes.GET("/:id", jenisTransaksiController.FindJenisTransaksiById)
			jenisTransaksiRoutes.POST("/create", jenisTransaksiController.CreateJenisTransaksi)
			jenisTransaksiRoutes.PUT("/:id", jenisTransaksiController.UpdateJenisTransaksi)
			jenisTransaksiRoutes.DELETE("/:id", jenisTransaksiController.DeleteJenisTransaksi)
			// jenisTransaksiRoutes.DELETE("", jenisTransaksiController.DeleteMappingCodeBulk)
		}

		// list op code
		listOpCodeRoute := prefix.Group("/listOpCode")
		{
			listOpCodeRoute.GET("", listOpCodeController.FindListOpCode)
			listOpCodeRoute.GET("/:id", listOpCodeController.FindListOpCodeById)
			listOpCodeRoute.POST("/create", listOpCodeController.CreateListOpCode)
			listOpCodeRoute.PUT("/:id", listOpCodeController.UpdateListOpCode)
			listOpCodeRoute.DELETE("/:id", listOpCodeController.DeleteListOpCode)
		}
		
		binKartuRoute := prefix.Group("/binKartu")
		{
			binKartuRoute.GET("", binKartuController.FindBinKartu)
			binKartuRoute.GET("/:id", binKartuController.FindBinKartuById)
			binKartuRoute.POST("/create", binKartuController.CreateBinKartu)
			binKartuRoute.PUT("/:id", binKartuController.UpdateBinKartu)
			binKartuRoute.DELETE("/:id", binKartuController.DeleteBinKartu)
		}

		cardRetainedRoute := prefix.Group("/cardRetained")
		{
			cardRetainedRoute.GET("", cardRetainedController.GetCardRetained)
		}
	}

}
