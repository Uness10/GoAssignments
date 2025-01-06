package main

import "fmt"

func part1(n int) float32 {
	var sum float32
	for i := 0; i < n; i++ {
		var x float32
		fmt.Scanf("%f", &x)
		sum += x
	}
	average := sum / float32(n)
	return average
}

func part3() float32 {

	var sum float32
	var n int
	for {
		var x float32
		fmt.Printf("grade %d : ", 1+n)

		fmt.Scanln(&x)

		if x == -1 {
			break
		}
		sum += x
		n++

	}
	return sum / float32(n)
}
func part2(av float32) {
	switch {
	case av >= 90:
		fmt.Println("Grade: A")
	case av >= 80:
		fmt.Println("Grade: B")
	case av >= 70:
		fmt.Println("Grade: C")
	case av >= 60:
		fmt.Println("Grade: D")
	default:
		fmt.Println("Grade: F")
	}
}

func main() {
	average := part3()
	fmt.Printf("Average: %.2f\n", average)
	part2(average)
}
