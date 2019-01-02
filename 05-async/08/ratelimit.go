package main

import (
	"sync"
	"fmt"
	"strings"
	"time"
	"runtime"
)

const (
	iterationsNum = 7
	goroutinesNum = 5
	quotaLimit = 2
)


func startWorker(in int, wg *sync.WaitGroup, quotaCh chan struct{}) {
	quotaCh <- struct{}{} // берем свободный слот
	defer wg.Done()
	for j :=0; j < iterationsNum; j++ {
		fmt.Printf(formatWork(in, j))

		if j%2 == 0 {
			<-quotaCh  // возвращаем слот
			quotaCh <- struct{}{} // берем слот
		}

		runtime.Gosched()
	}
	<-quotaCh // возвращаем слот
}

func main() {
	wg := &sync.WaitGroup{} // инициализируем группу
	quotaCh := make(chan struct{}, quotaLimit)
	for i := 0; i < goroutinesNum; i++ {
		wg.Add(1) // добавляем воркер
		go startWorker(i, wg, quotaCh)
	}
	time.Sleep(time.Millisecond)

	wg.Wait() // ожидаем, пока waiter.Done() не приведет счетчик к 0
}


func formatWork(in, j int) string {
	return fmt.Sprintln(strings.Repeat("  ", in), "*",
		strings.Repeat("  ", goroutinesNum-in),
		"th", in,
		"iter", j, strings.Repeat("*", j))
}