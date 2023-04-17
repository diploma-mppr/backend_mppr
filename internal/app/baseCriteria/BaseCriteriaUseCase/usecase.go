package BaseCriteriaUseCase

import (
	"encoding/json"
	"fmt"
	"gitgub.com/diploma-mppr/backend_mppr/internal/app/baseCriteria"
	"gitgub.com/diploma-mppr/backend_mppr/internal/app/models"
)

type UseCaseBaseCriteria struct {
	Repository baseCriteria.Repository
}

func NewUseCaseBaseCriteria(Repository baseCriteria.Repository) *UseCaseBaseCriteria {
	return &UseCaseBaseCriteria{
		Repository: Repository,
	}
}

func (u *UseCaseBaseCriteria) GetBaseCriteria(id int) (*models.BaseCriteriaJson, error) {
	Data, err := u.Repository.GetBaseCriteria(id)
	if err != nil {
		fmt.Println("UseCasePareto GetPareto", err)
		return nil, err
	}

	var BasicCriteria models.BaseCriteria
	err = json.Unmarshal(Data.Data, &BasicCriteria)
	if err != nil {
		fmt.Println("UseCasePareto GetPareto", err)
		return nil, err
	}

	DataResponse := &models.BaseCriteriaJson{
		Id:   Data.Id,
		Name: Data.Name,
		Var1: BasicCriteria.Var1,
	}

	return DataResponse, nil
}

func (u *UseCaseBaseCriteria) SetBaseCriteria(DataRequest *models.BaseCriteriaJson) (*models.BaseCriteriaJson, error) {
	BasicCriteria, err := json.Marshal(
		models.BaseCriteria{
			Var1: DataRequest.Var1,
		},
	)
	if err != nil {
		fmt.Println("UseCasePareto SetPareto", err)
		return nil, err
	}

	Data, err := u.Repository.SetBaseCriteria(&models.BaseCriteriaDb{
		Name: DataRequest.Name,
		Data: BasicCriteria,
	})
	if err != nil {
		fmt.Println("UseCasePareto SetPareto", err)
		return nil, err
	}

	var BasicCriteria1 models.BaseCriteria
	err = json.Unmarshal(Data.Data, &BasicCriteria1)
	if err != nil {
		fmt.Println("UseCasePareto SetPareto", err)
		return nil, err
	}

	DataResponse := &models.BaseCriteriaJson{
		Id:   Data.Id,
		Name: Data.Name,
		Var1: BasicCriteria1.Var1,
	}

	return DataResponse, nil
}

func (u *UseCaseBaseCriteria) UpdateBaseCriteria(DataRequest *models.BaseCriteriaJson) (*models.BaseCriteriaJson, error) {
	BasicCriteria, err := json.Marshal(
		models.BaseCriteria{
			Var1: DataRequest.Var1,
		},
	)
	if err != nil {
		fmt.Println("UseCasePareto UpdatePareto", err)
		return nil, err
	}

	Data, err := u.Repository.UpdateBaseCriteria(&models.BaseCriteriaDb{
		Id:   DataRequest.Id,
		Data: BasicCriteria,
	})
	if err != nil {
		fmt.Println("UseCasePareto UpdatePareto", err)
		return nil, err
	}

	var BasicCriteria1 models.BaseCriteria
	err = json.Unmarshal(Data.Data, &BasicCriteria1)
	if err != nil {
		fmt.Println("UseCasePareto UpdatePareto", err)
		return nil, err
	}

	DataResponse := &models.BaseCriteriaJson{
		Id:   Data.Id,
		Name: Data.Name,
		Var1: BasicCriteria1.Var1,
	}

	return DataResponse, nil
}

func (u *UseCaseBaseCriteria) DeleteBaseCriteria(DataRequest *models.BaseCriteriaJson) error {
	err := u.Repository.DeleteBaseCriteria(&models.BaseCriteriaDb{
		Id: DataRequest.Id,
	})
	if err != nil {
		fmt.Println("UseCasePareto DeletePareto", err)
		return err
	}
	return nil
}
