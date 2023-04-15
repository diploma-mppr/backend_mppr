package task

import "gitgub.com/diploma-mppr/backend_mppr/internal/app/models"

type UseCase interface {
	GetData(id int) (*models.DataJson, error)
	SetData(DataRequest *models.DataJson) (*models.DataJson, error)
}
