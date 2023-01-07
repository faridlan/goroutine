package test

import (
	"fmt"
	"strconv"
	"testing"
	"time"
)

func TestBufferChannel(t *testing.T) {
	channel := make(chan string, 3)

	fmt.Println(cap(channel))
	fmt.Println(len(channel))

	go func() {
		channel <- "faridlan"
		channel <- "nul"
		channel <- "hakim"
	}()

	go func() {
		fmt.Println(<-channel)
		fmt.Println(<-channel)
		fmt.Println(<-channel)
	}()

	time.Sleep(2 * time.Second)

	fmt.Println("selesai")
}

func TestRangeChannel(t *testing.T) {
	channel := make(chan string)

	go func() {
		for i := 1; i <= 10; i++ {
			channel <- "Perulangan Ke " + strconv.Itoa(i)
		}
		close(channel)
	}()

	for x := range channel {
		fmt.Println("Menerima Data", x)
	}

	fmt.Println("Selesai")
}
