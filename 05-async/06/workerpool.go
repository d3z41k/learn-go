package main

import (
	"fmt"
	"runtime"
	"time"
	"strings"
)

const goroutinesNum = 3

func startWorker(workerNum int, in <-chan  string) {
	for input := range in {
		fmt.Printf(formatWork(workerNum, input))
		runtime.Gosched() //try to comment
	}
	printFinishWork(workerNum)
}

func main() {
	workerInput := make(chan string, 3) //try increase chanel size

	for i := 0; i < goroutinesNum ; i++  {
		go startWorker(i, workerInput)
	}

	months := []string{
		"Январь", "Февраль", "Март",
		"Апрель", "Май", "Июнь",
		"Июль", "Август", "Сентябрь",
		"Октябрь", "Ноябрь", "Декабрь"}

	for _, monthName := range months {
		workerInput <- monthName
	}
	close(workerInput) // try to comment

	time.Sleep(time.Millisecond)
}

func formatWork(in int, input string) string {
	return fmt.Sprintln(strings.Repeat("  ", in), "*",
		strings.Repeat("  ", goroutinesNum-in),
			"th", in,
			"recieved", input)
}

func printFinishWork(workerNum int) {
	fmt.Printf("=== %d finished\r\n", workerNum)
}
