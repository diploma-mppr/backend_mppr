package pareto

import "gitgub.com/diploma-mppr/backend_mppr/internal/app/models"

type Repository interface {
	GetPareto(id int) (*models.ParetoDb, error)
	SetPareto(DataRequest *models.ParetoDb) (*models.ParetoDb, error)
	UpdatePareto(DataRequest *models.ParetoDb) (*models.ParetoDb, error)
	DeletePareto(DataRequest *models.ParetoDb) error
}
