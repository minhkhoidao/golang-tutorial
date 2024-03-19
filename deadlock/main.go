package main

import "fmt"

func dummy(dummyChan chan int) {
	dummyChan <- 10
}

func main() {
	chanInt := make(chan int)
	go dummy(chanInt)
	fmt.Println(<-chanInt)
}
