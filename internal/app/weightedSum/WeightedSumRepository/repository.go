package WeightedSumRepository

import (
	"fmt"
	"gitgub.com/diploma-mppr/backend_mppr/internal/app/models"
	"github.com/jackc/pgx"
)

type RepositoryWeightedSum struct {
	DB *pgx.ConnPool
}

func NewRepositoryWeightedSum(db *pgx.ConnPool) *RepositoryWeightedSum {
	return &RepositoryWeightedSum{
		DB: db,
	}
}

func (r *RepositoryWeightedSum) GetWeightedSum(id int) (*models.WeightedSumDb, error) {
	DataResponse := &models.WeightedSumDb{}
	sql := `select id, name, data from "tdata" where id=$1`
	err := r.DB.QueryRow(sql, id).Scan(
		&DataResponse.Id,
		&DataResponse.Name,
		&DataResponse.Data,
	)
	if err != nil {
		fmt.Println("RepositoryWeightedSum GetWeightedSum", err)
		return nil, err
	}
	return DataResponse, nil
}

func (r *RepositoryWeightedSum) SetWeightedSum(DataRequest *models.WeightedSumDb) (*models.WeightedSumDb, error) {
	DataResponse := &models.WeightedSumDb{}
	sql := `insert into "tdata" (name, data) values ($1, $2) returning id, name, data;`
	err := r.DB.QueryRow(sql, DataRequest.Name, DataRequest.Data).Scan(
		&DataResponse.Id,
		&DataResponse.Name,
		&DataResponse.Data,
	)
	if err != nil {
		fmt.Println("RepositoryWeightedSum SetWeightedSum", err)
		return nil, err
	}

	return DataResponse, nil
}

func (r *RepositoryWeightedSum) UpdateWeightedSum(DataRequest *models.WeightedSumDb) (*models.WeightedSumDb, error) {
	DataResponse := &models.WeightedSumDb{}
	sql := `UPDATE "tdata" SET "data" = $1 WHERE "id"=$2 returning id, name, data;`
	err := r.DB.QueryRow(sql, DataRequest.Data, DataRequest.Id).Scan(
		&DataResponse.Id,
		&DataResponse.Name,
		&DataResponse.Data,
	)
	if err != nil {
		fmt.Println("RepositoryWeightedSum UpdateWeightedSum", err)
		return nil, err
	}

	return DataResponse, nil
}

func (r *RepositoryWeightedSum) DeleteWeightedSum(DataRequest *models.WeightedSumDb) error {
	sql := `DELETE FROM "tdata" WHERE "id"=$1;`
	_, err := r.DB.Query(sql, DataRequest.Id)
	if err != nil {
		fmt.Println("RepositoryWeightedSum DeleteWeightedSum", err)
		return err
	}

	return nil
}
