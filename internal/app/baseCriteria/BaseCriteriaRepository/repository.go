package BaseCriteriaRepository

import (
	"fmt"
	"gitgub.com/diploma-mppr/backend_mppr/internal/app/models"
	"github.com/jackc/pgx"
)

type RepositoryBaseCriteria struct {
	DB *pgx.ConnPool
}

func NewRepositoryBaseCriteria(db *pgx.ConnPool) *RepositoryBaseCriteria {
	return &RepositoryBaseCriteria{
		DB: db,
	}
}

func (r *RepositoryBaseCriteria) GetBaseCriteria(id int) (*models.BaseCriteriaDb, error) {
	DataResponse := &models.BaseCriteriaDb{}
	sql := `select id, name, data from "tdata" where id=$1`
	err := r.DB.QueryRow(sql, id).Scan(
		&DataResponse.Id,
		&DataResponse.Name,
		&DataResponse.Data,
	)
	if err != nil {
		fmt.Println("RepositoryBaseCriteria GetBasicCriteria", err)
		return nil, err
	}
	return DataResponse, nil
}

func (r *RepositoryBaseCriteria) SetBaseCriteria(DataRequest *models.BaseCriteriaDb) (*models.BaseCriteriaDb, error) {
	DataResponse := &models.BaseCriteriaDb{}
	sql := `insert into "tdata" (name, data) values ($1, $2) returning id, name, data;`
	err := r.DB.QueryRow(sql, DataRequest.Name, DataRequest.Data).Scan(
		&DataResponse.Id,
		&DataResponse.Name,
		&DataResponse.Data,
	)
	if err != nil {
		fmt.Println("RepositoryBaseCriteria SetBasicCriteria", err)
		return nil, err
	}

	return DataResponse, nil
}

func (r *RepositoryBaseCriteria) UpdateBaseCriteria(DataRequest *models.BaseCriteriaDb) (*models.BaseCriteriaDb, error) {
	DataResponse := &models.BaseCriteriaDb{}
	sql := `UPDATE "tdata" SET "data" = $1 WHERE "id"=$2 returning id, name, data;`
	err := r.DB.QueryRow(sql, DataRequest.Data, DataRequest.Id).Scan(
		&DataResponse.Id,
		&DataResponse.Name,
		&DataResponse.Data,
	)
	if err != nil {
		fmt.Println("RepositoryBaseCriteria UpdateBasicCriteria", err)
		return nil, err
	}

	return DataResponse, nil
}

func (r *RepositoryBaseCriteria) DeleteBaseCriteria(DataRequest *models.BaseCriteriaDb) error {
	sql := `DELETE FROM "tdata" WHERE "id"=$1;`
	_, err := r.DB.Query(sql, DataRequest.Id)
	if err != nil {
		fmt.Println("RepositoryBaseCriteria DeleteBasicCriteria", err)
		return err
	}

	return nil
}
