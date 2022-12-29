package music_crud

import (
	"context"
	"fmt"
	"testing"
)

func TestExecSqlQuery(t *testing.T) {
	db := GetConnection()
	defer db.Close()

	ctx := context.Background()

	sqlString := "INSERT INTO ARTISTS(ID,NAME) VALUES(2, 'ZAHRIYONO')"
	_, err := db.ExecContext(ctx, sqlString)
	if err != nil {
		panic(err)
	}

	fmt.Println("Data has been inserted")

}
