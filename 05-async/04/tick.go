package main

import (
	"time"
	"fmt"
)

func main() {
	ticker := time.NewTicker(time.Second)
	i := 0
	for tickTime := range ticker.C {
		i++
		fmt.Println("step", i, "time", tickTime)
		if i >= 5 {
			// надо остановить
			ticker.Stop()
			break
		}

		// не может быть остановлен и собран сборщиком мусора
		// используйте если должен работать вечно
		c := time.Tick(time.Second)
		i = 0
		for tickTime := range c {
			i++
			fmt.Println("step", i, "time", tickTime)
			if i >= 5 {
				break
			}
		}
	}
}
