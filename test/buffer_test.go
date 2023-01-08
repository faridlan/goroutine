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

func TestSelectChannel(t *testing.T) {
	channel1 := make(chan string)
	channel2 := make(chan string)
	defer close(channel1)
	defer close(channel2)

	go GiveMeResponse(channel1)
	go GiveMeResponse(channel2)

	counter := 0
	for {
		select {
		case data := <-channel1:
			fmt.Println("Data dari channel 1", data)
			counter++
		case data := <-channel2:
			fmt.Println("Data dari channel 2", data)
			counter++
		}

		if counter == 2 {
			break
		}
	}
}

func TestDefaultSelect(t *testing.T) {
	channel1 := make(chan string)
	channel2 := make(chan string)
	defer close(channel1)
	defer close(channel2)

	go GiveMeResponse(channel1)
	go GiveMeResponse(channel2)

	counter := 0
	for {
		select {
		case data := <-channel1:
			fmt.Println("Data dari channel 1", data)
			counter++
		case data := <-channel2:
			fmt.Println("Data dari channel 2", data)
			counter++
		default:
			fmt.Println("Menunggu Data")
		}

		if counter == 2 {
			break
		}
	}
}

func TestRaceCondition(t *testing.T) {

	x := 0
	for i := 1; i <= 1000; i++ {
		go func() {
			for j := 1; j <= 100; j++ {
				x += 1
			}
		}()
	}

	time.Sleep(5 * time.Second)
	fmt.Println("Perulang Ke", x)
}
