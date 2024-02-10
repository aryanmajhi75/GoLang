/*
Q5
You want to build a simple interest calculator in GoLang.
The program should ask the user to input multiple sets of
data, each containing the principal amount, the annual interest rate,
and the number of years for which the interest is to be calculated.
Use arrays to store the input data for each set,
calculate the simple interest for each set using the formula:
Simple Interest = (Principal Amount * Annual Interest Rate * Number of Years) / 100, and display the results.
*/

package main

import "fmt"

type SimpleInterest struct {
	principle float64
	rate      float64
	years     float64
	interest  float64
}

func main() {
	//variables
	var p [100]SimpleInterest
	var i, n int
	var ch string

	fmt.Println("Enter the number of customers ")
	fmt.Scan(&n)

	for i = 0; i < n; i++ {
		fmt.Println("\n-----------------------------")
		fmt.Println("Customer ", i+1)
		fmt.Println("Enter principle amount ")
		fmt.Scan(&p[i].principle)
		fmt.Println("Enter rate of interest ")
		fmt.Scan(&p[i].rate)
		fmt.Println("Enter number of years ")
		fmt.Scan(&p[i].years)
		p[i].interest = (p[i].principle * p[i].rate * p[i].years) / 100
	}

	fmt.Println("\nDo you want to show the customer details? (y/n)")
	fmt.Scan(&ch)

	if ch == "Y" || ch == "y" {
		for i = 0; i < n; i++ {
			fmt.Println("\n-----------------------------------------")
			fmt.Println("| Principle | Rate | Years | Simple Interest |")
			fmt.Printf("| %v        | %v   | %v    | %v               |", p[i].principle, p[i].rate, p[i].years, p[i].interest)
		}
	} else {
		fmt.Println("\nBye..")
	}
	fmt.Println("\nGood Bye..")
}