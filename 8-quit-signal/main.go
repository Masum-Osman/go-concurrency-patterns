package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	quitSignal := make(chan string)
	rcvChannel := boring("John", quitSignal)

	for i := 5; i >= 0; i-- {
		fmt.Println(<-rcvChannel)
	}

	quitSignal <- "bye"
	fmt.Println("John said: ", <-quitSignal)
}

func boring(msg string, signal chan string) <-chan string {
	boringChan := make(chan string)

	go func() {
		for i := 0; ; i++ {
			select {
			case boringChan <- fmt.Sprintf("%s %d", msg, i):

			case <-signal:
				fmt.Println("Cleaning up...")
				signal <- "Goodbye!"
				return
			}
			time.Sleep(time.Duration(rand.Intn(1e3)) * time.Millisecond)
		}
	}()

	return boringChan
}
