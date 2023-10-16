package repository

import (
	"Valter/db/sqlc"
	"context"
	"database/sql"
)

type UserRepository interface {
	AddNewUser(arg sqlc.AddNewUserParams) (sql.Result, error)
	GetAllUsers() ([]sqlc.User, error)
	GetUserbyId(id uint32) (sqlc.User, error)
	GetUserbyUsername(username string) (sqlc.User, error)
	GetUserbyEmail(email string) (sqlc.User, error)
	UpdateUser(arg sqlc.UpdateUserParams) error
	ForgotPass(arg sqlc.ForgotPassParams) error
	SetToken(arg sqlc.SetTokenParams) error
	DeleteUser(id uint32) error
}

type userRepository struct {
	db *sqlc.Queries
}

func NewUserRepository(db *sqlc.Queries) UserRepository {
	return &userRepository{db}
}

func (u *userRepository) AddNewUser(arg sqlc.AddNewUserParams) (sql.Result, error) {
	result, err := u.db.AddNewUser(context.Background(), arg)
	if err != nil {
		return nil, err
	}
	return result, nil
}

func (u *userRepository) GetAllUsers() ([]sqlc.User, error) {
	result, err := u.db.GetAllUsers(context.Background())
	if err != nil {
		return nil, err
	}
	return result, nil
}

func (u *userRepository) GetUserbyId(id uint32) (sqlc.User, error) {
	result, err := u.db.GetUserbyId(context.Background(), id)
	if err != nil {
		return result, err
	}
	return result, nil
}

func (u *userRepository) GetUserbyUsername(username string) (sqlc.User, error) {
	result, err := u.db.GetUserbyUsername(context.Background(), username)
	if err != nil {
		return result, err
	}
	return result, nil
}

func (u *userRepository) GetUserbyEmail(email string) (sqlc.User, error) {
	result, err := u.db.GetUserbyEmail(context.Background(), email)
	if err != nil {
		return result, err
	}
	return result, nil
}

func (u *userRepository) UpdateUser(arg sqlc.UpdateUserParams) error {
	err := u.db.UpdateUser(context.Background(), arg)
	if err != nil {
		return err
	}
	return nil
}

func (u *userRepository) ForgotPass(arg sqlc.ForgotPassParams) error {
	err := u.db.ForgotPass(context.Background(), arg)
	if err != nil {
		return err
	}
	return nil
}

func (u *userRepository) SetToken(arg sqlc.SetTokenParams) error {
	err := u.db.SetToken(context.Background(), arg)
	if err != nil {
		return err
	}
	return nil
}

func (u *userRepository) DeleteUser(id uint32) error {
	err := u.db.DeleteUser(context.Background(), id)
	if err != nil {
		return err
	}
	return nil
}
