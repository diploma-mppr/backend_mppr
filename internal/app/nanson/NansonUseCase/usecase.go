package NansonUseCase

import (
	"encoding/json"
	"fmt"
	"gitgub.com/diploma-mppr/backend_mppr/internal/app/models"
	"gitgub.com/diploma-mppr/backend_mppr/internal/app/nanson"
)

type UseCaseNanson struct {
	Repository nanson.Repository
}

func NewUseCaseNanson(Repository nanson.Repository) *UseCaseNanson {
	return &UseCaseNanson{
		Repository: Repository,
	}
}

func (u *UseCaseNanson) GetNanson(id int, UserId int) (*models.NansonJson, error) {
	Data, err := u.Repository.GetNanson(id, UserId)
	if err != nil {
		fmt.Println("UseCaseNanson GetNanson", err)
		return nil, err
	}

	var Nanson models.Nanson
	err = json.Unmarshal(Data.Data, &Nanson)
	if err != nil {
		fmt.Println("UseCaseNanson GetNanson", err)
		return nil, err
	}

	DataResponse := &models.NansonJson{
		Id:   Data.Id,
		Name: Data.Name,
		Var1: Nanson.Var1,
	}

	return DataResponse, nil
}

func (u *UseCaseNanson) SetNanson(DataRequest *models.NansonJson, UserId int) (*models.NansonJson, error) {
	Nanson, err := json.Marshal(
		models.Nanson{
			Var1: DataRequest.Var1,
		},
	)
	if err != nil {
		fmt.Println("UseCaseNanson SetNanson", err)
		return nil, err
	}

	Data, err := u.Repository.SetNanson(&models.NansonDb{
		UserId: UserId,
		Name:   DataRequest.Name,
		Data:   Nanson,
	})
	if err != nil {
		fmt.Println("UseCaseNanson SetNanson", err)
		return nil, err
	}

	var Nanson1 models.Nanson
	err = json.Unmarshal(Data.Data, &Nanson1)
	if err != nil {
		fmt.Println("UseCaseNanson SetNanson", err)
		return nil, err
	}

	DataResponse := &models.NansonJson{
		Id:   Data.Id,
		Name: Data.Name,
		Var1: Nanson1.Var1,
	}

	return DataResponse, nil
}

func (u *UseCaseNanson) UpdateNanson(DataRequest *models.NansonJson) (*models.NansonJson, error) {
	Nanson, err := json.Marshal(
		models.Nanson{
			Var1: DataRequest.Var1,
		},
	)
	if err != nil {
		fmt.Println("UseCaseNanson UpdateNanson", err)
		return nil, err
	}

	Data, err := u.Repository.UpdateNanson(&models.NansonDb{
		Id:   DataRequest.Id,
		Data: Nanson,
	})
	if err != nil {
		fmt.Println("UseCaseNanson UpdateNanson", err)
		return nil, err
	}

	var Nanson1 models.Nanson
	err = json.Unmarshal(Data.Data, &Nanson1)
	if err != nil {
		fmt.Println("UseCaseNanson UpdateNanson", err)
		return nil, err
	}

	DataResponse := &models.NansonJson{
		Id:   Data.Id,
		Name: Data.Name,
		Var1: Nanson1.Var1,
	}

	return DataResponse, nil
}

func (u *UseCaseNanson) DeleteNanson(DataRequest *models.NansonJson) error {
	err := u.Repository.DeleteNanson(&models.NansonDb{
		Id: DataRequest.Id,
	})
	if err != nil {
		fmt.Println("UseCaseNanson DeleteNanson", err)
		return err
	}
	return nil
}
