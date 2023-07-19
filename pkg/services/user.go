package services

import (
	"EJM/config"
	"EJM/dto"
	"EJM/pkg/models"
	"EJM/pkg/repository"
	"EJM/utils"

	"errors"

	"gorm.io/gorm"
)

type RegisterService struct {
	RegisterRepository repository.UserRepository
	RoleRepository     repository.RoleRepository
	Config             *config.Config
	Db                 *gorm.DB
}

type IUserService interface {
	ToggleActive(id uint, payload *bool) error                              // [activating the user]
	RegisterUser(user *dto.CreateNewUser) (models.User, error)              // [register new user]
	FindUsers(users *dto.GetUsers) ([]models.User, *models.Paginate, error) //[get all users]
	UpdateUser(id uint, user *dto.UpdateUser) error   
	ChangePassword(id uint, newPassword string) error
	DeleteUser(id uint) error                                               // [delete user by id]
	FindUserById(id uint) (models.User, error)                              // [find user by id]
}

func NewRegisterService(constructor *RegisterService) *RegisterService {
	return constructor
}

func (register *RegisterService) ToggleActive(id uint, payload *bool) error {
	var registerRepo repository.UserRepository = register.RegisterRepository

	// registerRepo.Begin(register.Db)

	_, errFindUser := registerRepo.FindFirst(id)

	// kalo user gak ada
	if errFindUser != nil {
		return utils.ErrUserNotFound
	}

	if err := registerRepo.ToggleActive(id, payload); err != nil {
		return err
	}

	// defer registerRepo.Rollback() //Kalo ada error di defer wae

	// registerRepo.Commit()
	return nil
}

// // register user [tested]
func (register *RegisterService) RegisterUser(user *dto.CreateNewUser) (models.User, error) {
	user.Password, _ = utils.HashPassword(user.Password)

	var (
		registerRepo repository.UserRepository = register.RegisterRepository
		roleRepo     repository.RoleRepository = register.RoleRepository
	)

	// cari role dulu
	_, rolenotfound := roleRepo.FindRoleById(uint(user.RoleId))
	if rolenotfound != nil {
		if errors.Is(rolenotfound, gorm.ErrRecordNotFound) {
			return models.User{}, utils.ErrRoleNotExists
		} else {
			return models.User{}, rolenotfound
		}
	}

	// cari username dulu
	data, err := registerRepo.FindByUsername(user.Username)

	// err != nil berarti username belum ada
	if err != nil {
		createdUser, errCreateUser := registerRepo.CreateUser(user) //using services
		if errCreateUser != nil {
			return models.User{}, errCreateUser
		}

		return createdUser, nil

	}
	
	return data, utils.ErrUsernameExist
}

// find all usres [tested]
func (register *RegisterService) FindUsers(users *dto.GetUsers) (dto.GetUsers, *models.Paginate, error) {
	pagination := models.Paginate{
		Page:     users.Page,
		PageSize: users.PageSize,
	}

	var registerRepo repository.UserRepository = register.RegisterRepository

	data, meta, err := registerRepo.FindUsers(&pagination, users.Search, users.Value)
	if err != nil {
		return dto.GetUsers{}, meta, err
	}

	// Buat slice baru untuk menyimpan data pengguna dengan role name
	var usersWithRoleName []dto.CreateNewUserResponse

	// Lakukan iterasi untuk mengisi data pengguna dan mengambil role name berdasarkan RoleID
	for _, user := range data {
		// Pastikan Anda sudah menggunakan Preload untuk mengambil data role
		// saat melakukan query ke database (seperti pada contoh repository.FindUsers)
		userDTO := dto.CreateNewUserResponse{
			Name:       user.Name,
			Username:   user.Username,
			RoleId:     user.RoleId,
			Active:     dto.ActiveEnum(user.Active),
			RoleName:   user.Role.Name, // Mengambil role name dari relasi Role
		}

		usersWithRoleName = append(usersWithRoleName, userDTO)
	}

	responseDTO := dto.GetUsers{
		BasePagination: users.BasePagination,
		Users:          usersWithRoleName,
	}

	return responseDTO, meta, nil
}

// update user [tested]
func (register *RegisterService) UpdateUser(id uint, user *dto.UpdateUser) error {
	// cari id user
	var (
		registerRepo repository.UserRepository = register.RegisterRepository
		roleRepo     repository.RoleRepository = register.RoleRepository
	)

	_, errFindUser := registerRepo.FindFirst(id)
	if errFindUser != nil { // kalo user gak ada
		return utils.ErrUserNotFound
	}

	_, errFindRole := roleRepo.FindRoleById(uint(user.RoleId))
	if errFindRole != nil {
		return utils.ErrRoleNotExists
	}

	hashed, _ := utils.HashPassword(user.Password)
	user.Password = hashed

	return registerRepo.UpdateUser(id, user)
}

// update user [tested]
func (register *RegisterService) ChangePassword(id uint, newPassword string) error {
	// Cari id pengguna
	registerRepo := register.RegisterRepository
	_, err := registerRepo.FindFirst(id)
	if err != nil {
		return utils.ErrUserNotFound
	}

	return registerRepo.ChangePassword(id, newPassword)
}



// delete user [tested]
func (register *RegisterService) DeleteUser(id uint) error {
	var registerRepo repository.UserRepository = register.RegisterRepository

	// cari id dulu
	_, err := registerRepo.FindFirst(id)
	if err != nil {
		return utils.ErrUserNotFound
	}

	// delete
	return registerRepo.DeleteUser(id)
}

// Find User By Id
func (register *RegisterService) FindUserById(id uint) (models.User, error) {
	var userRepo repository.UserRepository = register.RegisterRepository

	data, err := userRepo.FindFirst(id)

	if err != nil {
		return models.User{}, utils.ErrUserNotFound
	}

	return data, nil
}
