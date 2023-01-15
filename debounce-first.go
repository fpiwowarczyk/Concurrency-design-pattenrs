package main

import (
	"fmt"
	"time"
)

// START OMIT
func debounceFirst(interval time.Duration, input chan string, callback func(arg string)) {
	var item string
	timer := time.NewTimer(interval) // HL
	called := false                  // HL
	for {
		select {
		case item = <-input:
			if !called {
				callback(item)        // HL
				called = true         // HL
				timer.Reset(interval) // HL
			}
		case <-timer.C:
			called = false //HL
		}
	}
}
func main() {
	messagesChan := make(chan string)
	go debounceFirst(time.Second, messagesChan, echoMessage) // HL
	sendMessages(messagesChan)
}

// END OMIT

func echoMessage(msg string) {
	fmt.Println(msg)
}

func sendMessages(msgs chan string) {
	for true {
		msgs <- "message 1"
		msgs <- "message 2"
		time.Sleep(time.Second)
	}
}
