package task

import (
	"encoding/json"
	"github.com/tp-study-ai/backend/internal/app/models"
)

type UseCaseTask struct {
	Repo Repository
}

func NewUseCaseTask(TaskRepo Repository) *UseCaseTask {
	return &UseCaseTask{
		Repo: TaskRepo,
	}
}

func (u *UseCaseTask) GetData(id int) (*models.DataJson, error) {
	Data, err := u.Repo.GetData(id)
	if err != nil {
		return nil, err
	}

	var var11 models.Float64Json
	err = json.Unmarshal(Data.Var1, &var11)
	if err != nil {
		return nil, err
	}
	var var21 models.Float64Json
	err = json.Unmarshal(Data.Var2, &var21)
	if err != nil {
		return nil, err
	}
	var var31 models.Float64Json
	err = json.Unmarshal(Data.Var3, &var31)
	if err != nil {
		return nil, err
	}

	DataResponse := &models.DataJson{
		Id:   Data.Id,
		Name: Data.Name,
		Var1: var11.MyFloat,
		Var2: var21.MyFloat,
		Var3: var31.MyFloat,
	}

	return DataResponse, nil
}

func (u *UseCaseTask) SetData(DataRequest *models.DataJson) (*models.DataJson, error) {
	var1, err := json.Marshal(models.Float64Json{MyFloat: DataRequest.Var1})
	if err != nil {
		return nil, err
	}
	var2, err := json.Marshal(models.Float64Json{MyFloat: DataRequest.Var2})
	if err != nil {
		return nil, err
	}
	var3, err := json.Marshal(models.Float64Json{MyFloat: DataRequest.Var3})
	if err != nil {
		return nil, err
	}

	Data, err := u.Repo.SetData(&models.DataDb{
		Name: DataRequest.Name,
		Var1: var1,
		Var2: var2,
		Var3: var3,
	})
	if err != nil {
		return nil, err
	}

	var var11 models.Float64Json
	err = json.Unmarshal(Data.Var1, &var11)
	if err != nil {
		return nil, err
	}
	var var21 models.Float64Json
	err = json.Unmarshal(Data.Var2, &var21)
	if err != nil {
		return nil, err
	}
	var var31 models.Float64Json
	err = json.Unmarshal(Data.Var3, &var31)
	if err != nil {
		return nil, err
	}

	DataResponse := &models.DataJson{
		Id:   Data.Id,
		Name: Data.Name,
		Var1: var11.MyFloat,
		Var2: var21.MyFloat,
		Var3: var31.MyFloat,
	}

	return DataResponse, nil
}
