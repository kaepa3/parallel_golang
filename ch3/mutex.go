package main

import (
	"fmt"
	"sync"
)

func main() {
	var count int
	var lock sync.Mutex

	increament := func() {
		lock.Lock()
		defer lock.Unlock()
		count++
		fmt.Printf("Increamenting: %d\n", count)
	}
	decrement := func() {
		lock.Lock()
		defer lock.Unlock()
		count--
		fmt.Printf("Decreamenting: %d\n", count)
	}

	var arithmetic sync.WaitGroup

	for i := 0; i <= 5; i++ {
		arithmetic.Add(1)
		go func() {
			defer arithmetic.Done()
			increament()
		}()
	}
	for i := 0; i <= 5; i++ {
		arithmetic.Add(1)
		go func() {
			defer arithmetic.Done()
			decrement()
		}()
	}
	arithmetic.Wait()

	fmt.Println("Arithmetic complete.")
}
