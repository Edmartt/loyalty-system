package ports

import (
	"github.com/jmoiron/sqlx"
)

type IConnection interface {
	GetConnection() (*sqlx.DB, error)
}
