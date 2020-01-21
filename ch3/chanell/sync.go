package main

import (
	"fmt"
	"sync"
)

func main() {
	begin := make(chan interface{})
	var wg sync.WaitGroup
	for i := 0; i < 4; i++ {
		wg.Add(1)
		go func(i int) {
			defer wg.Done()
			<-begin
			fmt.Printf("%v has Begun\n", i)
		}(i)
	}
	fmt.Println("Unblocking goroutiens...")
	close(begin)
	wg.Wait()

}
