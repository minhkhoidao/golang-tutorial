package main

import (
	"log"
)

func printChar(strchan chan string) {
	var result string
	for i := 'A'; i < 'A'+26; i++ {
		result += string(i)
	}
	strchan <- result
}

func printNumber(numchan chan int) {
	var result int

	for i := 0; i < 26; i++ {
		result += i
	}
	numchan <- result
}

func main() {
	chanPrintNumber := make(chan int)
	chanPrintChar := make(chan string)

	go printNumber(chanPrintNumber)
	go printChar(chanPrintChar)

	log.Println("channel: ", <-chanPrintNumber)
	log.Println("channel: ", <-chanPrintChar)

}
