package main

import (
	"fmt"
	"time"
)

// This is example of feed reader design pattern

func main() {
	//Subscribe to some feeds, and create a merger update stream 
	merged := Merge(
		Subscribe(Fetch("blog.golang.org")),
		Subscribe(Fetch("googleblog.blogspot.com")),
		Subscribe(Fetch("googledevelopers.blogspot.com"))
	)

	// Close the subscriptions after some time. 
	time.AfterFunc(3*time.Second, func() {
		fmt.Println("closed:", merged.Close()))
	})

	// Print the steam 
	for it := range merged.Updates() {
		fmt.Println(it.Channel, it.Title)
	}

	panic("show me the stacks")
}
