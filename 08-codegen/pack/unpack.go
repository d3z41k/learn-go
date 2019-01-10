// go build gen/* && ./codegen pack/unpack.go  pack/marshaller.go
package main

import "fmt"

// lets generate code for this struct
// cgen: binpack
type User struct {
	ID       int
	RealName string `cgen:"-"`
	Login    string
	Flags    int
}

type Avatar struct {
	ID  int
	Url string
}

func main() {

	//perl -E '$b = pack("L L/a* L", 1_123_456, "den", 16); print map { ord.", "  } split("", $b); '

	data := []byte{128, 36, 17, 0, 3, 0, 0, 0, 100, 101, 110, 16, 0, 0, 0,}

	u := User{}
	u.Unpack(data)
	fmt.Printf("Unpacked user %#v", u)
}
