package main

import (
	"fmt"
	"time"
)

func worker(id int, channel chan int) {
	for num := range channel {
		fmt.Printf("Worker %d received %d\n", id, num)
		time.Sleep(time.Millisecond * 10)
	}
}

func main() {
	channel := make(chan int)
	for i := 0; i < 3; i++ {
		go worker(i, channel)
	}
	for i := 0; i < 5; i++ {
		channel <- i
	}
	close(channel)
}
