package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	in := bufio.NewScanner(os.Stdin)
	// alreadySeen := make(map[string]bool)
	var prev string
	for in.Scan() {
		txt := in.Text()

		if txt == prev {
			continue
		}

		if txt < prev {
			panic("File not sorted")
		}

		prev = txt

		// if _, found := alreadySeen[txt]; found {
		// 	continue
		// }

		// alreadySeen[txt] = true

		fmt.Println(txt)
	}

}
