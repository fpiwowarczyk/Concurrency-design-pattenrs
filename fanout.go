package main

import (
	"fmt"
	"sync"
	"time"
)

func Split(source <-chan string, n int) []<-chan string {
	dests := make([]<-chan string, 0)

	for i := 0; i <= n; i++ { // HL
		ch := make(chan string)
		dests = append(dests, ch)
		go func() { // HL
			defer close(ch)           // HL
			for val := range source { // HL
				ch <- val // HL
			} // HL
		}() // HL
	}
	return dests // HL
}

// START OMIT
func main() {
	source := splitMessage("Secret message goes in here") // HL
	dests := Split(source, 5)                             // HL
	var wg sync.WaitGroup
	wg.Add(len(dests))

	for i, ch := range dests {
		go func(i int, d <-chan string) {
			defer wg.Done()
			for val := range d {
				fmt.Printf("#%d got %s\n", i, val) // HL
			}
		}(i, ch)
	}
	wg.Wait()
}

// END OMIT

func splitMessage(msg string) <-chan string {
	c := make(chan string)
	go func() {
		defer close(c)
		for _, val := range msg {
			c <- string(val)
			time.Sleep(time.Millisecond * 100)
		}
	}()
	return c
}
