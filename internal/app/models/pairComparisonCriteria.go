package models

type PairComparisonCriteriaDb struct {
	Id     int    `db:"id"`
	UserId int    `db:"user_id"`
	Name   string `db:"name"`
	Data   []byte `db:"data"`
}

type PairComparisonCriteria struct {
	Var1  []float64 `json:"var_1"`
	Var2  []float64 `json:"var_2"`
	Var3  []float64 `json:"var_3"`
	Var4  []float64 `json:"var_4"`
	Var5  []float64 `json:"var_5"`
	Var6  []float64 `json:"var_6"`
	Var7  []float64 `json:"var_7"`
	Var8  []float64 `json:"var_8"`
	Var9  []float64 `json:"var_9"`
	Var10 []float64 `json:"var_10"`
}

type PairComparisonCriteriaJson struct {
	Id    int       `json:"id"`
	Name  string    `json:"name"`
	Var1  []float64 `json:"var_1"`
	Var2  []float64 `json:"var_2"`
	Var3  []float64 `json:"var_3"`
	Var4  []float64 `json:"var_4"`
	Var5  []float64 `json:"var_5"`
	Var6  []float64 `json:"var_6"`
	Var7  []float64 `json:"var_7"`
	Var8  []float64 `json:"var_8"`
	Var9  []float64 `json:"var_9"`
	Var10 []float64 `json:"var_10"`
}
