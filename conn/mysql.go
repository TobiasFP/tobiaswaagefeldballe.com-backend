package conn

import (
	"database/sql"

	_ "github.com/go-sql-driver/mysql"
)

// GetMysqlDB returns a pointer to the standard sql db.
func GetMysqlDB() *sql.DB {
	// Open up our database connection.
	// I've set up a database on my local machine using phpmyadmin.
	// The database is called testDb
	server := "localhost"
	userName := "root"
	password := "What4reUL00k1ng4"
	database := "lttr"
	port := "3306"
	db, err := sql.Open("mysql", userName+":"+password+"@tcp("+server+":"+port+")/"+database)

	// if there is an error opening the connection, handle it
	if err != nil {
		panic(err.Error())
	}
	return db
}
