package main

import (
	"fmt"
	"math/rand"
	"time"
)

// START OMIT

func main() {
	c := generateMessage("generator is interesing pattern")
	for i := 0; i < 5; i++ {
		fmt.Printf("You say: %q\n", <-c)
	}

}

func generateMessage(msg string) <-chan string {
	c := make(chan string)
	defer close(c)
	go func() {
		for i := 0; ; i++ {
			c <- fmt.Sprintf("%s %d", msg, i)
			time.Sleep(time.Duration(rand.Intn(1e3)) * time.Millisecond)
		}
	}()
	return c
}

// END OMIT
