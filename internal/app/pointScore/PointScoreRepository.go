package pointScore

import "gitgub.com/diploma-mppr/backend_mppr/internal/app/models"

type Repository interface {
	GetPointScore(id int, UserId int) (*models.PointScoreDb, error)
	SetPointScore(DataRequest *models.PointScoreDb) (*models.PointScoreDb, error)
	UpdatePointScore(DataRequest *models.PointScoreDb) (*models.PointScoreDb, error)
	DeletePointScore(DataRequest *models.PointScoreDb) error
}
