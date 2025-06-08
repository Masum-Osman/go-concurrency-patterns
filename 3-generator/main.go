package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	zaif := boring("Zaif")
	raka := boring("Raka")

	for msg := range zaif {
		fmt.Println(msg)
	}

	for msg := range raka {
		fmt.Println(msg)
	}

	fmt.Println("You both are awesome. But I have to leave now.")
}

func boring(msg string) <-chan string {
	c := make(chan string)

	go func() {
		for i := 0; i < 10; i++ {
			if i%3 == 0 {
				continue
			}
			c <- fmt.Sprintf("Hello from channel. This is %s %d", msg, i)
			time.Sleep(time.Duration(rand.Intn(1e3)) * time.Millisecond)
		}

		close(c)
	}()

	return c
}
