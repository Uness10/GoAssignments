package main

import "fmt"

const m2f = 10.7639

func calcArea(x float32, y float32) float32 {
	return x * y
}
func meter2ToFoot2(m float32) float32 {
	return m * m2f
}
func main() {
	var x float32
	var y float32
	fmt.Scanf("%f %f", &x, &y)
	var s float32 = calcArea(x, y)
	println(&x, &y)
	fmt.Printf("area in meters: %.2f \n", s)
	fmt.Printf("area in feet: %.2f", meter2ToFoot2(s))
}
