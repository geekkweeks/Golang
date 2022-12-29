package music_crud

import (
	"context"
	"fmt"
	"testing"
)

func TestExecInsertSql(t *testing.T) {
	db := GetConnection()
	defer db.Close()

	ctx := context.Background()

	sqlString := "INSERT INTO ARTISTS(ID,NAME) VALUES(4, 'Gulung')"
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
		var name string
		err = rows.Scan(&id, &name)
		if err != nil {
			panic(err)
		}
		fmt.Println("ID: ", id)
		fmt.Println("Name: ", name)
	}
}
