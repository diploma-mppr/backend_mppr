package nanson

import "gitgub.com/diploma-mppr/backend_mppr/internal/app/models"

type Repository interface {
	GetNanson(id int) (*models.NansonDb, error)
	SetNanson(DataRequest *models.NansonDb) (*models.NansonDb, error)
	UpdateNanson(DataRequest *models.NansonDb) (*models.NansonDb, error)
	DeleteNanson(DataRequest *models.NansonDb) error
}
