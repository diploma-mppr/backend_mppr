package models

type BaseCriteriaDb struct {
	Id   int    `db:"id"`
	Name string `db:"name"`
	Data []byte `db:"data"`
}

type BaseCriteria struct {
	Var1 []bool `json:"var_1"`
}

type BaseCriteriaJson struct {
	Id   int    `json:"id"`
	Name string `json:"name"`
	Var1 []bool `json:"var1"`
}
