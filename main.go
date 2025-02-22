package main

import (
	"fmt"

	"github.com/MikeJRalph/go-calc-basic/functions"
)

func main() {

	a := 10
	b := 6

	fmt.Println(functions.Add(a, b))
	fmt.Println(functions.Subtr(a, b))
	fmt.Println(float64(functions.Divide(a, b)))
	fmt.Println(functions.Multiply(a, b))
}
