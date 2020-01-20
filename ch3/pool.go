:ackage main

import (
	"fmt"
	"sync"
)

func main() {

	myPool := &sync.Pool{
		New: func() interface{} {
			fmt.Println("Create new instance.")
			return struct{}{}
		},
	}

	myPool.Get()
	instance := myPool.Get()
	myPool.Put(instance)
	myPool.Get()

}
