package models

type WeightedSumDb struct {
	Id     int    `db:"id"`
	UserId int    `db:"user_id"`
	Name   string `db:"name"`
	Data   []byte `db:"data"`
}

type WeightedSum struct {
	Var1 []string  `json:"var_1"`
	Var2 []float64 `json:"var_2"`
	Var3 []float64 `json:"var_3"`
	Var4 []float64 `json:"var_4"`
	Var5 []float64 `json:"var_5"`
}

type WeightedSumJson struct {
	Id   int       `json:"id"`
	Name string    `json:"name"`
	Var1 []string  `json:"var_1"`
	Var2 []float64 `json:"var_2"`
	Var3 []float64 `json:"var_3"`
	Var4 []float64 `json:"var_4"`
	Var5 []float64 `json:"var_5"`
}
