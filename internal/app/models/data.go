package models

type DataDb struct {
	Id         int    `db:"id"`
	MethodName string `db:"method_name"`
	Name       string `db:"name"`
}

type DataJson struct {
	Id         int    `json:"id"`
	MethodName string `json:"method_name"`
	Name       string `json:"name"`
}
