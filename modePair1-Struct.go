package main

import "fmt"

type chambers struct {
	artifact string
	puzzle   string
	solved   bool
}

func exploreTemple(chambers) {
	fmt.Println("\tSolved puzzle")
}

func main() {
	var c [10]chambers
	var i, n int
	fmt.Print("\nEnter the number of chambers discovered : ")
	fmt.Scan(&n)
	for i = 0; i < n; i++ {
		fmt.Println("\n+-------------------------------")
		fmt.Print("\n| Artifact: ")
		fmt.Scanln(&c[i].artifact)
		fmt.Print("\n| Puzzle :")
		fmt.Scanln(&c[i].puzzle)
		c[i].solved = false
		fmt.Println("\n+-------------------------------")
	}

	for i = 0; i < n; i++ {
		fmt.Print("Solving puzzle for chamber %v", i)
		exploreTemple(c[i])
		c[i].solved = true
	}
	fmt.Println("\nReport of the chambers : ")
	fmt.Println("\n+-----------------------------------------------------------------------------------------+")
	fmt.Println("\n|      Artifact Description       |           Puzzle for the Artifact           | Solved? |")
	for i = 0; i < n; i++ {
		fmt.Print("\n| %v | %v | %v |", c[i].artifact, c[i].puzzle, c[i].solved)
	}
	fmt.Println("\n-----------------+")
}
