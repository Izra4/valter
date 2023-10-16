package service

import (
	"Valter/db/sqlc"
	"Valter/repository"
	"Valter/utility"
	"database/sql"
	"errors"
	"github.com/gin-gonic/gin"
)

type UserService interface {
	AddNewUser(c *gin.Context, username string, email string, number string, address string, password string) (sql.Result, error)
	GetAllUsers() ([]sqlc.User, error)
	GetUserbyId(id uint32) (sqlc.User, error)
	GetUserbyUsername(username string) (sqlc.User, error)
	GetUserbyEmail(email string) (sqlc.User, error)
	UpdateUser(c *gin.Context, id uint32, newUsername string, newEmail string, newNumber string, newAddress string) error
	ForgotPass(c *gin.Context, id uint32, password string) error
	SetToken(id uint32, token string) error
	DeleteUser(c *gin.Context, id uint32) error
	LoginUser(c *gin.Context, username string, password string) (string, error)
}

type userService struct {
	repository repository.UserRepository
}

func NewUserService(repository repository.UserRepository) UserService {
	return &userService{repository}
}

func (u *userService) AddNewUser(c *gin.Context, username string, email string, number string, address string, password string) (sql.Result, error) {
	if len(password) < 8 {
		utility.HttpBadRequest(c, "password length less than 8")
		return nil, errors.New("password length less than 8")
	}

	hashedPass, err := utility.Hashing(password)
	if err != nil {
		utility.HttpInternalErrorResponse(c, "Failed to hash password", err)
		return nil, err
	}
	data := sqlc.AddNewUserParams{
		Username: username,
		Email:    email,
		Number:   number,
		Address:  address,
		Password: hashedPass,
	}

	return u.repository.AddNewUser(data)
}

func (u *userService) GetAllUsers() ([]sqlc.User, error) {
	return u.repository.GetAllUsers()
}

func (u *userService) GetUserbyId(id uint32) (sqlc.User, error) {
	return u.repository.GetUserbyId(id)
}

func (u *userService) GetUserbyUsername(username string) (sqlc.User, error) {
	return u.repository.GetUserbyUsername(username)
}

func (u *userService) GetUserbyEmail(email string) (sqlc.User, error) {
	return u.repository.GetUserbyEmail(email)
}

func (u *userService) UpdateUser(c *gin.Context, id uint32, newUsername string, newEmail string, newNumber string, newAddress string) error {
	data, err := u.repository.GetUserbyId(id)
	if err != nil {
		utility.HttpDataNotFound(c, "Data not found", err)
		return err
	}

	if newUsername == "" {
		newUsername = data.Username
	}
	if newEmail == "" {
		newEmail = data.Email
	}
	if newNumber == "" {
		newNumber = data.Number
	}
	if newAddress == "" {
		newAddress = data.Address
	}
	input := sqlc.UpdateUserParams{
		Username: newUsername,
		Email:    newEmail,
		Number:   newNumber,
		Address:  newAddress,
		ID:       id,
	}
	return u.repository.UpdateUser(input)
}

func (u *userService) ForgotPass(c *gin.Context, id uint32, password string) error {
	_, err := u.repository.GetUserbyId(id)
	if err != nil {
		utility.HttpDataNotFound(c, "Data not found", err)
		return err
	}
	hashedPass, err := utility.Hashing(password)
	if err != nil {
		utility.HttpInternalErrorResponse(c, "Failed to hash", err)
		return err
	}
	data := sqlc.ForgotPassParams{
		Password: hashedPass,
		ID:       id,
	}
	return u.repository.ForgotPass(data)

}

func (u *userService) SetToken(id uint32, token string) error {
	data := sqlc.SetTokenParams{
		Token: token,
		ID:    id,
	}
	return u.repository.SetToken(data)
}

func (u *userService) DeleteUser(c *gin.Context, id uint32) error {
	_, err := u.repository.GetUserbyId(id)
	if err != nil {
		utility.HttpDataNotFound(c, "Data not found", err)
		return err
	}
	if err = u.repository.DeleteUser(id); err != nil {
		utility.HttpInternalErrorResponse(c, "Failed to delete user", err)
		return err
	}
	return nil
}

func (u *userService) LoginUser(c *gin.Context, username string, password string) (string, error) {
	if username == "" || password == "" {
		utility.HttpBadRequest(c, "Fill the empty field")
		return "", errors.New("username is empty")
	}
	data, err := u.repository.GetUserbyUsername(username)
	if err != nil {
		utility.HttpDataNotFound(c, "invalid username / passowrd", err)
		return "", err
	}
	err = utility.ComparePass(data.Password, password)
	if err != nil {
		utility.HttpBadRequest(c, "invalid username / password")
		return "", err
	}
	token, err := utility.Token(data)
	if err != nil {
		utility.HttpInternalErrorResponse(c, "Failed to create token", err)
		return "", err
	}
	return token, nil
}
