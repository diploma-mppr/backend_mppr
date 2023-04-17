package baseCriteria

import "gitgub.com/diploma-mppr/backend_mppr/internal/app/models"

type UseCase interface {
	GetBaseCriteria(id int) (*models.BaseCriteriaJson, error)
	SetBaseCriteria(DataRequest *models.BaseCriteriaJson) (*models.BaseCriteriaJson, error)
	UpdateBaseCriteria(DataRequest *models.BaseCriteriaJson) (*models.BaseCriteriaJson, error)
	DeleteBaseCriteria(DataRequest *models.BaseCriteriaJson) error
}
