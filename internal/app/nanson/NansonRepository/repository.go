package NansonRepository

import (
	"fmt"
	"gitgub.com/diploma-mppr/backend_mppr/internal/app/models"
	"github.com/jackc/pgx"
)

type RepositoryNanson struct {
	DB *pgx.ConnPool
}

func NewRepositoryNanson(db *pgx.ConnPool) *RepositoryNanson {
	return &RepositoryNanson{
		DB: db,
	}
}

func (r *RepositoryNanson) GetNanson(id int) (*models.NansonDb, error) {
	DataResponse := &models.NansonDb{}
	sql := `select id, name, data from "tdata" where id=$1`
	err := r.DB.QueryRow(sql, id).Scan(
		&DataResponse.Id,
		&DataResponse.Name,
		&DataResponse.Data,
	)
	if err != nil {
		fmt.Println("RepositoryNanson GetNanson", err)
		return nil, err
	}
	return DataResponse, nil
}

func (r *RepositoryNanson) SetNanson(DataRequest *models.NansonDb) (*models.NansonDb, error) {
	DataResponse := &models.NansonDb{}
	sql := `insert into "tdata" (name, data) values ($1, $2) returning id, name, data;`
	err := r.DB.QueryRow(sql, DataRequest.Name, DataRequest.Data).Scan(
		&DataResponse.Id,
		&DataResponse.Name,
		&DataResponse.Data,
	)
	if err != nil {
		fmt.Println("RepositoryNanson SetNanson", err)
		return nil, err
	}

	return DataResponse, nil
}

func (r *RepositoryNanson) UpdateNanson(DataRequest *models.NansonDb) (*models.NansonDb, error) {
	DataResponse := &models.NansonDb{}
	sql := `UPDATE "tdata" SET "data" = $1 WHERE "id"=$2 returning id, name, data;`
	err := r.DB.QueryRow(sql, DataRequest.Data, DataRequest.Id).Scan(
		&DataResponse.Id,
		&DataResponse.Name,
		&DataResponse.Data,
	)
	if err != nil {
		fmt.Println("RepositoryNanson UpdateNanson", err)
		return nil, err
	}

	return DataResponse, nil
}

func (r *RepositoryNanson) DeleteNanson(DataRequest *models.NansonDb) error {
	sql := `DELETE FROM "tdata" WHERE "id"=$1;`
	_, err := r.DB.Query(sql, DataRequest.Id)
	if err != nil {
		fmt.Println("RepositoryNanson DeleteNanson", err)
		return err
	}

	return nil
}
