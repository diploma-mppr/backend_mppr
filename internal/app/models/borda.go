package models

type BordaDb struct {
	Id   int    `db:"id"`
	Name string `db:"name"`
	Data []byte `db:"data"`
}

type Borda struct {
	Var1 []float64 `json:"var_1"`
	Var2 []string  `json:"var_2"`
	Var3 []string  `json:"var_3"`
	Var4 []string  `json:"var_4"`
}

type BordaJson struct {
	Id   int       `json:"id"`
	Name string    `json:"name"`
	Var1 []float64 `json:"var_1"`
	Var2 []string  `json:"var_2"`
	Var3 []string  `json:"var_3"`
	Var4 []string  `json:"var_4"`
}
