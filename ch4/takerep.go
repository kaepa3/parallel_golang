package main

import (
	"./common"
	"fmt"
)

func main() {
	done := make(chan interface{})
	defer close(done)
	for num := range common.Take(done, common.Repeat(done, 1), 10) {
		fmt.Printf("%v\n", num)
	}
}
