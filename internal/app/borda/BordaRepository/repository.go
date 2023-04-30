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

func (r *RepositoryBorda) GetBorda(id int, UserId int) (*models.BordaDb, error) {
	DataResponse := &models.BordaDb{}
	sql := `select "id", "name", "data" from "method" where id=$1 and "user_id" = $2`
	err := r.DB.QueryRow(sql, id, UserId).Scan(
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
	sql := `insert into "method" ("user_id", "name", "data", "method_name") values ($1, $2, $3, 'borda') returning id, name, data;`
	err := r.DB.QueryRow(sql, DataRequest.UserId, DataRequest.Name, DataRequest.Data).Scan(
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
	sql := `UPDATE "method" SET "data" = $1 WHERE "id"=$2 returning id, name, data;`
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
	sql := `DELETE FROM "method" WHERE "id"=$1;`
	_, err := r.DB.Query(sql, DataRequest.Id)
	if err != nil {
		fmt.Println("RepositoryBorda DeleteBorda", err)
		return err
	}

	return nil
}
