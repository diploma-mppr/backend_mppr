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

func (r *RepositoryData) GetAll(UserId int) (*[]models.DataDb, error) {
	var DataResponse []models.DataDb

	var Data []interface{}
	Data = append(Data, UserId)

	sql := `select "id", "name", "method_name" from "method" where "user_id" = $1 ORDER BY "id";`
	rows, err := r.DB.Query(sql, Data...)
	if err != nil {
		fmt.Println("RepositoryData GetAll", err)
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var buff models.DataDb
		err = rows.Scan(
			&buff.Id,
			&buff.MethodName,
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
