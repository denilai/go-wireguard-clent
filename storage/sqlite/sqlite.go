package sqlite

import (
	"database/sql"
	"fmt"

	_ "github.com/mattn/go-sqlite3"
)

type Storage struct {
	db *sql.DB
}

func New(storagePath string) (*Storage, error) {
	op := "storage.sqlite.New"

	db, err := sql.Open("sqlite3", storagePath)
	if err != nil {
		return nil, fmt.Errorf("%s:%w", op, err)
	}

	stmt, err := db.Prepare(`
        CREATE TABLE IF NOT EXISTS clients(
            id INTEGER PRIMARY KEY,
            name TEXT NOT NULL UNIQUE,
            config_id INTEGER NOT NULL UNIQUE
        );
		CREATE INDEX IF NOT EXISTS idx_clients_id ON clients(id);
	`)
	if err != nil {
		return nil, fmt.Errorf("%s:%w", op, err)
	}

	_, err = stmt.Exec()
	if err != nil {
		return nil, fmt.Errorf("%s:%w", op, err)
	}
	return &(Storage{db}), nil

}
