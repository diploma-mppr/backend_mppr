package data

import "gitgub.com/diploma-mppr/backend_mppr/internal/app/models"

type UseCase interface {
	GetAll(UserId int) (*[]models.DataDb, error)
}
