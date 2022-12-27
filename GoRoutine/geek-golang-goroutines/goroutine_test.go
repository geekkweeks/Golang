package geek_golang_goroutines

import (
	"fmt"
	"testing"
	"time"
)

func RunHelloWorld() {
	fmt.Println("Hello world")
}

func TestCreateGoRoutine(t *testing.T) {
	go RunHelloWorld()
	fmt.Println("Sukses selalu")

	time.Sleep(1 * time.Second)
}

func DisplayNumber(number int) {
	fmt.Println("Display number: ", number)
}

func TestManyGoroutine(t *testing.T) {
	for i := 0; i < 100000; i++ {
		go DisplayNumber(i)
	}

	time.Sleep(2 * time.Second)
}
