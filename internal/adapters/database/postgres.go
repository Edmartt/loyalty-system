package database

import (
	"errors"
	"fmt"
	"os"

	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq" // sqlx uses it indirectly
)

type Postgres struct {
	user       string
	db         string
	password   string
	host       string
	port       string
	connection *sqlx.DB
}

func NewPostgresConn() (*Postgres, error) {
	pgConfig := &Postgres{
		user:     os.Getenv("PG_USER"),
		db:       os.Getenv("PG_DB"),
		password: os.Getenv("PG_PASSWORD"),
		host:     os.Getenv("HOST"),
		port:     os.Getenv("PORT"),
	}

	if pgConfig.user == "" || pgConfig.db == "" || pgConfig.password == "" || pgConfig.host == "" || pgConfig.port == "" {
		return nil, errors.New("environment variables not set")
	}

	return pgConfig, nil

}

// GetConnection connects to specific database
func (p *Postgres) GetConnection() (*sqlx.DB, error) {

	if p.connection != nil {
		return p.connection, nil
	}

	db, err := sqlx.Connect("postgres", fmt.Sprintf("user=%s dbname=%s password=%s host=%s port=%s sslmode=disable", p.user, p.db, p.password, p.host, p.port))

	if err != nil {
		return nil, fmt.Errorf("DB connection error: %v", err)
	}

	p.connection = db

	return p.connection, nil
}
