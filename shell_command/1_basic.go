package main

import (
	"fmt"
	"os/exec"
)

// let’s list the files in our current directory using ls, and print the output from our code:
func main() {
	cmd := exec.Command("ls", "./")
	out, err := cmd.Output()
	if err != nil {
		fmt.Println("Could not run command", err)
	}
	fmt.Println("Output", string(out))
}

// $ Output :
// $ go run basic.go
// Output README.md
// basic.go
