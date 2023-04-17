package DataRepository

import (
	"fmt"
	"gitgub.com/diploma-mppr/backend_mppr/internal/app/models"
	"github.com/jackc/pgx"
)

type RepositoryData struct {
	DB *pgx.ConnPool
}

func NewRepositoryData(db *pgx.ConnPool) *RepositoryData {
	return &RepositoryData{
		DB: db,
	}
}

func (r *RepositoryData) GetAll() (*[]models.DataDb, error) {
	var DataResponse []models.DataDb
	sql := `select "id", "name" from "tdata";`
	rows, err := r.DB.Query(sql)
	if err != nil {
		fmt.Println("RepositoryData GetAll", err)
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var buff models.DataDb
		err = rows.Scan(
			&buff.Id,
			&buff.Name,
		)
		if err != nil {
			fmt.Println("RepositoryData GetAll", err)
			return nil, err
		}

		DataResponse = append(DataResponse, buff)
	}

	return &DataResponse, nil
}
