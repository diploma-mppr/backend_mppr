package task

import (
	"gitgub.com/diploma-mppr/backend_mppr/internal/app/models"
	"github.com/jackc/pgx"
)

type RepositoryTask struct {
	DB *pgx.ConnPool
}

func NewRepositoryTask(db *pgx.ConnPool) *RepositoryTask {
	return &RepositoryTask{DB: db}
}

func (r *RepositoryTask) GetData(id int) (*models.DataDb, error) {
	DataResponse := &models.DataDb{}
	err := r.DB.QueryRow(
		`select id, name, var1, var2, var3 from DataResponse where id=$1`, id,
	).Scan(
		&DataResponse.Id,
		&DataResponse.Name,
		&DataResponse.Var1,
		&DataResponse.Var2,
		&DataResponse.Var3,
	)
	if err != nil {
		return nil, err
	}
	return DataResponse, nil
}

func (r *RepositoryTask) SetData(DataRequest *models.DataDb) (*models.DataDb, error) {
	DataResponse := &models.DataDb{}
	sql := `insert into pareto (name, var1, var2, var3) values ($1, $2, $3, $4) returning id, name, var1, var2, var3;`
	err := r.DB.QueryRow(sql, DataRequest.Name, DataRequest.Var1, DataRequest.Var2, DataRequest.Var3).Scan(
		&DataResponse.Id,
		&DataResponse.Name,
		&DataResponse.Var1,
		&DataResponse.Var2,
		&DataResponse.Var3,
	)
	if err != nil {
		return nil, err
	}

	return DataResponse, nil
}
