package initialize

import (
	"database/sql"
)

func OpenDB() (*sql.DB, error) {
	return sql.Open("sqlite3", "./backend/database/database.db")
}
