package models

type NansonDb struct {
	Id     int    `db:"id"`
	UserId int    `db:"user_id"`
	Name   string `db:"name"`
	Data   []byte `db:"data"`
}

type Nanson struct {
	Var1 []float64 `json:"var_1"`
}

type NansonJson struct {
	Id   int       `json:"id"`
	Name string    `json:"name"`
	Var1 []float64 `json:"var1"`
}
