package main

import (
	"fmt"
	"math/rand/v2"
)

func part1() {
	rg := 20
	n := rand.IntN(rg) + 1
	i := 1
	fmt.Println("Let's start out game !, the range is ", rg)

	trying := true
	for trying {
		var g int
		fmt.Print("Enter your guess : ")
		fmt.Scanln(&g)
		if 1 > g || g > rg {
			fmt.Println("Out of range")
			continue
		}

		switch {
		case g < n:
			fmt.Println("Too low")
		case g > n:
			fmt.Println("Too high")
		default:
			trying = false
			fmt.Println("Correct !, you took ", i, " tries")
		}

		i++
	}

}

func part2() {
	rg := 60

	fmt.Println("Let's start out game !, the range is ", rg)
	bestScore := -1

	replay := true
	for replay {
		n := rand.IntN(rg) + 1
		trying := true
		i := 1
		var limit int
		fmt.Print("Enter the limit of tries : ")
		fmt.Scanln(&limit)
		for trying && i <= limit {
			var g int
			fmt.Print("Enter your guess : ")
			fmt.Scanln(&g)
			if 1 > g || g > rg {
				fmt.Println("Out of range")
				continue
			}

			switch {
			case g < n:
				fmt.Println("Too low")
			case g > n:
				fmt.Println("Too high")
			default:
				trying = false
				fmt.Println("Correct !, you took ", i, " tries")
				if bestScore == -1 {
					bestScore = i
				} else {
					bestScore = min(bestScore, i)
					fmt.Println("$_ Your best score is : ", bestScore)
				}
			}

			i++
		}
		if trying {
			fmt.Println("You lost")
		}
		replay := 0
		fmt.Print("Do you want to play again ? (1 for yes, 0 for no) : ")
		fmt.Scanln(&replay)
		if replay == 0 {
			fmt.Println("Good bye !")
			break
		}
		x := 0
		fmt.Println("You can configure the range difficulty : ")
		fmt.Println("	choose 1 for easy : 50 \n	choose 2 for medium : 100 \n	choose 3 for hard : 200")
		fmt.Print("Difficulty : ")
		fmt.Scanln(&x)

		switch x {
		case 1:
			rg = 50
		case 2:
			rg = 100
		case 3:
			rg = 200
		default:
			fmt.Println("Invalid input, the range will be set to medium")
			rg = 50
		}

	}

}
func main() {
	part2()
	//part1()
}
