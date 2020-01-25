package main

import (
	"fmt"
	"math/rand"
	"runtime"
	"time"

	"github.com/kaepa3/parallel_golang/ch4/common"
)

func main() {
	done := make(chan interface{})
	defer close(done)

	start := time.Now()
	rand := func() interface{} { return rand.Intn(50000000) }
	randIntStream := toInt(done, common.RepeatFn(done, rand))

	numFinders := runtime.NumCPU()
	fmt.Printf("spinning up %d prime finders\n", numFinders)
	finders := make([]<-chan interface{}, numFinders)

	fmt.Println("Primes:")
	for i := 0; i < numFinders; i++ {
		finders[i] = primeFinder(done, randIntStream)
	}
	for prime := range common.Take(done, common.FanIn(done, finders...), 10) {
		fmt.Printf("\t%d\n", prime)
	}
	fmt.Printf("Search took; %v", time.Since(start))
}
func primeFinder(done <-chan interface{}, intStream <-chan int) <-chan interface{} {
	primeStream := make(chan interface{})
	go func() {
		defer close(primeStream)
		for integer := range intStream {
			integer -= 1
			prime := true
			for divisor := integer - 1; divisor > 1; divisor-- {
				if integer%divisor == 0 {
					prime = false
					break
				}
			}

			if prime {
				select {
				case <-done:
					return
				case primeStream <- integer:
				}
			}
		}
	}()
	return primeStream
}

func toInt(done <-chan interface{}, val <-chan interface{}) <-chan int {
	integerStream := make(chan int)
	go func() {
		defer close(integerStream)
		for v := range val {
			select {
			case <-done:
				return
			case integerStream <- v.(int):
			}
		}
	}()
	return integerStream
}
