package postgres

import (
	"database/sql"
	"fmt"
	"log"
	"time"

	"github.com/diyliv/anime/configs"
	_ "github.com/lib/pq"
)

func InitPsqlDB() *sql.DB {
	cfg := configs.ReadConfig()

	postgresConnect := fmt.Sprintf(
		"host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		cfg.Postgres.Host,
		cfg.Postgres.Port,
		cfg.Postgres.User,
		cfg.Postgres.Password,
		cfg.Postgres.DBName)

	client, err := sql.Open("postgres", postgresConnect)
	if err != nil {
		log.Fatalf("Error while conencting to PostgresDB: %v\n", err)
	}
	client.SetConnMaxLifetime(time.Minute * 3)
	client.SetMaxOpenConns(10)
	client.SetMaxIdleConns(10)

	return client
}
