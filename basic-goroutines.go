package main

import (
	"fmt"
	"time"
)

// START OMIT
func main() {

	for i := 0; i <= 20; i++ {
		go printNumber(i) // HL
	}

	time.Sleep(time.Second)
}

func printNumber(n int) {
	fmt.Println(n)
}

// END OMIT
