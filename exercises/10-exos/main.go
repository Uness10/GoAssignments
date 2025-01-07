package main

import "fmt"

func area(len float64, wid float64) (float64, bool) {
	if len <= 0 || wid <= 0 {
		return 0, false
	}
	return len * wid, true
}

func clos() func() int {
	i := 0
	return func() int {
		i++
		fmt.Println(&i)
		return i
	}

}
func main() {
	inc := clos()
	fmt.Println(clos)

	fmt.Println(inc())

	fmt.Println(inc())
}
