package main

import (
	"fmt"

	"github.com/kaepa3/parallel_golang/ch4/common"
)

func main() {
	done := make(chan interface{})
	defer close(done)

	out1, out2 := common.Tee(done, common.Take(done, common.Repeat(done, 1, 2), 4))

	for vall := range out1 {
		fmt.Printf("out1: %v, out2: %v\n", vall, <-out2)
	}
}
