package PointScoreRepository

import (
	"fmt"
	"gitgub.com/diploma-mppr/backend_mppr/internal/app/models"
	"github.com/jackc/pgx"
)

type RepositoryPointScore struct {
	DB *pgx.ConnPool
}

func NewRepositoryPointScore(db *pgx.ConnPool) *RepositoryPointScore {
	return &RepositoryPointScore{
		DB: db,
	}
}

func (r *RepositoryPointScore) GetPointScore(id int) (*models.PointScoreDb, error) {
	DataResponse := &models.PointScoreDb{}
	sql := `select id, name, data from "tdata" where id=$1`
	err := r.DB.QueryRow(sql, id).Scan(
		&DataResponse.Id,
		&DataResponse.Name,
		&DataResponse.Data,
	)
	if err != nil {
		fmt.Println("RepositoryPareto GetPointScore", err)
		return nil, err
	}
	return DataResponse, nil
}

func (r *RepositoryPointScore) SetPointScore(DataRequest *models.PointScoreDb) (*models.PointScoreDb, error) {
	DataResponse := &models.PointScoreDb{}
	sql := `insert into "tdata" (name, data) values ($1, $2) returning id, name, data;`
	err := r.DB.QueryRow(sql, DataRequest.Name, DataRequest.Data).Scan(
		&DataResponse.Id,
		&DataResponse.Name,
		&DataResponse.Data,
	)
	if err != nil {
		fmt.Println("RepositoryPareto SetPointScore", err)
		return nil, err
	}

	return DataResponse, nil
}

func (r *RepositoryPointScore) UpdatePointScore(DataRequest *models.PointScoreDb) (*models.PointScoreDb, error) {
	DataResponse := &models.PointScoreDb{}
	sql := `UPDATE "tdata" SET "data" = $1 WHERE "id"=$2 returning id, name, data;`
	err := r.DB.QueryRow(sql, DataRequest.Data, DataRequest.Id).Scan(
		&DataResponse.Id,
		&DataResponse.Name,
		&DataResponse.Data,
	)
	if err != nil {
		fmt.Println("RepositoryPareto UpdatePointScore", err)
		return nil, err
	}

	return DataResponse, nil
}

func (r *RepositoryPointScore) DeletePointScore(DataRequest *models.PointScoreDb) error {
	sql := `DELETE FROM "tdata" WHERE "id"=$1;`
	_, err := r.DB.Query(sql, DataRequest.Id)
	if err != nil {
		fmt.Println("RepositoryPareto DeletePointScore", err)
		return err
	}

	return nil
}
