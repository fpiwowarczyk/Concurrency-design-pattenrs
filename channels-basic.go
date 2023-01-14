package main

import (
	"fmt"
	"math/rand"
	"time"
)

// START OMIT

func work(id, input int, output chan int) {
	fmt.Printf("Worker with id %d starts work\n", id)
	time.Sleep(500 * time.Millisecond)
	output <- input * input
}

func main() {

	c := make(chan int)
	for i := 0; i < 3; i++ {
		go work(i, rand.Intn(5), c) // HL
	}
	w1, w2, w3 := <-c, <-c, <-c // HL

	fmt.Println(w1, w2, w3)
}

//END OMIT
