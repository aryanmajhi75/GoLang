/*
Q4
You are tasked with creating a grade calculator program in GoLang.
The program should prompt the user to input their scores in three subjects:
Math, Science, and English. Based on these scores,
calculate the average grade (assuming each subject is equally weighted) and
display the corresponding letter grade (A, B, C, D, or F) using control flow.
*/

package main

import "fmt"

func main() {

	//variables
	var eng, sci, math, avg float32
	var grade string

	fmt.Println("Enter english marks")
	fmt.Scan(&eng)
	fmt.Println("Enter science marks")
	fmt.Scan(&sci)
	fmt.Println("Enter mathematics marks")
	fmt.Scan(&math)
	avg = (eng + sci + math) / 3
	if avg >= 90 {
		grade = "A"
	} else if avg >= 80 && avg < 90 {
		grade = "B"
	} else if avg >= 70 && avg < 80 {
		grade = "C"
	} else if avg >= 60 && avg < 70 {
		grade = "D"
	} else {
		grade = "F"
	}
	fmt.Println("\n-----------------------------------------")
	fmt.Println("| English | Mathematics | Science | Grade |")
	fmt.Printf("| %v      | %v          | %v      | %v    |", eng, math, sci, grade)
}