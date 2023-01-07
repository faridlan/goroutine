package test

import (
	"fmt"
	"testing"
	"time"
)

func GetNumber(number int) {
	fmt.Println("Number", number)
}

func TestNumber(t *testing.T) {
	for i := 0; i < 10000; i++ {
		go GetNumber(i)
	}
	time.Sleep(2 * time.Second)
}

func TestChannel(t *testing.T) {

	channel := make(chan string)
	defer close(channel)

	go func() {
		time.Sleep(2 * time.Second)
		channel <- "faridlan nul hakim"
		fmt.Println("Succes")
	}()

	data := <-channel
	fmt.Println(data)

	time.Sleep(5 * time.Second)

}
