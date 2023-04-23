package models

type ParetoDb struct {
	Id     int    `db:"id"`
	UserId int    `db:"user_id"`
	Name   string `db:"name"`
	Data   []byte `db:"data"`
}

type Pareto struct {
	Var1 []float64 `json:"var_1"`
	Var2 []float64 `json:"var_2"`
	Var3 []float64 `json:"var_3"`
}

type ParetoJson struct {
	Id   int       `json:"id"`
	Name string    `json:"name"`
	Var1 []float64 `json:"var1"`
	Var2 []float64 `json:"var2"`
	Var3 []float64 `json:"var3"`
}
