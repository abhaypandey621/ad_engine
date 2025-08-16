package repository

import (
	"database/sql"
	"fmt"
	"os"
	"strings"

	_ "github.com/go-sql-driver/mysql"
)

// NewMySQLDB creates a new MySQL DB connection using the MYSQL_DSN environment variable.
func NewMySQLDB() (*sql.DB, error) {
	dsn := os.Getenv("MYSQL_DSN") // e.g. user:password@tcp(localhost:3306)/dbname
	if dsn == "" {
		return nil, fmt.Errorf("MYSQL_DSN environment variable not set")
	}
	return sql.Open("mysql", dsn)
}

// NewMySQLDBWithDSN creates a new MySQL DB connection using the provided DSN string.
func NewMySQLDBWithDSN(dsn string) (*sql.DB, error) {
	if dsn == "" {
		return nil, fmt.Errorf("MySQL DSN is empty")
	}
	return sql.Open("mysql", dsn)
}

// FetchAllAppIDs returns a map of app_identifier (lowercase) to app_id.
func FetchAllAppIDs(db *sql.DB) (map[string]int, error) {
	query := `SELECT app_identifier, app_id FROM app`
	rows, err := db.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	result := make(map[string]int)
	for rows.Next() {
		var ident string
		var id int
		if err := rows.Scan(&ident, &id); err != nil {
			return nil, err
		}
		result[strings.ToLower(ident)] = id
	}
	return result, nil
}
