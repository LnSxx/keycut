package storage

import (
	"context"
	"database/sql"
	"fmt"
	"keycut/keycut/settings"

	_ "github.com/lib/pq"
)

type DatabaseConnection interface {
	QueryRow(string, ...any) *sql.Row
	Query(string, ...any) (*sql.Rows, error)
	QueryContext(context.Context, string, ...any) (*sql.Rows, error)
	Exec(string, ...any) (sql.Result, error)
}

func NewDatabaseConnection(
	dbSettings settings.DatabaseSettings,
) (*sql.DB, error) {
	connection := &sql.DB{}

	psqlInfo := fmt.Sprintf(`
		host=%s 
		port=%d 
		user=%s 
		dbname=%s
		sslmode=disable
	`,
		dbSettings.Host,
		dbSettings.Port,
		dbSettings.User,
		dbSettings.Dbname)

	var err error

	connection, err = sql.Open("postgres", psqlInfo)

	if err != nil {
		return connection, err
	}

	return connection, nil
}
