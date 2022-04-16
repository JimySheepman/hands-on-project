package main

import "fmt"

func main() {
	x := map[string]string{"one": "a", "two": "", "three": "c"}

	if _, ok := x["two"]; !ok {
		fmt.Println("no entry")
	}
}
