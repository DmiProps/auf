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

// GetConnect return connection to database
func GetConnect() {

	// Get connection
	conn, err := pgx.Connect(context.Background(), settings.AppSettings.DbConnection)
	if err != nil {
		log.Fatalln("Error GetConnect(): ", err)
		return
	}
	defer conn.Close(context.Background())

	// Check tables
	//TO-DO

}
