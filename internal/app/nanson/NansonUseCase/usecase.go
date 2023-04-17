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

func (u *UseCaseNanson) GetNanson(id int) (*models.NansonJson, error) {
	Data, err := u.Repository.GetNanson(id)
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
		Var2: Nanson.Var2,
		Var3: Nanson.Var3,
		Var4: Nanson.Var4,
	}

	return DataResponse, nil
}

func (u *UseCaseNanson) SetNanson(DataRequest *models.NansonJson) (*models.NansonJson, error) {
	Nanson, err := json.Marshal(
		models.Nanson{
			Var1: DataRequest.Var1,
			Var2: DataRequest.Var2,
			Var3: DataRequest.Var3,
			Var4: DataRequest.Var4,
		},
	)
	if err != nil {
		fmt.Println("UseCaseNanson SetNanson", err)
		return nil, err
	}

	Data, err := u.Repository.SetNanson(&models.NansonDb{
		Name: DataRequest.Name,
		Data: Nanson,
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
		Var2: Nanson1.Var2,
		Var3: Nanson1.Var3,
		Var4: Nanson1.Var4,
	}

	return DataResponse, nil
}

func (u *UseCaseNanson) UpdateNanson(DataRequest *models.NansonJson) (*models.NansonJson, error) {
	Nanson, err := json.Marshal(
		models.Nanson{
			Var1: DataRequest.Var1,
			Var2: DataRequest.Var2,
			Var3: DataRequest.Var3,
			Var4: DataRequest.Var4,
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
		Var2: Nanson1.Var2,
		Var3: Nanson1.Var3,
		Var4: Nanson1.Var4,
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
