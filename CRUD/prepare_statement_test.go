package music_crud

import (
	"context"
	"fmt"
	"strconv"
	"testing"
)

func TestPrepareStatementInSql(t *testing.T) {
	// get connection
	db := GetConnection()
	defer db.Close()

	ctx := context.Background()

	sqlString := "INSERT INTO ARTISTS (NAME) VALUES ($1)"
	statement, err := db.PrepareContext(ctx, sqlString)
	if err != nil {
		panic(err)
	}
	// close prepare statemen
	defer statement.Close()

	for i := 0; i < 100; i++ {
		artistName := "Bondan" + strconv.Itoa(i)

		_, err := statement.ExecContext(ctx, artistName)
		if err != nil {
			panic(err)
		}
		fmt.Println("Artist ID: ", artistName)

	}

}
