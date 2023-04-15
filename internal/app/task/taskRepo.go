package task

import "gitgub.com/diploma-mppr/backend_mppr/internal/app/models"

type Repository interface {
	GetData(id int) (*models.DataDb, error)
	SetData(DataRequest *models.DataDb) (*models.DataDb, error)
}
