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
	// var c [10]chambers
	var i, n int
	fmt.Print("\nEnter the number of chambers discovered : ")
	fmt.Scan(&n)
	// for i = 0; i < n; i++ {
	// 	fmt.Println("\n+-------------------------------")
	// 	fmt.Print("\n| Artifact: ")
	// 	fmt.Scanln(&c[i].artifact)
	// 	fmt.Print("\n| Puzzle :")
	// 	fmt.Scanln(&c[i].puzzle)
	// 	c[i].solved = false
	// 	fmt.Println("\n+-------------------------------")
	// }
	c := []chambers{
		{Artifact: "Ancient Sword", Puzzle: "I speak without a mouth and hear without ears. I have no body, but I come alive with the wind. What am I?", Solved: false},
		{Artifact: "Golden Amulet", Puzzle: "The more you take, the more you leave behind. What am I?", Solved: false},
		{Artifact: "Crystal Ball", Puzzle: "I have keys but open no locks. I have space but no room. You can enter, but you can't go inside. What am I?", Solved: false},
		{Artifact: "Ultimate Treasure", Puzzle: "The key to this treasure lies in the heart of knowledge. Solve the final puzzle to reveal the ultimate secret.", Solved: false},
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
