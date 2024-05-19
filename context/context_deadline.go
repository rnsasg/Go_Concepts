package main

import (
	"context"
	"fmt"
	"net/http"
	"time"
)

// The context will be cancelled after 3 seconds
// If it needs to be cancelled earlier, the `cancel` function can
// be used, like before
// ctx, cancel := context.WithTimeout(ctx, 3*time.Second)

// Setting a context deadline is similar to setting a timeout, except
// you specify a time when you want the context to cancel, rather than a duration.
// Here, the context will be cancelled on 2009-11-10 23:00:00
// ctx, cancel := context.WithDeadline(ctx, time.Date(2009, time.November, 10, 23, 0, 0, 0, time.UTC)

func main() {
	// Create a new context
	// With a deadline of 100 milliseconds
	ctx := context.Background()
	ctx, _ = context.WithTimeout(ctx, 100*time.Millisecond)

	// Make a request, that will call the google homepage
	req, _ := http.NewRequest(http.MethodGet, "https://www.sohamkamani.com/", nil)

	// Associate the cancellable context we just created to the request
	req = req.WithContext(ctx)

	client := http.Client{}
	res, err := client.Do(req)

	if err != nil {
		fmt.Println("Request Failed", err)
		return
	}

	fmt.Println("Response Recevived,Status Code", res.StatusCode)
}
