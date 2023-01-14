package main

import "fmt"

// START OMIT
func main() {
	for number := range generateFibbonaci(10) {
		fmt.Println(number)
	}
}
func generateFibbonaci(n int) <-chan int {
	c := make(chan int)
	defer close(c)
	go func() {
		a, b := 1, 1
		for i := 0; i <= n; i++ {
			if i == 0 || i == 1 {
				c <- a
				continue
			}
			b, a = a+b, b
			c <- b
		}
		close(c)
	}()
	return c
}

// END OMIT
