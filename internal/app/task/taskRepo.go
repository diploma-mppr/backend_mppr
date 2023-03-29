package task

import "gitgub.com/diploma-mppr/backend_mppr/internal/app/models"

type Repository interface {
	GetData() (data models.Data, err error)
	SetData(data models.DataReq) (data1 models.Data, err error)
}
