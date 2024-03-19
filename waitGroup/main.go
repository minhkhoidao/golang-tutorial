package main

import (
	"fmt"
	"sync"
)

func printNum(wg *sync.WaitGroup) {
	for i := 0; i < 26; i++ {
		fmt.Printf("%d", i)
	}
	wg.Done()
}

func printChar(wg *sync.WaitGroup) {
	for i := 'A'; i < 'A'+26; i++ {
		fmt.Printf("%c", i)
	}
	wg.Done()
}

func main() {
	var wg sync.WaitGroup
	wg.Add(2)

	go printChar(&wg)
	go printNum(&wg)

	wg.Wait()

}
