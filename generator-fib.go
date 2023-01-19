package main

import (
	"fmt"
	"math/rand"
	"time"
)

// START OMIT
func main() {
	for number := range generateFibbonaci(10) { // HL
		fmt.Println(number)
	}
}
func generateFibbonaci(n int) <-chan int {
	c := make(chan int) // HL
	go func() {
		defer close(c) // HL
		a, b := 1, 1
		for i := 0; i <= n; i++ {
			if i == 0 || i == 1 {
				c <- a
				continue
			}
			b, a = a+b, b // HL
			time.Sleep(time.Duration(rand.Intn(1e3)) * time.Millisecond)
			c <- b // HL
		}
	}()
	return c // HL
}

// END OMIT
