package ParetoRepository

import (
	"fmt"
	"gitgub.com/diploma-mppr/backend_mppr/internal/app/models"
	"github.com/jackc/pgx"
)

type RepositoryPareto struct {
	DB *pgx.ConnPool
}

func NewRepositoryPareto(db *pgx.ConnPool) *RepositoryPareto {
	return &RepositoryPareto{
		DB: db,
	}
}

func (r *RepositoryPareto) GetPareto(id int, UserId int) (*models.ParetoDb, error) {
	DataResponse := &models.ParetoDb{}
	sql := `select id, name, data from "method" where id=$1 and user_id = $2;`
	err := r.DB.QueryRow(sql, id, UserId).Scan(
		&DataResponse.Id,
		&DataResponse.Name,
		&DataResponse.Data,
	)
	if err != nil {
		fmt.Println("RepositoryPareto GetPareto", err)
		return nil, err
	}
	return DataResponse, nil
}

func (r *RepositoryPareto) SetPareto(DataRequest *models.ParetoDb) (*models.ParetoDb, error) {
	DataResponse := &models.ParetoDb{}
	sql := `insert into "method" (user_id, name, data, method_name) values ($1, $2, $3, 'pareto') returning id, name, data;`
	err := r.DB.QueryRow(sql, DataRequest.UserId, DataRequest.Name, DataRequest.Data).Scan(
		&DataResponse.Id,
		&DataResponse.Name,
		&DataResponse.Data,
	)
	if err != nil {
		fmt.Println("RepositoryPareto SetPareto", err)
		return nil, err
	}

	return DataResponse, nil
}

func (r *RepositoryPareto) UpdatePareto(DataRequest *models.ParetoDb) (*models.ParetoDb, error) {
	DataResponse := &models.ParetoDb{}
	sql := `UPDATE "method" SET "data" = $1 WHERE "id"=$2 returning id, name, data;`
	err := r.DB.QueryRow(sql, DataRequest.Data, DataRequest.Id).Scan(
		&DataResponse.Id,
		&DataResponse.Name,
		&DataResponse.Data,
	)
	if err != nil {
		fmt.Println("RepositoryPareto UpdatePareto", err)
		return nil, err
	}

	return DataResponse, nil
}

func (r *RepositoryPareto) DeletePareto(DataRequest *models.ParetoDb) error {
	sql := `DELETE FROM "method" WHERE "id"=$1;`
	_, err := r.DB.Query(sql, DataRequest.Id)
	if err != nil {
		fmt.Println("RepositoryPareto DeletePareto", err)
		return err
	}

	return nil
}
