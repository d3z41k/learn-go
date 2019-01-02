package main

import (
	"time"
	"fmt"
	"sync/atomic"
)

var totalOperations int32 = 0

func inc() {
	//totalOperations++
	atomic.AddInt32(&totalOperations, 1)
}

func main() {
	for i := 0; i < 1000; i++ {
		go inc()
	}
	time.Sleep(2 * time.Millisecond)
	// ожидаештся 1000 но по факту будет меньше
	fmt.Println("total operation: ", totalOperations)
}