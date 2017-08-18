// https://tour.golang.org/basics/1
// https://golang.org/pkg/math/rand/#Rand.Int
package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	fmt.Println(r.Int())
}
