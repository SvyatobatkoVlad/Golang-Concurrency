package main

import "fmt"

func main() {

	/*
		var dataStream <-chan interface{}
		dataStream := make(<-chan interface{})


		var dataStream chan<- interface{}
		dataStream := make(chan<- interface{})
	*/


	chanOwner := func() <-chan int { resultStream := make(chan int, 5) go func() {
		defer close(resultStream)
		for i := 0; i <= 5; i++ { resultStream <- i
		} }()
		return resultStream }
	resultStream := chanOwner()
	for result := range resultStream {
		fmt.Printf("Received: %d\n", result)
	}
	fmt.Println("Done receiving!")

}
