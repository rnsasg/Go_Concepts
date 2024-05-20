package main

import (
	"fmt"
	"os"
	"os/exec"
)

// we run the ping command, we get continuous output at periodic intervals:
func main() {
	cmd := exec.Command("ping", "google.com")
	cmd.Stdout = os.Stdout

	if err := cmd.Run(); err != nil {
		fmt.Println("could not run command", err)
	}
}

// $ Output
// $ go run long_running_cmd.go
// PING google.com (142.250.195.174): 56 data bytes
// 64 bytes from 142.250.195.174: icmp_seq=0 ttl=116 time=8.172 ms
// 64 bytes from 142.250.195.174: icmp_seq=1 ttl=116 time=8.172 ms
// 64 bytes from 142.250.195.174: icmp_seq=2 ttl=116 time=8.484 ms
// 64 bytes from 142.250.195.174: icmp_seq=3 ttl=116 time=14.662 ms
