package main

import "fmt"

type addRequest struct {
	response chan int
	summands []int
}

func SumRespond(c chan addRequest) {
	result := 0
	req := <-c
	println("Sums: received request: ", req.summands)
	for _, s := range req.summands {
		result = result + s
	}
	req.response <- result
}

func SumRoutine() chan addRequest {
	c := make(chan addRequest)
	go SumRespond(c)
	return c
}

func SumRespondLoop(c chan addRequest) {
	for {
		println("Awaiting sum request")
		SumRespond(c)
		println("Next loop")
	}
}

func MakeSumRespondLoop() chan addRequest {
	c := make(chan addRequest)
	go SumRespondLoop(c)
	return c
}

func main() {
	sums := MakeSumRespondLoop()

	req1 := addRequest{make(chan int), []int{5, 5, 10, 80}}
	req2 := addRequest{make(chan int), []int{10, 5, 10, 75}}

	fmt.Printf("Requesting sum of: req1: %v\n", req1.summands)
	fmt.Printf("Requesting sum of: req2: %v\n", req2.summands)
	sums <- req1
	println("Requesting sum of: ", req2.summands)
	sums <- req2
	println("Received response: ", <-req1.response)
	println("Received response: ", <-req2.response)
}
