package music_crud

import (
	"context"
	"database/sql"
	"fmt"
	"testing"
)

func TestExecInsertSql(t *testing.T) {
	db := GetConnection()
	defer db.Close()

	ctx := context.Background()

	sqlString := "INSERT INTO ARTISTS (NAME) VALUES ('Cinta Laura')"
	_, err := db.ExecContext(ctx, sqlString)
	if err != nil {
		panic(err)
	}

	fmt.Println("Data has been inserted")

}

func TestExecSelectSql(t *testing.T) {
	db := GetConnection()
	defer db.Close()

	ctx := context.Background()

	sqlString := "SELECT ID, NAME FROM ARTISTS"
	rows, err := db.QueryContext(ctx, sqlString)
	if err != nil {
		panic(err)
	}
	// dont forget to always close row
	defer rows.Close()

	// iterate data from return of query select
	// if rows.Next = false iterate will stop
	for rows.Next() {
		var id int

		// to handle nullable from SQL, using sql.NullString
		var name sql.NullString

		err = rows.Scan(&id, &name)
		if err != nil {
			panic(err)
		}
		fmt.Println("ID: ", id)
		fmt.Println("Name: ", name)
	}
}

// To avoid SQL injection, the best way is execute the SQL by using paramter
func TestExecSelectWithParamSql(t *testing.T) {
	db := GetConnection()
	defer db.Close()

	ctx := context.Background()

	var paramId int = 1
	sqlString := "SELECT ID, NAME FROM ARTISTS WHERE ID = $1"
	rows, err := db.QueryContext(ctx, sqlString, paramId)
	if err != nil {
		panic(err)
	}
	// dont forget to always close row
	defer rows.Close()

	// if rows.Next = false iterate will stop
	if rows.Next() {
		var id int

		// to handle nullable from SQL, using sql.NullString
		var name sql.NullString

		err = rows.Scan(&id, &name)
		if err != nil {
			panic(err)
		}
		fmt.Println("ID: ", id)
		fmt.Println("Name: ", name)
	}
}

func TestGetLastIdSql(t *testing.T) {
	// get DB connection
	db := GetConnection()
	// dont forget to close db connection
	defer db.Close()

	// context allow to cancel the process
	ctx := context.Background()
	artistName := "cantik manis manja"

	sqlString := "INSERT INTO ARTISTS (NAME) VALUES ($1)"
	lastInsertId  := 0
	err := db.QueryRowContext(ctx, sqlString, artistName).Scan(&lastInsertId )
	if err != nil {
		panic(err)
	}

	fmt.Println("Data has been inserted with id: ", lastInsertId )

}
