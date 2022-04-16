package main

import "fmt"

type info struct {
	result int
}

func work() (int, error) {
	return 13, nil
}

func main() {
	var data info
	var err error

	data.result, err = work()
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Printf("%v\n", data)
}
