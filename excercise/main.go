package main

import (
	"log"
	"os"
	"strings"
)

// Count word occurrences in a file

func countKeywordFile(result chan int, filepath string, keyword string) {
	var numberOfOccurrences int

	fileContent, err := os.ReadFile(filepath)

	if err != nil {
		log.Println(err)
		result <- 0
		return
	}

	numberOfOccurrences = strings.Count(string(fileContent), keyword)

	result <- numberOfOccurrences
	defer close(result)
}

func main() {
	countFirstChan := make(chan int)
	countSecondChan := make(chan int)

	go countKeywordFile(countFirstChan, "text1.txt", "go")
	go countKeywordFile(countSecondChan, "text2.txt", "go")

	log.Println("Number of occurrences in file: ", <-countFirstChan)
}
