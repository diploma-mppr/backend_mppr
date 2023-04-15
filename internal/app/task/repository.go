package task

import (
	"fmt"
	"gitgub.com/diploma-mppr/backend_mppr/internal/app/models"
	"github.com/jackc/pgx"
)

type RepositoryTask struct {
	DB *pgx.ConnPool
}

func NewRepositoryTask(db *pgx.ConnPool) *RepositoryTask {
	return &RepositoryTask{DB: db}
}

//func (r *RepositoryTask) GetData() (data models.DataDb, err error) {
//	err = r.DB.QueryRow(
//		`select id, name, var1, var2, var3 from data where id=1`,
//	).Scan(
//		&data.Id,
//		&data.Name,
//		&data.Var1,
//		&data.Var2,
//		&data.Var3)
//	fmt.Println(data)
//	fmt.Println(err)
//	return
//}

func (r *RepositoryTask) SetData(data models.DataDb) (data1 models.DataDb, err error) {
	err = r.DB.QueryRow(
		`insert into pareto (name, var1, var2, var3) values ($1, $2, $3, $4)
			returning id, name, var1, var2, var3;`,
		data.Name, data.Var1, data.Var2, data.Var3,
	).Scan(
		&data1.Id,
		&data1.Name,
		&data1.Var1,
		&data1.Var2,
		&data1.Var3,
	)
	fmt.Println("repository", err)
	if err != nil {
		return models.DataDb{}, err
	}

	fmt.Println(data1)
	return
}
