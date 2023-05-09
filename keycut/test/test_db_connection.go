package test

import (
	"database/sql"
	"keycut/keycut/settings"
	"keycut/keycut/storage"
)

func OpenTestDBConnection() (*sql.DB, error) {
	return storage.NewDatabaseConnection(settings.TestAppDatabase)
}
