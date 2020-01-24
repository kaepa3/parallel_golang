package main

import (
	"fmt"
	"github.com/kaepa3/parallel_golang/ch4/common/or"
	"time"
)

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
	<-or.Or(
		sig(2*time.Hour),
		sig(2*time.Minute),
		sig(1*time.Second),
		sig(1*time.Hour),
		sig(1*time.Minute),
	)

	fmt.Println("done afet %v", time.Since(start))
}
