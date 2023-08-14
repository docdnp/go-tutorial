package main

import (
	"fmt"
	"time"
)

var intchan = make(chan int)

func producer() {
	for i := 0; i < 10; i++ {
		intchan <- i
		time.Sleep(100 * time.Millisecond)
	}
}
func consumer() {
	for {
		c := <-intchan
		fmt.Print(c, " ")
	}
}

func main() {
	go producer()
	consumer()
}
