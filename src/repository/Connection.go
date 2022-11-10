package repository

import (
	"fmt"
	"os"
	"truphone-api/src/utils"

	"github.com/jackc/pgx"
)

type DB struct {
	sql *pgx.Conn
}

// func (s *SqlSupplier) InitConnection() *pgx.Conn {
// 	s.sql, _ = GetConn()
// 	return s.sql
// }

func GetConn() (*DB, error) {
	utils.LoadEnv()

	DSN := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%v sslmode=disable TimeZone=America/Sao_Paulo",
		os.Getenv("HOSTNAME"),
		os.Getenv("DB_USER"),
		os.Getenv("DB_PASS"),
		os.Getenv("DB_NAME"),
		os.Getenv("PORT"))

	config, err := pgx.ParseDSN(DSN)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error to parse DSN: %v\n", err)
		os.Exit(1)
	}

	conn, err := pgx.Connect(config)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Unable to connect to database: %v\n", err)
		os.Exit(1)
	}

	return &DB{sql: conn}, nil
}
