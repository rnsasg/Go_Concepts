package main

import "fmt"

// Channels are the pipes that connect concurrent goroutines.
// You can send values into channels from one goroutine and receive those values into another goroutine.

// Create a new channel with make(chan val-type). Channels are typed by the values they convey.
// Send a value into a channel using the channel <- syntax.
// Here we send "ping" to the messages channel we made above, from a new goroutine.

// The <-channel syntax receives a value from the channel.
// Here we’ll receive the "ping" message we sent above and print it out.

func main() {

	ch := make(chan string)

	go func() { ch <- "ping" }()

	msg := <-ch
	fmt.Println(msg)

}
