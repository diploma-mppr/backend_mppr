package baseCriteria

import "gitgub.com/diploma-mppr/backend_mppr/internal/app/models"

type Repository interface {
	GetBaseCriteria(id int, UserId int) (*models.BaseCriteriaDb, error)
	SetBaseCriteria(DataRequest *models.BaseCriteriaDb) (*models.BaseCriteriaDb, error)
	UpdateBaseCriteria(DataRequest *models.BaseCriteriaDb) (*models.BaseCriteriaDb, error)
	DeleteBaseCriteria(DataRequest *models.BaseCriteriaDb) error
}
