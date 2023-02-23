package main

import (
	"database/sql"
	"fmt"
	_ "github.com/lib/pq"
	"goskill/api"
	db "goskill/db/sqlc"
	"log"
)

var driverName = "postgres"
var driverSource = "postgresql://root:secret@localhost:1234/go_skill_db?sslmode=disable"

// @title           Go Skill
// @version         1.0
// @description     A skill tracker and development tool
// @termsOfService  https://tos.santoshk.dev

// @contact.name   David Oheji
// @contact.url    https://twitter.com/ejedavy
// @contact.email  ejeohejidavid@gmail.com

// @license.name  Apache 2.0
// @license.url   http://www.apache.org/licenses/LICENSE-2.0.html

// @host      localhost:8000
// @BasePath  /
func main() {
	var store *db.Store
	conn, err := sql.Open(driverName, driverSource)
	if err != nil {
		log.Fatal(err)
	}
	store = db.NewStore(conn)
	server := api.NewServer(store)
	fmt.Println("Starting the server")
	server.Start(":8000")
}
