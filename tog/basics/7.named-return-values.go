package main

import "fmt"

func split(sum int) (x, y int, z string) {
	x = sum * 4 / 9
	y = sum - x
	z = "poe"
	// can override (but not partial)
	return 10, 2, "50"
}

func main() {
	fmt.Println(split(17))
}
