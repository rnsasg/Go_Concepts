package main

import (
	"context"
	"log"
	"math/rand"
	"time"
)

// we need to set a key that tells us where the data is stored
const keyID = "id"

func main() {
	// create a request ID as a random number
	rand.Seed(time.Now().Unix())
	requestID := rand.Intn(1000)

	ctx := context.WithValue(context.Background(), keyID, requestID)
	operation1(ctx)
}

func operation1(ctx context.Context) {
	log.Println("opration1 for id :", ctx.Value(keyID), " Completed")
	operation2(ctx)
}

func operation2(ctx context.Context) {
	log.Println("opration2 for id:", ctx.Value(keyID), " Completed")
}
