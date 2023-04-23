package DataUseCase

import (
	"fmt"
	"gitgub.com/diploma-mppr/backend_mppr/internal/app/data"
	"gitgub.com/diploma-mppr/backend_mppr/internal/app/models"
)

type UseCaseData struct {
	Repository data.Repository
}

func NewUseCaseData(Repository data.Repository) *UseCaseData {
	return &UseCaseData{
		Repository: Repository,
	}
}

func (u *UseCaseData) GetAll(UserId int) (*[]models.DataDb, error) {
	Data, err := u.Repository.GetAll(UserId)
	if err != nil {
		fmt.Println("UseCaseData GetAll", err)
		return nil, err
	}

	var DataResponse []models.DataDb

	for _, data := range *Data {
		DataResponse = append(
			DataResponse,
			models.DataDb{
				Id:         data.Id,
				MethodName: data.MethodName,
				Name:       data.Name,
			},
		)
	}

	return &DataResponse, err
}
