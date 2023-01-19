package main

import (
	"context"
	"fmt"
	"sync"
	"time"
)

type Future interface {
	GetResult() string // HL
}

type InnerFuture struct {
	once sync.Once // HL
	wg   sync.WaitGroup

	res        string
	resultChan <-chan string // HL
}

func (f *InnerFuture) GetResult() (result string) {
	f.once.Do(func() { // HL
		f.wg.Add(1)
		defer f.wg.Done()
		f.res = <-f.resultChan // HL
	})
	f.wg.Wait()
	return f.res
}

// START OMIT
func SlowFunction(ctx context.Context) Future {
	resultChan := make(chan string) // HL
	go func() {
		select {
		case <-time.After(time.Second * 2):
			resultChan <- "Secret message with delay" // HL
		case <-ctx.Done():
			resultChan <- ""
		}
	}()
	return &InnerFuture{resultChan: resultChan} // HL
}

func main() {

	ctx := context.Background()
	future := SlowFunction(ctx)
	res := future.GetResult() // HL

	fmt.Println(res) // HL
}

// END OMIT
