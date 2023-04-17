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

func (r *RepositoryPareto) GetPareto(id int) (*models.ParetoDb, error) {
	DataResponse := &models.ParetoDb{}
	sql := `select id, name, data from "tdata" where id=$1`
	err := r.DB.QueryRow(sql, id).Scan(
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
	sql := `insert into "tdata" (name, data) values ($1, $2) returning id, name, data;`
	err := r.DB.QueryRow(sql, DataRequest.Name, DataRequest.Data).Scan(
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
	sql := `UPDATE "tdata" SET "data" = $1 WHERE "id"=$2 returning id, name, data;`
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
	sql := `DELETE FROM "tdata" WHERE "id"=$1;`
	_, err := r.DB.Query(sql, DataRequest.Id)
	if err != nil {
		fmt.Println("RepositoryPareto DeletePareto", err)
		return err
	}

	return nil
}
