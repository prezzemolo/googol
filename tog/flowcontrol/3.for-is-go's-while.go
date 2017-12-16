package main

import "fmt"

func main() {
	sum := "go"
	for sum != "gogogogogo" {
		fmt.Println(sum)
		sum += "go"
	}
	fmt.Println(sum)
}
