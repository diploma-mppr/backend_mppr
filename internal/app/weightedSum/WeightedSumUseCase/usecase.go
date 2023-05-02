package WeightedSumUseCase

import (
	"encoding/json"
	"fmt"
	"gitgub.com/diploma-mppr/backend_mppr/internal/app/models"
	"gitgub.com/diploma-mppr/backend_mppr/internal/app/weightedSum"
)

type UseCaseWeightedSum struct {
	Repository weightedSum.Repository
}

func NewUseCaseWeightedSum(Repository weightedSum.Repository) *UseCaseWeightedSum {
	return &UseCaseWeightedSum{
		Repository: Repository,
	}
}

func (u *UseCaseWeightedSum) GetWeightedSum(id int, UserId int) (*models.WeightedSumJson, error) {
	Data, err := u.Repository.GetWeightedSum(id, UserId)
	if err != nil {
		fmt.Println("UseCaseWeightedSum GetWeightedSum", err)
		return nil, err
	}

	var WeightedSum models.WeightedSum
	err = json.Unmarshal(Data.Data, &WeightedSum)
	if err != nil {
		fmt.Println("UseCaseWeightedSum GetWeightedSum", err)
		return nil, err
	}

	DataResponse := &models.WeightedSumJson{
		Id:   Data.Id,
		Name: Data.Name,
		Var1: WeightedSum.Var1,
		Var2: WeightedSum.Var2,
		Var3: WeightedSum.Var3,
		Var4: WeightedSum.Var4,
		Var5: WeightedSum.Var5,
	}

	return DataResponse, nil
}

func (u *UseCaseWeightedSum) SetWeightedSum(DataRequest *models.WeightedSumJson, UserId int) (*models.WeightedSumJson, error) {
	WeightedSum, err := json.Marshal(
		models.WeightedSum{
			Var1: DataRequest.Var1,
			Var2: DataRequest.Var2,
			Var3: DataRequest.Var3,
			Var4: DataRequest.Var4,
			Var5: DataRequest.Var5,
		},
	)
	if err != nil {
		fmt.Println("UseCaseWeightedSum SetWeightedSum", err)
		return nil, err
	}

	Data, err := u.Repository.SetWeightedSum(&models.WeightedSumDb{
		UserId: UserId,
		Name:   DataRequest.Name,
		Data:   WeightedSum,
	})
	if err != nil {
		fmt.Println("UseCaseWeightedSum SetWeightedSum", err)
		return nil, err
	}

	var WeightedSum1 models.WeightedSum
	err = json.Unmarshal(Data.Data, &WeightedSum1)
	if err != nil {
		fmt.Println("UseCaseWeightedSum SetWeightedSum", err)
		return nil, err
	}

	DataResponse := &models.WeightedSumJson{
		Id:   Data.Id,
		Name: Data.Name,
		Var1: WeightedSum1.Var1,
		Var2: WeightedSum1.Var2,
		Var3: WeightedSum1.Var3,
		Var4: WeightedSum1.Var4,
		Var5: WeightedSum1.Var5,
	}

	return DataResponse, nil
}

func (u *UseCaseWeightedSum) UpdateWeightedSum(DataRequest *models.WeightedSumJson) (*models.WeightedSumJson, error) {
	WeightedSum, err := json.Marshal(
		models.WeightedSum{
			Var1: DataRequest.Var1,
			Var2: DataRequest.Var2,
			Var3: DataRequest.Var3,
			Var4: DataRequest.Var4,
			Var5: DataRequest.Var5,
		},
	)
	if err != nil {
		fmt.Println("UseCaseWeightedSum UpdateWeightedSum", err)
		return nil, err
	}

	Data, err := u.Repository.UpdateWeightedSum(&models.WeightedSumDb{
		Id:   DataRequest.Id,
		Data: WeightedSum,
	})
	if err != nil {
		fmt.Println("UseCaseWeightedSum UpdateWeightedSum", err)
		return nil, err
	}

	var WeightedSum1 models.WeightedSum
	err = json.Unmarshal(Data.Data, &WeightedSum1)
	if err != nil {
		fmt.Println("UseCaseWeightedSum UpdateWeightedSum", err)
		return nil, err
	}

	DataResponse := &models.WeightedSumJson{
		Id:   Data.Id,
		Name: Data.Name,
		Var1: WeightedSum1.Var1,
		Var2: WeightedSum1.Var2,
		Var3: WeightedSum1.Var3,
		Var4: WeightedSum1.Var4,
		Var5: WeightedSum1.Var5,
	}

	return DataResponse, nil
}

func (u *UseCaseWeightedSum) DeleteWeightedSum(DataRequest *models.WeightedSumJson) error {
	err := u.Repository.DeleteWeightedSum(&models.WeightedSumDb{
		Id: DataRequest.Id,
	})
	if err != nil {
		fmt.Println("UseCaseWeightedSum DeleteWeightedSum", err)
		return err
	}
	return nil
}
