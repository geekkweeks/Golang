package geek_golang_goroutines

import (
	"fmt"
	"strconv"
	"testing"
)

func TestRangeChannel(t *testing.T) {
	channel := make(chan string)

	go func() {
		for i := 0; i < 20; i++ {
			channel <- "Loop: " + strconv.Itoa(i)
		}
		close(channel)
	}()

	go func() {
		for data := range channel {
			fmt.Println("Data - ", data)
		}
	}()

	fmt.Println("Finished")
}
