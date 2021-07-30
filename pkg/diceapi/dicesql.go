package diceapi

import (
	"context"
	"fmt"
	"os"

	"github.com/jackc/pgx/v4/pgxpool"
)

func connect(conn_str string) *pgxpool.Pool {
	dbpool, err := pgxpool.Connect(context.Background(), conn_str)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Unable to connect to database %v\n", err)
		os.Exit(1)
	}
	return dbpool
}

func Hello() (string, error) {
	dbpool := connect(os.Getenv("DATABASE_URL"))

	var greeting string
	err := dbpool.QueryRow(context.Background(), "select 'Hello, world!'").Scan(&greeting)
	if err != nil {
		fmt.Fprintf(os.Stderr, "QueryRow failed %v\n", err)
		os.Exit(1)
	}
	defer dbpool.Close()

	return greeting, err
}
