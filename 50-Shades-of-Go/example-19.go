package main

func main() {
	x := []int{
		1,
		2,
	}
	x = x

	y := []int{3, 4} //no error
	y = y
}
