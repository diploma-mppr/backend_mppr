package models

type NansonDb struct {
	Id     int    `db:"id"`
	UserId int    `db:"user_id"`
	Name   string `db:"name"`
	Data   []byte `db:"data"`
}

type Nanson struct {
	Var1 []float64 `json:"var_1"`
	Var2 []string  `json:"var_2"`
	Var3 []string  `json:"var_3"`
	Var4 []string  `json:"var_4"`
}

type NansonJson struct {
	Id   int       `json:"id"`
	Name string    `json:"name"`
	Var1 []float64 `json:"var_1"`
	Var2 []string  `json:"var_2"`
	Var3 []string  `json:"var_3"`
	Var4 []string  `json:"var_4"`
}
