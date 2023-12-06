package database

import (
	"database/sql"
	"fmt"
	_ "github.com/lib/pq"
	"log"
)

var (
// TODO: query errors
)

type Postgres struct {
	DB *sql.DB
}

func New(username, password, host, dbname string, port int) (*Postgres, error) {
	const op = "postgres.New"

	url := fmt.Sprintf("postgres://%s:%s@%s:%d/%s?sslmode=disable",
		username,
		password,
		host,
		port,
		dbname,
	)

	db, err := sql.Open("postgres", url)
	if err != nil {
		return nil, fmt.Errorf("%s: %s", op, err)
	}

	if err = db.Ping(); err != nil {
		return nil, fmt.Errorf("%s: %s", op, err)
	}

	//TODO: migrations

	log.Println("postgres repository connected successfully!")

	return &Postgres{DB: db}, nil
}

func (p *Postgres) Close() error {
	const op = "postgres.Close"

	return fmt.Errorf("%s: %s", op, p.DB.Close())
}
