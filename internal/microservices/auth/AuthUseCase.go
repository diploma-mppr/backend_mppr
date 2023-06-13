package auth

import "gitgub.com/diploma-mppr/backend_mppr/internal/app/models"

type UseCase interface {
	Register(User *models.UserJson) (*models.ResponseUserJson, error)
	Login(*models.UserJson) (*models.ResponseUserJson, error)
	GetUserById(id models.UserId) (models.ResponseUserJson, error)
}
