package auth

import "gitgub.com/diploma-mppr/backend_mppr/internal/app/models"

type Repository interface {
	GetUserById(id models.UserId) (*models.UserDB, error)
	GetUser(UserRequest *models.UserDB) (*models.UserDB, error)
	CreateUser(UserRequest *models.UserDB) (*models.UserDB, error)
	Login(UserRequest *models.UserDB) (*models.UserDB, error)
}
