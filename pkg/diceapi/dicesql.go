package diceapi

import (
	"context"
	"fmt"
	"os"
	"time"

	"github.com/jackc/pgx/v4/pgxpool"
)

func Connect(conn_str string) *pgxpool.Pool {
	dbpool, err := pgxpool.Connect(context.Background(), conn_str)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Unable to connect to database %v\n", err)
		os.Exit(1)
	}
	return dbpool
}

func UserInsert(conn *pgxpool.Pool, username string, password string, date time.Time) error {
	commandTag, err := conn.Exec(context.Background(), "INSERT INTO users (username, password, date_created) VALUES ($1, $2, $3);", username, password, date)
	if err != nil {
		return err
	}
	if commandTag.RowsAffected() != 1 {
		err = fmt.Errorf("did not successfully insert %v into %v", username, date)
	}
	return err
}

func Hello() (string, error) {
	dbpool := Connect(os.Getenv("DATABASE_URL"))

	var greeting string
	err := dbpool.QueryRow(context.Background(), "select 'Hello, world!'").Scan(&greeting)
	if err != nil {
		fmt.Fprintf(os.Stderr, "QueryRow failed %v\n", err)
		os.Exit(1)
	}
	defer dbpool.Close()

	return greeting, err
}
