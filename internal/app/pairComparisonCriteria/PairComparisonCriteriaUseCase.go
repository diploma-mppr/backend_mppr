package pairComparisonCriteria

import "gitgub.com/diploma-mppr/backend_mppr/internal/app/models"

type UseCase interface {
	GetPairComparisonCriteria(id int) (*models.PairComparisonCriteriaJson, error)
	SetPairComparisonCriteria(DataRequest *models.PairComparisonCriteriaJson) (*models.PairComparisonCriteriaJson, error)
	UpdatePairComparisonCriteria(DataRequest *models.PairComparisonCriteriaJson) (*models.PairComparisonCriteriaJson, error)
	DeletePairComparisonCriteria(DataRequest *models.PairComparisonCriteriaJson) error
}
