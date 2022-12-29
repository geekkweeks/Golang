package geek_golang_goroutines


import (
	"fmt"
	"testing"
	"time"
)

func TestBufferedChannel(t *testing.T) {
	channel := make(chan string, 3)
	defer close(channel)

	// anonymous func
	go func() {
		channel <- "Alfan"
		channel <- "Zahriyono"
		channel <- "Tampan"
	}()

	go func() {
		fmt.Println(<-channel)
		fmt.Println(<-channel)
		fmt.Println(<-channel)
	}()

	time.Sleep(2 * time.Second)
	fmt.Println("Finished")
}
