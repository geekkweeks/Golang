package geek_golang_goroutines

import (
	"fmt"
	"testing"
	"time"
)

func GiveMeResponse(channel chan string) {
	time.Sleep(2 * time.Second)
	channel <- "Alfan Zahriyono"
}

func TestChannelAsParameter(_ *testing.T) {
	channel := make(chan string)
	// sukses atau berhasil close channel akan kepanggil
	defer close(channel)

	go GiveMeResponse(channel)

	data := <-channel
	fmt.Println(data)

	time.Sleep(5 * time.Second)
}