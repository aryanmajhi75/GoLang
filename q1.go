/*
Q1
 You're developing an online store application in GoLang.
 As part of the application, you need to keep track of various product details such as name,
 price, and quantity in stock. Design a set of variables and assign values to represent a
 specific product in the inventory. Ensure you use appropriate data types for each variable
 to accurately capture the information.
*/

package main

import "fmt"

type ProductCatalog struct {
	Name     string
	Price    float64
	Quantity int
}

func main() {

	//variables
	var p [100]ProductCatalog
	var i, n int
	var ch string

	fmt.Println("Enter the number of products ")
	fmt.Scan(&n)

	for i = 0; i < n; i++ {
		fmt.Println("\n-----------------------------")
		fmt.Println("Product ", i+1)
		fmt.Println("Enter product name ")
		fmt.Scan(&p[i].Name)
		fmt.Println("Enter product price ")
		fmt.Scan(&p[i].Price)
		fmt.Println("Enter product quantity ")
		fmt.Scan(&p[i].Quantity)
	}

	fmt.Println("\nDo you want to show the product details? (y/n)")
	fmt.Scan(&ch)

	if ch == "Y" || ch == "y" {
		for i = 0; i < n; i++ {
			fmt.Println("\n-----------------------------------------")
			fmt.Println("Product ", i)
			fmt.Println("| Name | Price | Quantity |")
			fmt.Printf("| %s   | %v    | %v       |", p[i].Name, p[i].Price, p[i].Quantity)
		}
	} else {
		fmt.Println("\nBye..")
	}
	fmt.Println("\nGood Bye..")
}
