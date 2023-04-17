package weightedSum

import "gitgub.com/diploma-mppr/backend_mppr/internal/app/models"

type UseCase interface {
	GetWeightedSum(id int) (*models.WeightedSumJson, error)
	SetWeightedSum(DataRequest *models.WeightedSumJson) (*models.WeightedSumJson, error)
	UpdateWeightedSum(DataRequest *models.WeightedSumJson) (*models.WeightedSumJson, error)
	DeleteWeightedSum(DataRequest *models.WeightedSumJson) error
}
