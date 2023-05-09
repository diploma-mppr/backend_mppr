package BordaUseCase

import (
	"encoding/json"
	"fmt"
	"gitgub.com/diploma-mppr/backend_mppr/internal/app/borda"
	"gitgub.com/diploma-mppr/backend_mppr/internal/app/models"
)

type UseCaseBorda struct {
	Repository borda.Repository
}

func NewUseCaseBorda(Repository borda.Repository) *UseCaseBorda {
	return &UseCaseBorda{
		Repository: Repository,
	}
}

func (u *UseCaseBorda) GetBorda(id int, UserId int) (*models.BordaJson, error) {
	Data, err := u.Repository.GetBorda(id, UserId)
	if err != nil {
		fmt.Println("UseCaseBorda GetBorda", err)
		return nil, err
	}

	var Borda models.Borda
	err = json.Unmarshal(Data.Data, &Borda)
	if err != nil {
		fmt.Println("UseCaseBorda GetBorda", err)
		return nil, err
	}

	DataResponse := &models.BordaJson{
		Id:   Data.Id,
		Name: Data.Name,
		Var1: Borda.Var1,
	}

	return DataResponse, nil
}

func (u *UseCaseBorda) SetBorda(DataRequest *models.BordaJson, UserId int) (*models.BordaJson, error) {
	Borda, err := json.Marshal(
		models.Borda{
			Var1: DataRequest.Var1,
		},
	)
	if err != nil {
		fmt.Println("UseCaseBorda SetBorda", err)
		return nil, err
	}

	Data, err := u.Repository.SetBorda(&models.BordaDb{
		UserId: UserId,
		Name:   DataRequest.Name,
		Data:   Borda,
	})
	if err != nil {
		fmt.Println("UseCaseBorda SetBorda", err)
		return nil, err
	}

	var Borda1 models.Borda
	err = json.Unmarshal(Data.Data, &Borda1)
	if err != nil {
		fmt.Println("UseCaseBorda SetBorda", err)
		return nil, err
	}

	DataResponse := &models.BordaJson{
		Id:   Data.Id,
		Name: Data.Name,
		Var1: Borda1.Var1,
	}

	return DataResponse, nil
}

func (u *UseCaseBorda) UpdateBorda(DataRequest *models.BordaJson) (*models.BordaJson, error) {
	Borda, err := json.Marshal(
		models.Borda{
			Var1: DataRequest.Var1,
		},
	)
	if err != nil {
		fmt.Println("UseCaseBorda UpdateBorda", err)
		return nil, err
	}

	Data, err := u.Repository.UpdateBorda(&models.BordaDb{
		Id:   DataRequest.Id,
		Data: Borda,
		Name: DataRequest.Name,
	})
	if err != nil {
		fmt.Println("UseCaseBorda UpdateBorda", err)
		return nil, err
	}

	var Borda1 models.Borda
	err = json.Unmarshal(Data.Data, &Borda1)
	if err != nil {
		fmt.Println("UseCaseBorda UpdateBorda", err)
		return nil, err
	}

	DataResponse := &models.BordaJson{
		Id:   Data.Id,
		Name: Data.Name,
		Var1: Borda1.Var1,
	}

	return DataResponse, nil
}

func (u *UseCaseBorda) DeleteBorda(DataRequest *models.BordaJson) error {
	err := u.Repository.DeleteBorda(&models.BordaDb{
		Id: DataRequest.Id,
	})
	if err != nil {
		fmt.Println("UseCaseBorda DeleteBorda", err)
		return err
	}
	return nil
}
