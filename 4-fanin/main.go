package main

import (
	"context"
	"fmt"
	"time"
)

func main() {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	c := fanIn(ctx, boring(ctx, "Dafi"), boring(ctx, "Zora"), boring(ctx, "Raha"))

	for i := 0; i < 25; i++ {
		fmt.Println(<-c)
	}
	cancel()

	fmt.Println("You both are awesome. But I have to leave now.")
	time.Sleep(500 * time.Millisecond)
}

func fanIn(ctx context.Context, channels ...<-chan string) <-chan string {
	fan_in_channel := make(chan string)

	for _, chan_input := range channels {
		go func(chan_value <-chan string) {
			for {
				select {
				case <-ctx.Done():
					return
				case msg, ok := <-chan_value:
					if !ok {
						return
					}
					// Read from the channel and send to fan_in_channel
					fan_in_channel <- msg
				}
			}
		}(chan_input)
	}

	return fan_in_channel
}

func boring(ctx context.Context, msg string) <-chan string {
	boring_channel := make(chan string)

	go func() {
		defer close(boring_channel)
		for i := 0; ; i++ {
			select {
			case <-ctx.Done():
				return
			case boring_channel <- fmt.Sprintf("Hi, from boring func %s %d", msg, i):
				time.Sleep(time.Second)
			}

			// time.Sleep(time.Duration(rand.Intn(1e3)) * time.Millisecond)
		}
	}()

	return boring_channel
}
