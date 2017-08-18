package main

import "fmt"

func test (x, y int, z float32) int {
	return x + y + int(z-0.5)
}

func main () {
	fmt.Println(test(50, 40, 43.5))
}