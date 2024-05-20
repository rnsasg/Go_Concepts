package main

import (
	"fmt"
	"os/exec"
)

type CustomWriter struct{}

func (c CustomWriter) Write(p []byte) (int, error) {
	fmt.Println("Received Output: ", string(p))
	return len(p), nil
}

// we run the ping command, we get continuous output at periodic intervals:
func main() {
	cmd := exec.Command("ping", "google.com")
	cmd.Stdout = CustomWriter{}

	if err := cmd.Run(); err != nil {
		fmt.Println("could not run command", err)
	}
}

// $ Output
// $  go run 3_custom_writer.go
// Received Output:  PING google.com (142.250.195.142): 56 data bytes
// 64 bytes from 142.250.195.142: icmp_seq=0 ttl=116 time=7.881 ms

// Received Output:  64 bytes from 142.250.195.142: icmp_seq=1 ttl=116 time=14.254 ms

// Received Output:  64 bytes from 142.250.195.142: icmp_seq=2 ttl=116 time=14.717 ms

// Received Output:  64 bytes from 142.250.195.142: icmp_seq=3 ttl=116 time=7.882 ms
