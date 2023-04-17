package pairComparisonCriteria

import "gitgub.com/diploma-mppr/backend_mppr/internal/app/models"

type Repository interface {
	GetPairComparisonCriteria(id int) (*models.PairComparisonCriteriaDb, error)
	SetPairComparisonCriteria(DataRequest *models.PairComparisonCriteriaDb) (*models.PairComparisonCriteriaDb, error)
	UpdatePairComparisonCriteria(DataRequest *models.PairComparisonCriteriaDb) (*models.PairComparisonCriteriaDb, error)
	DeletePairComparisonCriteria(DataRequest *models.PairComparisonCriteriaDb) error
}
