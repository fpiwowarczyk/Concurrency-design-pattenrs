package main

import (
	"context"
	"fmt"
	"sync"
	"time"
)

type Future interface {
	GetResult() (string, error)
}

type InnerFuture struct {
	once sync.Once
	wg   sync.WaitGroup

	res        string
	err        error
	resultChan <-chan string
	errChan    <-chan error
}

func (f *InnerFuture) GetResult() (result string, err error) {
	f.once.Do(func() {
		f.wg.Add(1)
		defer f.wg.Done()
		f.res = <-f.resultChan
		f.err = <-f.errChan
	})
	f.wg.Wait()
	return f.res, f.err
}

// START OMIT
func SlowFunction(ctx context.Context) Future {
	resultChan := make(chan string)
	errChan := make(chan error)
	go func() {
		select {
		case <-time.After(time.Second * 2):
			resultChan <- "Secret message with delay"
			errChan <- nil
		case <-ctx.Done():
			resultChan <- ""
			errChan <- ctx.Err()
		}
	}()
	return &InnerFuture{resultChan: resultChan, errChan: errChan}
}

func main() {
	ctx := context.Background()
	future := SlowFunction(ctx)
	res, err := future.GetResult()
	if err != nil {
		fmt.Println("error:", err)
		return
	}
	fmt.Println(res)
}

// END OMIT
