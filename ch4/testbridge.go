package main

import (
	"fmt"

	"github.com/kaepa3/parallel_golang/ch4/common"
)

func main() {
	genVals := func() <-chan <-chan interface{} {
		chanSteram := make(chan (<-chan interface{}))

		go func() {
			defer close(chanSteram)
			for i := 0; i < 10; i++ {
				stream := make(chan interface{}, 1)
				stream <- i
				close(stream)
				chanSteram <- stream
			}
		}()
		return chanSteram
	}
	for v := range common.Bridge(nil, genVals()) {
		fmt.Printf("%v", v)
	}
}
