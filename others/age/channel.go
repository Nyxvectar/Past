package main

import (
	"fmt"
)

func main() {
	ch := make(chan int)
	ch <- 1
	close(ch)
	fmt.Println(len(ch))
}
