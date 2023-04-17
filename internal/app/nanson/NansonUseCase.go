package nanson

import "gitgub.com/diploma-mppr/backend_mppr/internal/app/models"

type UseCase interface {
	GetNanson(id int) (*models.NansonJson, error)
	SetNanson(DataRequest *models.NansonJson) (*models.NansonJson, error)
	UpdateNanson(DataRequest *models.NansonJson) (*models.NansonJson, error)
	DeleteNanson(DataRequest *models.NansonJson) error
}
