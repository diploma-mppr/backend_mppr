package PairComparisonCriteriaRepository

import (
	"fmt"
	"gitgub.com/diploma-mppr/backend_mppr/internal/app/models"
	"github.com/jackc/pgx"
)

type RepositoryPairComparisonCriteria struct {
	DB *pgx.ConnPool
}

func NewRepositoryPairComparisonCriteria(db *pgx.ConnPool) *RepositoryPairComparisonCriteria {
	return &RepositoryPairComparisonCriteria{
		DB: db,
	}
}

func (r *RepositoryPairComparisonCriteria) GetPairComparisonCriteria(id int) (*models.PairComparisonCriteriaDb, error) {
	DataResponse := &models.PairComparisonCriteriaDb{}
	sql := `select id, name, data from "tdata" where id=$1`
	err := r.DB.QueryRow(sql, id).Scan(
		&DataResponse.Id,
		&DataResponse.Name,
		&DataResponse.Data,
	)
	if err != nil {
		fmt.Println("RepositoryPairComparisonCriteriaDb GetPairComparisonCriteriaDb", err)
		return nil, err
	}
	return DataResponse, nil
}

func (r *RepositoryPairComparisonCriteria) SetPairComparisonCriteria(DataRequest *models.PairComparisonCriteriaDb) (*models.PairComparisonCriteriaDb, error) {
	DataResponse := &models.PairComparisonCriteriaDb{}
	sql := `insert into "tdata" (name, data) values ($1, $2) returning id, name, data;`
	err := r.DB.QueryRow(sql, DataRequest.Name, DataRequest.Data).Scan(
		&DataResponse.Id,
		&DataResponse.Name,
		&DataResponse.Data,
	)
	if err != nil {
		fmt.Println("RepositoryPairComparisonCriteriaDb SetPairComparisonCriteriaDb", err)
		return nil, err
	}

	return DataResponse, nil
}

func (r *RepositoryPairComparisonCriteria) UpdatePairComparisonCriteria(DataRequest *models.PairComparisonCriteriaDb) (*models.PairComparisonCriteriaDb, error) {
	DataResponse := &models.PairComparisonCriteriaDb{}
	sql := `UPDATE "tdata" SET "data" = $1 WHERE "id"=$2 returning id, name, data;`
	err := r.DB.QueryRow(sql, DataRequest.Data, DataRequest.Id).Scan(
		&DataResponse.Id,
		&DataResponse.Name,
		&DataResponse.Data,
	)
	if err != nil {
		fmt.Println("RepositoryPairComparisonCriteriaDb UpdatePairComparisonCriteriaDb", err)
		return nil, err
	}

	return DataResponse, nil
}

func (r *RepositoryPairComparisonCriteria) DeletePairComparisonCriteria(DataRequest *models.PairComparisonCriteriaDb) error {
	sql := `DELETE FROM "tdata" WHERE "id"=$1;`
	_, err := r.DB.Query(sql, DataRequest.Id)
	if err != nil {
		fmt.Println("RepositoryPairComparisonCriteriaDb DeletePairComparisonCriteriaDb", err)
		return err
	}

	return nil
}
