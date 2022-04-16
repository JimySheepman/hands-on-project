package main

import "fmt"

func main() {
	x := "text"
	fmt.Println(x[0])      //print 116
	fmt.Printf("%T", x[0]) //prints uint8
}
