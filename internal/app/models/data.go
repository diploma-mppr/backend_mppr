package models

type DataDb struct {
	Id   int    `db:"id"`
	Name string `db:"name"`
}

type DataJson struct {
	Id   int    `json:"id"`
	Name string `json:"name"`
}
