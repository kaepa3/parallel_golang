package main

import (
	"fmt"

	"./common"
)

func main() {
	done := make(chan interface{})

	defer close(done)
	var message string
	for token := range common.ToString(done, common.Take(done, common.Repeat(done, "I", "am", "you."), 5)) {
		message += token
	}
	fmt.Println(message)

}
