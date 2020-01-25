package main

import (
	"fmt"
	"math/rand"

	"github.com/kaepa3/parallel_golang/ch4/common"
)

func main() {
	done := make(chan interface{})
	defer close(done)

	rand := func() interface{} { return rand.Int() }
	for num := range common.Take(done, common.RepeatFn(done, rand), 4) {
		fmt.Println(num)
	}
}
