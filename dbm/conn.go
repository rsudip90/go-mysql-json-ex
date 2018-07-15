package dbm

import (
	"go-mysql-json-ex/dbm/internal"
)

// InitDB will initialize connection to db
func InitDB() {
	internal.InitDB()
}

// CloseDB will close the connection to db
func CloseDB() {
	internal.CloseDB()
}
