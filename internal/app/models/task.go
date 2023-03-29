package models

import "github.com/jackc/pgx/pgtype"

type Data struct {
	Id   int              `db:"id"`
	Name string           `db:"name"`
	Var1 pgtype.Int4Array `db:"var1"`
	Var2 pgtype.Int4Array `db:"var2"`
	Var3 pgtype.Int4Array `db:"var3"`
}

type DataReq struct {
	Name string `json:"name"`
	Var1 []int  `json:"var1"`
	Var2 []int  `json:"var2"`
	Var3 []int  `json:"var3"`
}
