package main

import "fmt"

func main() {

	//stringStream := make(chan string)
	intStream := make(chan int)

	go func() {
		//stringStream <- "hello world"
		defer close(intStream)
		for i := 1; i < 5; i++ {
			intStream <- i
		}
	}()

	//salution, ok := <-stringStream
	for integer := range intStream {

		fmt.Printf("(%v)", integer)
	}

}
