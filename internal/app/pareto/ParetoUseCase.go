package pareto

import "gitgub.com/diploma-mppr/backend_mppr/internal/app/models"

type UseCase interface {
	GetPareto(id int, UserId int) (*models.ParetoJson, error)
	SetPareto(DataRequest *models.ParetoJson, UserId int) (*models.ParetoJson, error)
	UpdatePareto(DataRequest *models.ParetoJson, UserId int) (*models.ParetoJson, error)
	DeletePareto(DataRequest *models.ParetoJson, UserId int) error
}
