package main

import (
	"fmt"
	"time"
)

var intchan = make(chan int)

func producer() {
	i := 0
	for {
		i++
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
