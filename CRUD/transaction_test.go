package music_crud

import (
	"context"
	"fmt"
	"strconv"
	"testing"
)


func TestTransactionInSql(t *testing.T) {
	// get connection
	db := GetConnection()
	defer db.Close()

	ctx := context.Background()

	tx, err := db.Begin()
	if err != nil {
		panic(err)
	}
	
	sqlString := "INSERT INTO ARTISTS (NAME) VALUES ($1)"
	// do transaction
	for i := 0; i < 20; i++ {
		artistName := "Brian " + strconv.Itoa(i)

		_, err := tx.ExecContext(ctx, sqlString, artistName)
		if err != nil {
			panic(err)
		}
		fmt.Println("Artist ID: ", artistName)

	}

	err = tx.Commit()
	if err != nil {
		panic(err)
	}

}
