/*
Utilize the principles of Variables, Control Flow, Arrays, Slices, Maps, and Structs within a specified domain of your choice.
Begin the program with a comprehensive overview detailing the scenario and the concepts being implemented.
Ensure adequate comments are provided throughout the code. Evaluation of this program will be based on the following criteria:

R1: Concept Clarity/Viva: 8 Marks
R2: Correctness: 8 Marks
R3: Validations( only with if statement expected): 8 Marks
R4: Ability to Relate to Real-Time Scenario: 8 Marks
R5: Complexity: 8 Marks
*/

package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

type client struct {
	name          string
	age           int
	job           string
	numberofbanks int
	finances      [10]finance
}

type finance struct {
	balance     float64
	bankName    string
	accountType string
}

func readInput() string {
	// to manage problem with input with spaces
	reader := bufio.NewReader(os.Stdin)
	input, _ := reader.ReadString('\n')
	return strings.TrimSpace(input)
}

func menu() {
	fmt.Printf("\n")
	fmt.Println(" [1] Add Details ")
	fmt.Println(" [2] Add Another bank ")
	fmt.Println(" [3] Display Details ")
	fmt.Println(" [4] Update Balance ")
	fmt.Println(" [5] Exit ")
}

func main() {
	// variables
	var choice, ran int
	var search string
	var c client
	var newBalance float64

	// function to show menu
	for { // to mimic do while loop
		menu()
		fmt.Print("\nEnter your choice (1/2/3/4/5) : ")
		fmt.Scan(&choice)

		// switch case for the menu
		switch choice {
		case 1:
			if c.name == "" {
				fmt.Println("Enter your name : ")
				c.name = readInput()
				for {
					fmt.Println("Enter your age : ")
					fmt.Scanln(&c.age)
					if c.age > 0 {
						break
					}
				}
				fmt.Println("Describe your job role : ")
				c.job = readInput()
				fmt.Println("Enter your Bank name : ")
				c.finances[0].bankName = readInput()
				for {
					fmt.Println("Enter your current balance : ")
					fmt.Scanln(&c.finances[0].balance)
					if c.finances[0].balance > 0 {
						break
					}
				}
				fmt.Println("Enter your Account Type : ")
				c.finances[0].accountType = readInput()
				c.numberofbanks++
			} else {
				fmt.Println("Your data is already present!")
			}

		case 2:

			if c.numberofbanks < 10 {
				fmt.Println("Enter your Bank name : ")
				fmt.Scanln(&c.finances[c.numberofbanks].bankName)
				fmt.Println("Enter your current balance : ")
				fmt.Scanln(&c.finances[c.numberofbanks].balance)
				fmt.Println("Enter your Account Type : ")
				fmt.Scanln(&c.finances[c.numberofbanks].accountType)
				c.numberofbanks++
			} else {
				fmt.Println("You can add only up to 10 banks.")
			}

		case 3:
			fmt.Println("--------------------------------")
			fmt.Printf("\n %v", c.name)
			fmt.Printf("\n \t-%v", c.age)
			fmt.Printf("\n \t-%v", c.job)
			fmt.Printf("\n * You have %v banks entries\n", c.numberofbanks)
			fmt.Println("\nEnter the name of bank name (press d for default) :")
			fmt.Scanln(&search)
			fmt.Println("\n| Bank Name | Account Type | Balance |")
			fmt.Printf("========================================\n")
			if search == "d" || search == "D" {
				sliced := c.finances[0:3]
				if c.numberofbanks >= 3 {
					ran = 3
				} else {
					ran = c.numberofbanks
				}
				for i := 0; i < ran; i++ {
					fmt.Printf("| %v | %v | %v |\n", sliced[i].bankName, sliced[i].accountType, sliced[i].balance)
				}
			} else {
				for i := 0; i < c.numberofbanks; i++ {
					if c.finances[i].bankName == search {
						fmt.Printf("\n %v", c.finances[i].bankName)
						fmt.Printf("\n %v", c.finances[i].balance)
						fmt.Printf("\n %v", c.finances[i].accountType)
					}
				}
			}

		case 4:
			fmt.Println("\nEnter bank name  :")
			search = readInput()
			for i := 0; i < c.numberofbanks; i++ {
				if c.finances[i].bankName == search {
					fmt.Println("\nEnter Balance : ")
					fmt.Scan(&newBalance)
					c.finances[i].balance = newBalance
				}
			}
		case 5:
			return
		}
	}
}
