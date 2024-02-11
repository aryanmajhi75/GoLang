package main

import "fmt"

type Chamber struct {
	Artifact string
	Puzzle   string
	Solved   bool
}

func exploreTemple(chambers []Chamber) string {
	for i := 0; i < len(chambers)-1; i++ {
		currentChamber := &chambers[i]

		if !currentChamber.Solved {
			fmt.Printf("Entering chamber %d...\n", i+1)
			fmt.Println("Artifact:", currentChamber.Artifact)
			fmt.Println("Puzzle:", currentChamber.Puzzle)

			currentChamber.Solved = true
			fmt.Printf("Puzzle in chamber %d solved!\n\n", i+1)
		} else {
			fmt.Printf("Chamber %d already solved. Moving to the next one.\n\n", i+1)
		}
	}

	return chambers[len(chambers)-1].Artifact
}

func main() {
	chambers := []Chamber{
		{Artifact: "Ancient Sword", Puzzle: "I speak without a mouth and hear without ears. I have no body, but I come alive with the wind. What am I?", Solved: false},
		{Artifact: "Golden Amulet", Puzzle: "The more you take, the more you leave behind. What am I?", Solved: false},
		{Artifact: "Crystal Ball", Puzzle: "I have keys but open no locks. I have space but no room. You can enter, but you can't go inside. What am I?", Solved: false},
		{Artifact: "Ultimate Treasure", Puzzle: "The key to this treasure lies in the heart of knowledge. Solve the final puzzle to reveal the ultimate secret.", Solved: false},
	}

	ultimateTreasure := exploreTemple(chambers)
	fmt.Println("The ultimate treasure is:", ultimateTreasure)
}
