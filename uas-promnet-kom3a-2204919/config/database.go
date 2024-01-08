package config

import "database/sql"
import _ "github.com/go-sql-driver/mysql"

func DBConnectiom() (*sql.DB, error) {
	dbDriver := "mysql"
	dbUser := "root"
	dbPass := ""
	dbName := "db_2204919_nanditarizkynaazzahra_uas"

	db, err := sql.Open(dbDriver, dbUser+":"+dbPass+"@/"+dbName)
	return db, err
}