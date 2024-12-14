package mymigrate

import (
	"database/sql"
	"fmt"
	"os"
	"path/filepath"
)

func Migrate(params Params) {
	pathDatabase := filepath.Join("deploy", "migrations", "000001_create_calls_database.up.sql")
	pathTables := filepath.Join("deploy", "migrations", "000001_create_calls_table.up.sql")
	createDatabaseBytes, _ := os.ReadFile(pathDatabase)
	createTablesBytes, _ := os.ReadFile(pathTables)
	createDatabaseStr := string(createDatabaseBytes)
	createTablesStr := string(createTablesBytes)

	db, _ := sql.Open("postgres", fmt.Sprintf("sslmode=disable user=%s password=%s host=%s ", params.DatabaseUser, params.DatabasePassword, params.DatabaseHost))
	db.Exec(createDatabaseStr)
	db.Close()

	dbTables, _ := sql.Open("postgres", fmt.Sprintf("sslmode=disable user=%s password=%s host=%s dbname=%s ", params.DatabaseUser, params.DatabasePassword, params.DatabaseHost, params.DatabaseName))
	dbTables.Exec(createTablesStr)
	dbTables.Close()
}

type Params struct {
	DatabaseHost     string
	DatabaseUser     string
	DatabasePassword string
	DatabaseName     string
}
