package geek_golang_goroutines

import (
	"fmt"
	"testing"
	"time"
)

func TestCreateChannel(t *testing.T) {
	channel := make(chan string)
	defer close(channel)

	// anonymous function
	go func ()  {
		time.Sleep(2 * time.Second)
		channel <- "Alfan zahriyono"
		fmt.Println("Selesai kirim data ke channel")
	}()

	// ambil data channel
	data := <-channel
	fmt.Println(data)
	time.Sleep(5 * time.Second)
	
}