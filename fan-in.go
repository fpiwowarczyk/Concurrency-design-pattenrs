package main

import (
	"fmt"
	"sync"
	"time"
)

func Merge(sources ...<-chan int) <-chan int { // HL
	dest := make(chan int)
	var wg sync.WaitGroup
	wg.Add(len(sources))
	for _, ch := range sources {
		go func(c <-chan int) { // HL
			defer wg.Done()    // HL
			for n := range c { // HL
				dest <- n // HL
			} // HL
		}(ch) // HL
	}
	go func() {
		wg.Wait()
		close(dest)
	}()
	return dest
}

// START OMIT
func main() {
	sources := make([]<-chan int, 0) // HL
	for i := 2; i <= 4; i++ {
		source := generateNumberToThePower(i, 5)
		sources = append(sources, source)
	}
	dest := Merge(sources...) // HL
	for d := range dest {
		fmt.Print(d, " ")
	}
}

// END OMIT
func generateNumberToThePower(number, powers int) <-chan int {
	c := make(chan int)
	go func() {
		defer close(c)
		output := number
		for i := 1; i <= powers; i++ {
			time.Sleep(time.Second)
			output = number * output
			c <- output
		}
	}()
	return c
}
