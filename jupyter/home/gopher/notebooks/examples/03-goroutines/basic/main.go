package main

import (
	"fmt"
	"time"
)

func count() {
	for i := 0; i < 10; i++ {
		fmt.Print(i, " ")
		time.Sleep(100 * time.Millisecond)
	}
	fmt.Println()
}

func main() {
	go count()
	count()
}
