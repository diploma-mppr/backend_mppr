package borda

import "gitgub.com/diploma-mppr/backend_mppr/internal/app/models"

type Repository interface {
	GetBorda(id int, UserId int) (*models.BordaDb, error)
	SetBorda(DataRequest *models.BordaDb) (*models.BordaDb, error)
	UpdateBorda(DataRequest *models.BordaDb) (*models.BordaDb, error)
	DeleteBorda(DataRequest *models.BordaDb) error
}
