package borda

import "gitgub.com/diploma-mppr/backend_mppr/internal/app/models"

type UseCase interface {
	GetBorda(id int) (*models.BordaJson, error)
	SetBorda(DataRequest *models.BordaJson) (*models.BordaJson, error)
	UpdateBorda(DataRequest *models.BordaJson) (*models.BordaJson, error)
	DeleteBorda(DataRequest *models.BordaJson) error
}
