package main

import "fmt"

func arrStuff() {
	arr := [5][5]int{}
	for i := 0; i < len(arr); i++ {
		for j := 0; j < len(arr[0]); j++ {
			arr[i][j] = 39888
			//fmt.Print(&arr[i][j], "\t")
		}
		//print("\n")
	}
	// print(len(arr))
	s := []int{1, 3, 4}
	s = append(s, s...)
	fmt.Println(s)
}
func freq(s []string) map[string]int {
	mp := map[string]int{}
	for _, val := range s {
		_, f := mp[val]
		if f {
			mp[val]++
		} else {
			mp[val] = 1
		}
	}
	return mp
}

func main() {
	s := []string{"ccc", "aa", "aa", "b", "bb", "b"}
	fmt.Println(freq(s))
}
