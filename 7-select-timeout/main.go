package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	c := boring("John")

	timeout := time.After(2 * time.Second)

	for {
		select {
		case s := <-c:
			fmt.Println(s)

		case <-timeout:
			fmt.Println("no response, you talk too slow")
			return
		}

	}
}

func boring(msg string) <-chan string {
	boringChan := make(chan string)

	go func() {
		for i := 0; ; i++ {
			boringChan <- fmt.Sprintf("%s %d", msg, i)
			time.Sleep(time.Duration(rand.Intn(1500)) * time.Millisecond)
			// time.Sleep(time.Second)
		}
	}()
	return boringChan
}
