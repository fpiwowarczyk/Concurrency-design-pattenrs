package main

import (
	"fmt"
	"time"
)

// START OMIT

type Ball struct{ hits int } // HL

func main() {
	table := make(chan *Ball)
	go player("ping", table) // HL
	go player("pong", table) // HL

	table <- new(Ball) // HL
	time.Sleep(time.Second)
	<-table // HL
}

func player(name string, table chan *Ball) {
	for {
		ball := <-table // HL
		ball.hits++
		fmt.Println(name, ball.hits)
		time.Sleep(100 * time.Millisecond)
		table <- ball // HL
	}
}

// END OMIT
