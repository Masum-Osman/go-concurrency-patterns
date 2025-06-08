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
	c := fanIn(boring("Dafi"), boring("Zora"), boring("Ahn"))

	for i := 0; i < 25; i++ {
		msg1 := <-c
		fmt.Println("Msg One: ", msg1)

		msg2 := <-c
		fmt.Println("Msg Two: ", msg2)

		msg1.wait <- true
		msg2.wait <- true
	}

	fmt.Println("You both are awesome. But I have to leave now.")

}

func fanIn(inputs ...<-chan Message) <-chan Message {
	fanIn := make(chan Message)

	for i := range inputs {
		input := inputs[i]
		go func() {
			for {
				fanIn <- <-input
			}
		}()

	}
	return fanIn
}

func boring(msg string) <-chan Message {
	boringChan := make(chan Message)
	waitForIt := make(chan bool)

	go func() {
		for i := 0; ; i++ {
			boringChan <- Message{
				str:  fmt.Sprintf("Hi, from boring func %s %d", msg, i),
				wait: waitForIt,
			}
			time.Sleep(time.Second)
			<-waitForIt // Wait for the main goroutine to signal that it has processed the message
		}
	}()

	return boringChan
}
