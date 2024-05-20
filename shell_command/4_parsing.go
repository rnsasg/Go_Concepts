package main

import (
	"fmt"
	"io"
	"os"
	"os/exec"
)

func main() {
	cmd := exec.Command("grep", "apple")
	// Create a new pipe, which gives us a reader/writer pair
	reader, writer := io.Pipe()

	cmd.Stdin = reader
	cmd.Stdout = os.Stdout

	go func() {
		defer writer.Close()
		// the writer is connected to the reader via the pipe
		// so all data written here is passed on to the commands
		// standard input
		writer.Write([]byte("1. pear\n"))
		writer.Write([]byte("2. grapes\n"))
		writer.Write([]byte("3. apple\n"))
		writer.Write([]byte("4. banana\n"))
	}()

	if err := cmd.Run(); err != nil {
		fmt.Println("could not run command: ", err)
	}
}
