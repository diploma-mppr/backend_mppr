package task

import (
	"encoding/json"
	"fmt"
	"gitgub.com/diploma-mppr/backend_mppr/internal/app/models"
)

type UcaseTask struct {
	Repo Repository
}

func NewUcaseTask(TaskRepo Repository) *UcaseTask {
	return &UcaseTask{
		Repo: TaskRepo,
	}
}

//func (u *UcaseTask) GetTask() (Task models.DataJson, err error) {
//	Task, err = u.Repo.GetTask()
//
//	if err != nil {
//		return
//	}
//	return
//}

func (u *UcaseTask) SetData(TaskRequest models.DataJson) (*models.DataJson, error) {
	fmt.Println("usecase", TaskRequest)
	var1, err := json.Marshal(models.Float64Json{MyFloat: TaskRequest.Var1})
	fmt.Println("usecase", var1, err)
	if err != nil {
		return nil, err
	}
	var2, err := json.Marshal(models.Float64Json{MyFloat: TaskRequest.Var2})
	fmt.Println("usecase", var2, err)
	if err != nil {
		return nil, err
	}
	var3, err := json.Marshal(models.Float64Json{MyFloat: TaskRequest.Var3})
	fmt.Println("usecase", var3, err)
	if err != nil {
		return nil, err
	}

	Task, err := u.Repo.SetData(models.DataDb{
		Name: TaskRequest.Name,
		Var1: var1,
		Var2: var2,
		Var3: var3,
	})
	if err != nil {
		return nil, err
	}
	fmt.Println("usecase", Task)

	var var11 models.Float64Json
	err = json.Unmarshal(Task.Var1, &var11)
	if err != nil {
		return nil, err
	}
	var var21 models.Float64Json
	err = json.Unmarshal(Task.Var2, &var21)
	if err != nil {
		return nil, err
	}
	var var31 models.Float64Json
	err = json.Unmarshal(Task.Var3, &var31)
	if err != nil {
		return nil, err
	}

	TaskResponse := &models.DataJson{
		Id:   Task.Id,
		Name: Task.Name,
		Var1: var11.MyFloat,
		Var2: var21.MyFloat,
		Var3: var31.MyFloat,
	}

	fmt.Println("usecase", TaskResponse)

	return TaskResponse, nil
}
