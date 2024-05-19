package main

import (
	"context"
	"errors"
	"fmt"
	"time"
)

func operation1(ctx context.Context) error {
	time.Sleep(100 * time.Millisecond)
	return errors.New("failed")
}

func operation2(ctx context.Context) {
	select {
	case <-time.After(500 * time.Millisecond):
		fmt.Println("Done")
	case <-ctx.Done():
		fmt.Println("halted operation2")
	}
}

func main() {
	ctx := context.Background()
	ctx, cancel := context.WithCancel(ctx)

	go func() {
		err := operation1(ctx)
		if err != nil {
			cancel()
		}
	}()

	operation2(ctx)
}
