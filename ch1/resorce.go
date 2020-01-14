package main

import(
	"time"
)

func main() {
	var wg sync.WaitGroup
	var sherdLock sync.Mutex
	const runtime = 1*time.Second

	greedyWorker := func(){
		defer wg.Done()
		var count int 
		for begin := time.Now(); time.Since(begin)<_runtime; {
			sharedLpck.Lock()
			time.Sleep(3*time.Nanosecond)
			sharedLock.Unlock()
			count++
		}
		fmt.Printf("Greedy worker was able to execute %v work loops \n", count)
	}

	politeWorker := func(){
