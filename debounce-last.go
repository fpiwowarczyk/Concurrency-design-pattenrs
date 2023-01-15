package main

import (
	"fmt"
	"time"
)

// START OMIT
func debounceLast(interval time.Duration, input chan string, callback func(arg string)) {
	var item string
	timer := time.NewTimer(interval) // HL
	for {
		select {
		case item = <-input:
			timer.Reset(interval) // HL
		case <-timer.C: // HL
			callback(item) // HL
		}
	}
}

func main() {
	messagesChan := make(chan string)
	go debounceLast(time.Second, messagesChan, echoMessage) // HL
	for true {
		// Will it print "message 1" or "message 2"?
		messagesChan <- "message 1"
		messagesChan <- "message 2"
		time.Sleep(time.Second)
	}
}

// END OMIT

func echoMessage(msg string) {
	fmt.Println(msg)
}
