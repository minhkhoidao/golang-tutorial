package main

import "fmt"

func printNum() {
	for i := 0; i < 26; i++ {
		fmt.Printf("%d", i)
	}
}

func printChar() {
	for i := 'A'; i < 'A'+26; i++ {
		fmt.Printf("%c", i)
	}
}

func main() {
	go printChar()
	go printNum()
}
