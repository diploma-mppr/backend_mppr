package task

import "gitgub.com/diploma-mppr/backend_mppr/internal/app/models"

type Ucase interface {
	SetData(TaskRequest models.DataJson) (*models.DataJson, error)
}
