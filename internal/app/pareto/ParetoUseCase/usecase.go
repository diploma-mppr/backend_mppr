package ParetoUseCase

import (
	"encoding/json"
	"fmt"
	"gitgub.com/diploma-mppr/backend_mppr/internal/app/models"
	"gitgub.com/diploma-mppr/backend_mppr/internal/app/pareto"
)

type UseCasePareto struct {
	Repository pareto.Repository
}

func NewUseCasePareto(Repository pareto.Repository) *UseCasePareto {
	return &UseCasePareto{
		Repository: Repository,
	}
}

func (u *UseCasePareto) GetPareto(id int, UserId int) (*models.ParetoJson, error) {
	Data, err := u.Repository.GetPareto(id, UserId)
	if err != nil {
		fmt.Println("UseCasePareto GetPareto", err)
		return nil, err
	}

	var Pareto models.Pareto
	err = json.Unmarshal(Data.Data, &Pareto)
	if err != nil {
		fmt.Println("UseCasePareto GetPareto", err)
		return nil, err
	}

	DataResponse := &models.ParetoJson{
		Id:   Data.Id,
		Name: Data.Name,
		Var1: Pareto.Var1,
		Var2: Pareto.Var2,
		Var3: Pareto.Var3,
	}

	return DataResponse, nil
}

func (u *UseCasePareto) SetPareto(DataRequest *models.ParetoJson, UserId int) (*models.ParetoJson, error) {
	Pareto, err := json.Marshal(
		models.Pareto{
			Var1: DataRequest.Var1,
			Var2: DataRequest.Var2,
			Var3: DataRequest.Var3,
		},
	)
	if err != nil {
		fmt.Println("UseCasePareto SetPareto", err)
		return nil, err
	}

	Data, err := u.Repository.SetPareto(&models.ParetoDb{
		UserId: UserId,
		Name:   DataRequest.Name,
		Data:   Pareto,
	})
	if err != nil {
		fmt.Println("UseCasePareto SetPareto", err)
		return nil, err
	}

	var Pareto1 models.Pareto
	err = json.Unmarshal(Data.Data, &Pareto1)
	if err != nil {
		fmt.Println("UseCasePareto SetPareto", err)
		return nil, err
	}

	DataResponse := &models.ParetoJson{
		Id:   Data.Id,
		Name: Data.Name,
		Var1: Pareto1.Var1,
		Var2: Pareto1.Var2,
		Var3: Pareto1.Var3,
	}

	return DataResponse, nil
}

func (u *UseCasePareto) UpdatePareto(DataRequest *models.ParetoJson, UserId int) (*models.ParetoJson, error) {
	Pareto, err := json.Marshal(
		models.Pareto{
			Var1: DataRequest.Var1,
			Var2: DataRequest.Var2,
			Var3: DataRequest.Var3,
		},
	)
	if err != nil {
		fmt.Println("UseCasePareto UpdatePareto", err)
		return nil, err
	}

	Data, err := u.Repository.UpdatePareto(&models.ParetoDb{
		Id:     DataRequest.Id,
		UserId: UserId,
		Data:   Pareto,
		Name:   DataRequest.Name,
	})
	if err != nil {
		fmt.Println("UseCasePareto UpdatePareto", err)
		return nil, err
	}

	var Pareto1 models.Pareto
	err = json.Unmarshal(Data.Data, &Pareto1)
	if err != nil {
		fmt.Println("UseCasePareto UpdatePareto", err)
		return nil, err
	}

	DataResponse := &models.ParetoJson{
		Id:   Data.Id,
		Name: Data.Name,
		Var1: Pareto1.Var1,
		Var2: Pareto1.Var2,
		Var3: Pareto1.Var3,
	}

	return DataResponse, nil
}

func (u *UseCasePareto) DeletePareto(DataRequest *models.ParetoJson, UserId int) error {
	err := u.Repository.DeletePareto(&models.ParetoDb{
		Id:     DataRequest.Id,
		UserId: UserId,
	})
	if err != nil {
		fmt.Println("UseCasePareto DeletePareto", err)
		return err
	}
	return nil
}
