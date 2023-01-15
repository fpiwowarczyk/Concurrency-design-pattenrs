package main

import (
	"context"
	"errors"
	"fmt"
	"time"
)

// START OMIT
type ToRetry func(context.Context) (string, error)

func Retry(toRetry ToRetry, retries int, delay time.Duration) ToRetry {
	return func(ctx context.Context) (string, error) {
		for r := 0; ; r++ {
			response, err := toRetry(ctx)   // HL
			if err == nil || r >= retries { //HL
				return response, err
			}
			fmt.Printf("Attempt %d failed; retrying in %v\n", r+1, delay) // HL
			select {
			case <-time.After(delay):
			case <-ctx.Done():
				return "", ctx.Err()
			}
		}
	}
}
func main() {
	r := Retry(FailThreeTimes, 5, time.Second) // HL
	res, err := r(context.Background())        // HL
	fmt.Println(res, err)                      // HL
}

// END OMIT

var count int

func FailThreeTimes(ctx context.Context) (string, error) {
	count++

	if count <= 3 {
		return "intentional fail", errors.New("error")
	} else {
		return "Success", nil
	}
}
