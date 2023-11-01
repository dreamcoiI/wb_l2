package main

import (
	"fmt"
	"time"
)

func or(ch ...<-chan interface{}) <-chan interface{} {
	res := make(chan interface{})

	for _, val := range ch {
		go func(channel <-chan interface{}) {
			<-channel
			close(res)
		}(val)
	}

	return res
}

func main() {
	sig := func(after time.Duration) <-chan interface{} {
		c := make(chan interface{})
		go func() {
			defer close(c)
			time.Sleep(after)
		}()
		return c
	}

	start := time.Now()
	<-or(
		sig(2*time.Hour),
		sig(5*time.Minute),
		sig(1*time.Second),
		sig(1*time.Hour),
		sig(1*time.Minute),
	)

	fmt.Printf("done after %v", time.Since(start))

}
