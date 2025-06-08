package main

import (
	"fmt"
	"time"
)

type Message struct {
	str  string
	wait chan bool
}

func main() {
	// Dafi and Zora wait for acknowledgment
	// Ahn does not wait
	c := fanIn(
		boring("Dafi", true),
		boring("Zora", true),
		boring("Ahn", false),
	)

	for i := 0; i < 25; i++ {
		msg := <-c
		fmt.Println("Received:", msg.str)

		// Only signal back if wait channel exists
		if msg.wait != nil {
			msg.wait <- true
		}
	}

	fmt.Println("Done. Mixed behavior demo complete.")
}

func boring(name string, shouldWait bool) <-chan Message {
	out := make(chan Message)
	var wait chan bool
	if shouldWait {
		wait = make(chan bool)
	}

	go func() {
		for i := 0; ; i++ {
			msg := Message{
				str:  fmt.Sprintf("[%s] message %d", name, i),
				wait: wait,
			}
			out <- msg

			time.Sleep(500 * time.Millisecond)

			// Wait only if applicable
			if shouldWait {
				<-wait
			}
		}
	}()

	return out
}

func fanIn(inputs ...<-chan Message) <-chan Message {
	out := make(chan Message)
	for _, ch := range inputs {
		go func(c <-chan Message) {
			for msg := range c {
				out <- msg
			}
		}(ch)
	}
	return out
}
