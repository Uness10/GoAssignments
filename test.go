package main

import "fmt"

func main() {
	var x float32 = 10
	var _ int
	var y float32 = 8

	fmt.Println(x / y)
	fmt.Printf("%f\n", x/y)
}
