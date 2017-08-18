package main

import "fmt"

func multiple (x, y string) (string, string, string) {
	return y, x, "arm"
}

func main () {
	fmt.Println(multiple("amd", "intel"))
}