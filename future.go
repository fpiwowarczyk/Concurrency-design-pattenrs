package main

type Item struct {
	val string
}

func Fetch(nae string) <-chan Item {
	c := make(chan Item, 1)
	go func() {
		// ...
		c <- Item
	}()
	return c
}
