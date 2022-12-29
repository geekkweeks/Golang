package music_crud

import (
	"database/sql"
	"fmt"
	_ "github.com/lib/pq"
	"testing"
)

func TestOpenConnection(t *testing.T) {
	const (
		host     = "localhost"
		port     = 5432
		user     = "postgres"
		password = "asdasdasd"
		dbname   = "fdfdf"
	)

	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s "+
		"password=%s dbname=%s sslmode=disable",
		host, port, user, password, dbname)

	db, err := sql.Open("postgres", psqlInfo)
	if err != nil {
		panic(err)
	}
	defer db.Close()
}
