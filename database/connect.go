package database

import (
	"context"
	"log"

	"github.com/jackc/pgx/v4"

	"github.com/DmiProps/auf/settings"
)

//+++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++
// 1. Change terminal user
// >sudo -i -u postgres
// 2. Run postgres CLI
// >psql
// P.S. For quit:
// >\q
// 3. Create role
// >CREATE USER dmiprops WITH PASSWORD 'password';
// P.S. List roles:
// >SELECT rolname FROM pg_roles;
// 4. Create database
// >CREATE DATABASE aufdb OWNER dmiprops;
// P.S. Get list of databases:
// >SELECT datname FROM pg_database;
// P.S. Get location databases:
// >ps auxw | grep postgres | grep -- -D
// 5. Grant privileges
// 5.1. Connect to database
// >\c aufdb
// 5.2. Grant privileges to role
// >GRANT ALL PRIVILEGES ON ALL TABLES IN SCHEMA public TO "dmiprops";
// 6. Return terminal user
// >su - dmiprops
// 7. Connect to database
// >psql -daufdb
// 8. List of tables
// >SELECT table_name FROM information_schema.tables WHERE table_schema NOT IN ('information_schema','pg_catalog');
//---------------------------------------------------------------------------------

// Connect make connect to database
func Connect() {

	// Get connection
	conn, err := pgx.Connect(context.Background(), settings.AppSettings.Database.DbConnection)
	if err != nil {
		log.Fatalln("Error GetConnect(): ", err)
		return
	}
	settings.DbConnect = conn

	// Get database schema version
	dbVersion, err := getVersion(settings.DbConnect)
	if err != nil {
		log.Fatalln("Error getVersion(): ", err)
		return
	}

	// Create or update database schema
	for dbVersion != settings.DbSchemaVersion {
		switch dbVersion {
		case "":
			err = createActualSchema(settings.DbConnect)
			if err != nil {
				log.Fatalln("Error createActualSchema(): ", err)
			} else {
				dbVersion = settings.DbSchemaVersion
			}
		case "0.1":
			err = updateFrom01To02(settings.DbConnect)
			if err != nil {
				log.Fatalln("Error updateFrom01To02(): ", err)
			} else {
				dbVersion = "0.2"
			}
		}
		if err != nil {
			break
		}
	}

}

func getVersion(conn *pgx.Conn) (string, error) {

	dbVersion := "" // version database schema

	// Create table 'version' if not exists
	_, err := conn.Exec(
		context.Background(),
		`create table if not exists version (
			id				integer,		-- always equal to 0
			schema_version	varchar(24),	-- db schema vertion,
			schema_date		date,			-- db schema date of creation
			primary key (id)
		)`)
	if err != nil {
		return "", err
	}

	// Get dabatase schema version
	rows, err := conn.Query(
		context.Background(),
		`select schema_version from version`)
	if err != nil {
		return "", err
	}
	defer rows.Close()

	if rows.Next() {
		rows.Scan(&dbVersion)
	}

	return dbVersion, nil

}
