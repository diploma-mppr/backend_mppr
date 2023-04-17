package pointScore

import "gitgub.com/diploma-mppr/backend_mppr/internal/app/models"

type UseCase interface {
	GetPointScore(id int) (*models.PointScoreJson, error)
	SetPointScore(DataRequest *models.PointScoreJson) (*models.PointScoreJson, error)
	UpdatePointScore(DataRequest *models.PointScoreJson) (*models.PointScoreJson, error)
	DeletePointScore(DataRequest *models.PointScoreJson) error
}
