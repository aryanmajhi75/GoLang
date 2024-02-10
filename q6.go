/*
Q6
Develop a scenario based on your domain that incorporates looping, control structures, variables, and constants:
Calculate the factorial of a number.
*/

package main

import "fmt"

func main() {

	//variables
	var i, num int
	var fact int = 1

	fmt.Print("Show factorial of : ")
	fmt.Scan(&num)
	if num > 0 {
		for i = 1; i <= num; i++ {
			fact = fact * i
		}
	}

	fmt.Printf("\nThe factorial of %v is: %v\n", num, fact)
}
