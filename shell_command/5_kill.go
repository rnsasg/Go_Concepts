package main

import (
	"context"
	"fmt"
	"os/exec"
	"time"
)

func main() {
	ctx := context.Background()
	// The context now times out after 1 second
	// alternately, we can call `cancel()` to terminate immediately
	ctx, cancel = context.WithTimeout(ctx, 1*time.Second)

	cmd := exec.CommandContext(ctx, "sleep", "100")

	out, err := cmd.Output()
	if err != nil {
		fmt.Println("could not run command: ", err)
	}
	fmt.Println("Output: ", string(out))

}
