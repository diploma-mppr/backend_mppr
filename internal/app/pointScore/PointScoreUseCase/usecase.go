package PointScoreUseCase

import (
	"encoding/json"
	"fmt"
	"gitgub.com/diploma-mppr/backend_mppr/internal/app/models"
	"gitgub.com/diploma-mppr/backend_mppr/internal/app/pointScore"
)

type UseCasePointScore struct {
	Repository pointScore.Repository
}

func NewUseCasePointScore(Repository pointScore.Repository) *UseCasePointScore {
	return &UseCasePointScore{
		Repository: Repository,
	}
}

func (u *UseCasePointScore) GetPointScore(id int, UserId int) (*models.PointScoreJson, error) {
	Data, err := u.Repository.GetPointScore(id, UserId)
	if err != nil {
		fmt.Println("UseCasePareto GetPointScore", err)
		return nil, err
	}

	var Pareto models.PointScore
	err = json.Unmarshal(Data.Data, &Pareto)
	if err != nil {
		fmt.Println("UseCasePareto GetPointScore", err)
		return nil, err
	}

	DataResponse := &models.PointScoreJson{
		Id:   Data.Id,
		Name: Data.Name,
		Var1: Pareto.Var1,
	}

	return DataResponse, nil
}

func (u *UseCasePointScore) SetPointScore(DataRequest *models.PointScoreJson, UserId int) (*models.PointScoreJson, error) {
	Pareto, err := json.Marshal(
		models.PointScore{
			Var1: DataRequest.Var1,
		},
	)
	if err != nil {
		fmt.Println("UseCasePareto SetPointScore", err)
		return nil, err
	}

	Data, err := u.Repository.SetPointScore(&models.PointScoreDb{
		UserId: UserId,
		Name:   DataRequest.Name,
		Data:   Pareto,
	})
	if err != nil {
		fmt.Println("UseCasePareto SetPointScore", err)
		return nil, err
	}

	var Pareto1 models.Pareto
	err = json.Unmarshal(Data.Data, &Pareto1)
	if err != nil {
		fmt.Println("UseCasePareto SetPointScore", err)
		return nil, err
	}

	DataResponse := &models.PointScoreJson{
		Id:   Data.Id,
		Name: Data.Name,
		Var1: Pareto1.Var1,
	}

	return DataResponse, nil
}

func (u *UseCasePointScore) UpdatePointScore(DataRequest *models.PointScoreJson) (*models.PointScoreJson, error) {
	Pareto, err := json.Marshal(
		models.Pareto{
			Var1: DataRequest.Var1,
		},
	)
	if err != nil {
		fmt.Println("UseCasePareto UpdatePointScore", err)
		return nil, err
	}

	Data, err := u.Repository.UpdatePointScore(&models.PointScoreDb{
		Id:   DataRequest.Id,
		Data: Pareto,
	})
	if err != nil {
		fmt.Println("UseCasePareto UpdatePointScore", err)
		return nil, err
	}

	var Pareto1 models.Pareto
	err = json.Unmarshal(Data.Data, &Pareto1)
	if err != nil {
		fmt.Println("UseCasePareto UpdatePointScore", err)
		return nil, err
	}

	DataResponse := &models.PointScoreJson{
		Id:   Data.Id,
		Name: Data.Name,
		Var1: Pareto1.Var1,
	}

	return DataResponse, nil
}

func (u *UseCasePointScore) DeletePointScore(DataRequest *models.PointScoreJson) error {
	err := u.Repository.DeletePointScore(&models.PointScoreDb{
		Id: DataRequest.Id,
	})
	if err != nil {
		fmt.Println("UseCasePareto DeletePointScore", err)
		return err
	}
	return nil
}
