package tools

import (
	"fmt"
	"github.com/jackc/pgx"
	"log"
)

func NewPostgresqlX() (*pgx.ConnPool, error) {
	dsn := fmt.Sprintf("user=%s dbname=%s password=%s host=%s port=%d", "yutfut", "yutfut", "yutfut", "db", 5432)
	conn, err := pgx.ParseConnectionString(dsn)
	if err != nil {
		log.Fatalln("cant parse config", err)
	}

	db, err := pgx.NewConnPool(pgx.ConnPoolConfig{
		ConnConfig:     conn,
		MaxConnections: 1000,
		AfterConnect:   nil,
		AcquireTimeout: 0,
	})

	if err != nil {
		fmt.Println("db error")
		fmt.Println(err.Error())
		log.Fatalf("Error %s occurred during connection to database", err)
	}
	fmt.Println("db connect done")

	return db, nil
}
