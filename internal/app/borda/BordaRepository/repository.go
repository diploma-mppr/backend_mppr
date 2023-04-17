package BordaRepository

import (
	"fmt"
	"gitgub.com/diploma-mppr/backend_mppr/internal/app/models"
	"github.com/jackc/pgx"
)

type RepositoryBorda struct {
	DB *pgx.ConnPool
}

func NewRepositoryBorda(db *pgx.ConnPool) *RepositoryBorda {
	return &RepositoryBorda{
		DB: db,
	}
}

func (r *RepositoryBorda) GetBorda(id int) (*models.BordaDb, error) {
	DataResponse := &models.BordaDb{}
	sql := `select id, name, data from "tdata" where id=$1`
	err := r.DB.QueryRow(sql, id).Scan(
		&DataResponse.Id,
		&DataResponse.Name,
		&DataResponse.Data,
	)
	if err != nil {
		fmt.Println("RepositoryBorda GetBorda", err)
		return nil, err
	}
	return DataResponse, nil
}

func (r *RepositoryBorda) SetBorda(DataRequest *models.BordaDb) (*models.BordaDb, error) {
	DataResponse := &models.BordaDb{}
	sql := `insert into "tdata" (name, data) values ($1, $2) returning id, name, data;`
	err := r.DB.QueryRow(sql, DataRequest.Name, DataRequest.Data).Scan(
		&DataResponse.Id,
		&DataResponse.Name,
		&DataResponse.Data,
	)
	if err != nil {
		fmt.Println("RepositoryBorda SetBorda", err)
		return nil, err
	}

	return DataResponse, nil
}

func (r *RepositoryBorda) UpdateBorda(DataRequest *models.BordaDb) (*models.BordaDb, error) {
	DataResponse := &models.BordaDb{}
	sql := `UPDATE "tdata" SET "data" = $1 WHERE "id"=$2 returning id, name, data;`
	err := r.DB.QueryRow(sql, DataRequest.Data, DataRequest.Id).Scan(
		&DataResponse.Id,
		&DataResponse.Name,
		&DataResponse.Data,
	)
	if err != nil {
		fmt.Println("RepositoryBorda UpdateBorda", err)
		return nil, err
	}

	return DataResponse, nil
}

func (r *RepositoryBorda) DeleteBorda(DataRequest *models.BordaDb) error {
	sql := `DELETE FROM "tdata" WHERE "id"=$1;`
	_, err := r.DB.Query(sql, DataRequest.Id)
	if err != nil {
		fmt.Println("RepositoryBorda DeleteBorda", err)
		return err
	}

	return nil
}
