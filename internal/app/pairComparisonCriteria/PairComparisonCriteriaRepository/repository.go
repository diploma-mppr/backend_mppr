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

func (r *RepositoryPairComparisonCriteria) GetPairComparisonCriteria(id int, UserId int) (*models.PairComparisonCriteriaDb, error) {
	DataResponse := &models.PairComparisonCriteriaDb{}
	sql := `select "id", "name", "data" from "method" where "id" = $1 and "user_id" = $2`
	err := r.DB.QueryRow(sql, id, UserId).Scan(
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
	sql := `insert into "method" ("user_id", "name", "data", "method_name") values ($1, $2, $3, 'pairComparisonCriteria') returning id, name, data;`
	err := r.DB.QueryRow(sql, DataRequest.UserId, DataRequest.Name, DataRequest.Data).Scan(
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
	sql := `UPDATE "method" SET "data" = $1 WHERE "id"=$2 returning id, name, data;`
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
	sql := `DELETE FROM "method" WHERE "id"=$1;`
	_, err := r.DB.Query(sql, DataRequest.Id)
	if err != nil {
		fmt.Println("RepositoryPairComparisonCriteriaDb DeletePairComparisonCriteriaDb", err)
		return err
	}

	return nil
}
