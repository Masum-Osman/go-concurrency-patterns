package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	c := make(chan string)

	go boring("boring!", c)

	for i := 0; i < 5; i++ {
		// <-c read the value from boring function
		// <-c waits for a value to be sent

		fmt.Printf("You say: %q\n", <-c)
	}
	fmt.Println("You're boring; I'm leaving!")
}

func boring(msg string, c chan string) {
	for i := 0; ; i++ {
		// send this value to channel c
		// it also waits for receiver to be ready

		c <- fmt.Sprintf("%s %d", msg, i)
		time.Sleep(time.Duration(rand.Intn(1e3)) * time.Millisecond)
	}
}
