package database

import (
	"database/sql"
	"errors"
	"fmt"
	"os"
	"strings"

	_ "modernc.org/sqlite"
)

type SQLiteManager struct{}

func NewSQLiteManager() *SQLiteManager {
	return &SQLiteManager{}
}

const defaultSchema = `
	CREATE TABLE IF NOT EXISTS secrets (
	id TEXT PRIMARY KEY,
	content TEXT NOT NULL,
	created_at TIMESTAMP,
	duration INTEGER NOT NULL,
	password TEXT,
	views INTEGER NOT NULL
);
`

func (s *SQLiteManager) InitDB() (*sql.DB, error) {
	path := os.Getenv("BYEMBER_SQLITE_DATABASE_NAME")
	if strings.TrimSpace(path) == "" {
		path = "byember.db"
	}

	firstTime := false
	if _, err := os.Stat(path); errors.Is(err, os.ErrNotExist) {
		firstTime = true
	}

	db, err := sql.Open("sqlite", path)
	if err != nil {
		return nil, fmt.Errorf("failed to open sqlite db: %w", err)
	}

	if firstTime {
		if _, err := db.Exec(defaultSchema); err != nil {
			return nil, fmt.Errorf("failed to initialize schema: %w", err)
		}
	}

	return db, nil
}

// func GetDB() (*sql.DB, error) {
// 	path := os.Getenv("BYEMBER_SQLITE_DATABASE_NAME")
// 	if strings.TrimSpace(path) == "" {
// 		path = "byember.db"
// 	}

// 	db, err := sql.Open("sqlite", path)
// 	if err != nil {
// 		return nil, fmt.Errorf("failed to open sqlite db: %w", err)
// 	}

// 	return db, nil
// }
