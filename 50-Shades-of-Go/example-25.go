package main

import "fmt"

func main() {
	var d uint8 = 2
	fmt.Printf("%08b\n", ^d)
}
