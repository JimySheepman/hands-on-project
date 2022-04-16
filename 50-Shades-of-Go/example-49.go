package main

import (
	"fmt"
)

func main() {
	i := 1
	defer func(in *int) { fmt.Println("result =>", *in) }(&i)

	i = 2
	//prints: result => 2
}
