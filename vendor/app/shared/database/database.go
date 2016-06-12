package database

import (
	"fmt"
	"log"

	_ "github.com/go-sql-driver/mysql" // MySQL driver
	"github.com/jmoiron/sqlx"
)

var (
	SQL      *sqlx.DB
	database Info
)

// MySQL Info is the details for the database connection
type Info struct {
	Username  string
	Password  string
	Name      string
	Hostname  string
	Port      int
	Parameter string
}

// DSN returns the Data Source Name
func DSN(ci Info) string {
	// Example: root:@tcp(localhost:3306)/test
	return ci.Username +
		":" +
		ci.Password +
		"@tcp(" +
		ci.Hostname +
		":" +
		fmt.Sprintf("%d", ci.Port) +
		")/" +
		ci.Name + ci.Parameter
}

// Connect to the database
func Connect(i Info) {
	var err error

	// Connect to MySQL
	if SQL, err = sqlx.Connect("mysql", DSN(i)); err != nil {
		log.Println("SQL Driver Error", err)
	}
}

// ReadConfig returns the database information
func ReadConfig() Info {
	return database
}
