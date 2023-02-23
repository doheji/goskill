package db

import (
	"database/sql"
	_ "github.com/lib/pq"
	"log"
	"os"
	"testing"
)

var testStore *Store
var driverName = "postgres"
var driverSource = "postgresql://root:secret@localhost:1234/go_skill_db?sslmode=disable"

func TestMain(m *testing.M) {
	conn, err := sql.Open(driverName, driverSource)
	if err != nil {
		log.Fatalln("Failed to run test as the connection to database was not established", err)
	}
	testStore = NewStore(conn)
	os.Exit(m.Run())
}
