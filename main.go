// _Channels_ are the pipes that connect concurrent
// goroutines. You can send values into channels from one
// goroutine and receive those values into another
// goroutine.
// Example from: https://gobyexample.com/channels

package main

import (
	"fmt"
	"time"
)

func main() {

	// Create a new channel with `make(chan val-type)`.
	// Channels are typed by the values they convey.
	messages := make(chan string)

	// _Send_ a value into a channel using the `channel <-`
	// syntax. Here we send `"ping"`  to the `messages`
	// channel we made above, from a new goroutine.
	go func() {
		messages <- "ping"
		time.Sleep(time.Second * 2)
		messages <- "We had our anonymous function complete!"
	}()

	// The `<-channel` syntax _receives_ a value from the
	// channel. Here we'll receive the `"ping"` message
	// we sent above and print it out.
	msg := <-messages

	fmt.Println(msg)
	fmt.Scanln()

	// Here time has elapsed and the messages channel has a new
	// value in it. So we ask to receive the latest message
	// and print it out.
	msg = <-messages
	fmt.Println(msg)
}
