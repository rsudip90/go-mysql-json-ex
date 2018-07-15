package internal

import (
	"database/sql"
	"log"
	"runtime/debug"

	"github.com/go-sql-driver/mysql"
)

// dbConf holds mysql connection configuration
var dbConf = mysql.Config{
	User:      "",
	Passwd:    "",
	Net:       "tcp",
	Addr:      "127.0.0.1:3306",
	DBName:    "jsontest",
	ParseTime: true,
}

// DBFields type
type DBFields map[string]string

// DBManager is global single object which holds DB and other stuff
var DBManager struct {
	DB *sql.DB
	PrepSQL
	DBFields
}

// CheckErr will exit the program with debug print stack
func CheckErr(err error) {
	if err != nil {
		debug.PrintStack()
		log.Fatal(err)
	}
}

// InitDB will initialize the connection using the configuration
// defined in dbConf variable. It also verifies the connection line
// by pinging. Later, it prepares all statements defined in
// buildStatements() function.
func InitDB() {

	// open
	var err error
	DBManager.DB, err = sql.Open("mysql", dbConf.FormatDSN())
	CheckErr(err)

	// ping the database line
	err = DBManager.DB.Ping()
	CheckErr(err)

	// init map
	DBManager.DBFields = make(DBFields)

	// build prepared statements
	buildStatements()
}

// CloseDB will close the connection held by DBManager struct's DB
// field
func CloseDB() { // close the database
	DBManager.DB.Close()
}
