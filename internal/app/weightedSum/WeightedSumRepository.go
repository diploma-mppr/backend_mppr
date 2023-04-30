package weightedSum

import "gitgub.com/diploma-mppr/backend_mppr/internal/app/models"

type Repository interface {
	GetWeightedSum(id int, UserId int) (*models.WeightedSumDb, error)
	SetWeightedSum(DataRequest *models.WeightedSumDb) (*models.WeightedSumDb, error)
	UpdateWeightedSum(DataRequest *models.WeightedSumDb) (*models.WeightedSumDb, error)
	DeleteWeightedSum(DataRequest *models.WeightedSumDb) error
}
