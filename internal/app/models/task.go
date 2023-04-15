package models

type DataDb struct {
	Id   int    `db:"id"`
	Name string `db:"name"`
	Var1 []byte `db:"var1"`
	Var2 []byte `db:"var2"`
	Var3 []byte `db:"var3"`
}

type Float64Json struct {
	MyFloat []float64 `json:"my_float"`
}

type DataJson struct {
	Id   int       `json:"id"`
	Name string    `json:"name"`
	Var1 []float64 `json:"var1"`
	Var2 []float64 `json:"var2"`
	Var3 []float64 `json:"var3"`
}
