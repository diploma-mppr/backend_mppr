package pareto

import "gitgub.com/diploma-mppr/backend_mppr/internal/app/models"

type UseCase interface {
	GetPareto(id int) (*models.ParetoJson, error)
	SetPareto(DataRequest *models.ParetoJson) (*models.ParetoJson, error)
	UpdatePareto(DataRequest *models.ParetoJson) (*models.ParetoJson, error)
	DeletePareto(DataRequest *models.ParetoJson) error
}
