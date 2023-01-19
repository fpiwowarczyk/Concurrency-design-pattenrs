package main

import (
	"fmt"
	"math/rand"
	"time"
)

func init() {
	rand.Seed(time.Now().UnixNano())
}

// START OMIT
func doSlowRequest(output chan<- string) {
	wait := rand.Intn(1000)
	time.Sleep(time.Millisecond * time.Duration(wait))
	output <- "Super secret message" // HL
}
func main() {
	c := make(chan string) // HL
	go doSlowRequest(c)    // HL
	select {
	case message := <-c: // HL
		fmt.Printf("Received message: %s!\n", message)
	case <-time.After(time.Millisecond * 500): // HL
		fmt.Println("TIMEOUT")
	}
}

// END OMIT
