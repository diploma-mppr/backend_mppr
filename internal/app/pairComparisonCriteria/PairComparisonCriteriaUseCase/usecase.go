package PairComparisonCriteriaUseCase

import (
	"encoding/json"
	"fmt"
	"gitgub.com/diploma-mppr/backend_mppr/internal/app/models"
	"gitgub.com/diploma-mppr/backend_mppr/internal/app/pairComparisonCriteria"
)

type UseCasePairComparisonCriteria struct {
	Repository pairComparisonCriteria.Repository
}

func NewUseCasePairComparisonCriteria(Repository pairComparisonCriteria.Repository) *UseCasePairComparisonCriteria {
	return &UseCasePairComparisonCriteria{
		Repository: Repository,
	}
}

func (u *UseCasePairComparisonCriteria) GetPairComparisonCriteria(id int) (*models.PairComparisonCriteriaJson, error) {
	Data, err := u.Repository.GetPairComparisonCriteria(id)
	if err != nil {
		fmt.Println("UseCasePairComparisonCriteria GetPairComparisonCriteria", err)
		return nil, err
	}

	var PairComparisonCriteria models.PairComparisonCriteria
	err = json.Unmarshal(Data.Data, &PairComparisonCriteria)
	if err != nil {
		fmt.Println("UseCasePairComparisonCriteria GetPairComparisonCriteria", err)
		return nil, err
	}

	DataResponse := &models.PairComparisonCriteriaJson{
		Id:    Data.Id,
		Name:  Data.Name,
		Var1:  PairComparisonCriteria.Var1,
		Var2:  PairComparisonCriteria.Var2,
		Var3:  PairComparisonCriteria.Var3,
		Var4:  PairComparisonCriteria.Var4,
		Var5:  PairComparisonCriteria.Var5,
		Var6:  PairComparisonCriteria.Var6,
		Var7:  PairComparisonCriteria.Var7,
		Var8:  PairComparisonCriteria.Var8,
		Var9:  PairComparisonCriteria.Var9,
		Var10: PairComparisonCriteria.Var10,
	}

	return DataResponse, nil
}

func (u *UseCasePairComparisonCriteria) SetPairComparisonCriteria(DataRequest *models.PairComparisonCriteriaJson) (*models.PairComparisonCriteriaJson, error) {
	PairComparisonCriteria, err := json.Marshal(
		models.PairComparisonCriteria{
			Var1:  DataRequest.Var1,
			Var2:  DataRequest.Var2,
			Var3:  DataRequest.Var3,
			Var4:  DataRequest.Var4,
			Var5:  DataRequest.Var5,
			Var6:  DataRequest.Var6,
			Var7:  DataRequest.Var7,
			Var8:  DataRequest.Var8,
			Var9:  DataRequest.Var9,
			Var10: DataRequest.Var10,
		},
	)
	if err != nil {
		fmt.Println("UseCasePairComparisonCriteria SetPairComparisonCriteria", err)
		return nil, err
	}

	Data, err := u.Repository.SetPairComparisonCriteria(&models.PairComparisonCriteriaDb{
		Name: DataRequest.Name,
		Data: PairComparisonCriteria,
	})
	if err != nil {
		fmt.Println("UseCasePairComparisonCriteria SetPairComparisonCriteria", err)
		return nil, err
	}

	var PairComparisonCriteria1 models.PairComparisonCriteria
	err = json.Unmarshal(Data.Data, &PairComparisonCriteria1)
	if err != nil {
		fmt.Println("UseCasePairComparisonCriteria SetPairComparisonCriteria", err)
		return nil, err
	}

	DataResponse := &models.PairComparisonCriteriaJson{
		Id:    Data.Id,
		Name:  Data.Name,
		Var1:  PairComparisonCriteria1.Var1,
		Var2:  PairComparisonCriteria1.Var2,
		Var3:  PairComparisonCriteria1.Var3,
		Var4:  PairComparisonCriteria1.Var4,
		Var5:  PairComparisonCriteria1.Var5,
		Var6:  PairComparisonCriteria1.Var6,
		Var7:  PairComparisonCriteria1.Var7,
		Var8:  PairComparisonCriteria1.Var8,
		Var9:  PairComparisonCriteria1.Var9,
		Var10: PairComparisonCriteria1.Var10,
	}

	return DataResponse, nil
}

func (u *UseCasePairComparisonCriteria) UpdatePairComparisonCriteria(DataRequest *models.PairComparisonCriteriaJson) (*models.PairComparisonCriteriaJson, error) {
	PairComparisonCriteria, err := json.Marshal(
		models.PairComparisonCriteria{
			Var1:  DataRequest.Var1,
			Var2:  DataRequest.Var2,
			Var3:  DataRequest.Var3,
			Var4:  DataRequest.Var4,
			Var5:  DataRequest.Var5,
			Var6:  DataRequest.Var6,
			Var7:  DataRequest.Var7,
			Var8:  DataRequest.Var8,
			Var9:  DataRequest.Var9,
			Var10: DataRequest.Var10,
		},
	)
	if err != nil {
		fmt.Println("UseCasePairComparisonCriteria UpdatePairComparisonCriteria", err)
		return nil, err
	}

	Data, err := u.Repository.UpdatePairComparisonCriteria(&models.PairComparisonCriteriaDb{
		Id:   DataRequest.Id,
		Data: PairComparisonCriteria,
	})
	if err != nil {
		fmt.Println("UseCasePairComparisonCriteria UpdatePairComparisonCriteria", err)
		return nil, err
	}

	var PairComparisonCriteria1 models.PairComparisonCriteria
	err = json.Unmarshal(Data.Data, &PairComparisonCriteria1)
	if err != nil {
		fmt.Println("UseCasePairComparisonCriteria UpdatePairComparisonCriteria", err)
		return nil, err
	}

	DataResponse := &models.PairComparisonCriteriaJson{
		Id:    Data.Id,
		Name:  Data.Name,
		Var1:  PairComparisonCriteria1.Var1,
		Var2:  PairComparisonCriteria1.Var2,
		Var3:  PairComparisonCriteria1.Var3,
		Var4:  PairComparisonCriteria1.Var4,
		Var5:  PairComparisonCriteria1.Var5,
		Var6:  PairComparisonCriteria1.Var6,
		Var7:  PairComparisonCriteria1.Var7,
		Var8:  PairComparisonCriteria1.Var8,
		Var9:  PairComparisonCriteria1.Var9,
		Var10: PairComparisonCriteria1.Var10,
	}

	return DataResponse, nil
}

func (u *UseCasePairComparisonCriteria) DeletePairComparisonCriteria(DataRequest *models.PairComparisonCriteriaJson) error {
	err := u.Repository.DeletePairComparisonCriteria(&models.PairComparisonCriteriaDb{
		Id: DataRequest.Id,
	})
	if err != nil {
		fmt.Println("UseCasePairComparisonCriteria DeletePairComparisonCriteria", err)
		return err
	}
	return nil
}
