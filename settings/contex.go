package settings

import (
	"github.com/jackc/pgx/v4"
)

const (
	// DbSchemaVersion contains database schema version
	DbSchemaVersion = "0.1"
	// DbSchemaDate contains database schema date creation in string format
	DbSchemaDate = "2020-Jan-01"
)

const (
	// ShortDateForm is short date format
	ShortDateForm = "2006-Jan-02"
)

var (
	// DbConnect contains connection to database
	DbConnect *pgx.Conn
)
