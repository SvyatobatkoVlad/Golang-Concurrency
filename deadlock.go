package main

import "fmt"

func main() {
	//======================== simple example deadlock ========================
	ch := make(chan int)
	ch <- 1
	fmt.Println(<-ch)
}
