package main

import "fmt"

func calcArea(x int, y int) int {
	return x * y
}
func meter2ToFoot2(m int) float32 {
	return float32(m) * 10.7639
}
func main() {
	var x int = 10
	var y int = 8
	fmt.Println("area in meters: ", calcArea(x, y))
	fmt.Printf("area in feet: %.2f", meter2ToFoot2(calcArea(x, y)))
}
