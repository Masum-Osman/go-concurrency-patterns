package main

import "fmt"

func main() {
	stageNames := []string{"Uploader", "Compressor", "Enricher", "Validator", "Parser"}
	var stages = len(stageNames)

	// Initial leftmost and rightmost channels
	leftmost := make(chan string)
	left := leftmost
	right := leftmost

	for i := 0; i < stages; i++ {
		right = make(chan string)
		go processStage(left, right, stageNames[i])
		left = right
	}

	// Start the chain with initial raw data
	go func(c chan string) {
		c <- "Raw Data"
	}(right)

	// Final processed output from the leftmost stage
	fmt.Println("Final Output: ", <-leftmost)
}

func processStage(left, right chan string, stageName string) {
	left <- (<-right + " -> " + stageName)
}
