/*
Q3
You're developing a weather monitoring system in GoLang for a research institute.
The system needs to store data about temperature, humidity, and wind speed.
Define variables to hold these measurements for a specific location at a given point in time.
Ensure you choose suitable data types for representing numerical measurements accurately.
*/

package main

import "fmt"

type WeatherInfo struct {
	Location string
	Temp     float64
	Humid    float64
	WSpeed   float64
}

func main() {
	//variables
	var p [100]WeatherInfo
	var i, n int
	var ch string

	fmt.Println("Enter the number of locations ")
	fmt.Scan(&n)

	for i = 0; i < n; i++ {
		fmt.Println("\n-----------------------------")
		fmt.Println(i + 1)
		fmt.Println("Enter location ")
		fmt.Scan(&p[i].Location)	
		fmt.Println("Enter temperature ")
		fmt.Scan(&p[i].Temp)
		fmt.Println("Enter humidity ")
		fmt.Scan(&p[i].Humid)
		fmt.Println("Enter wind speed ")
		fmt.Scan(&p[i].WSpeed)
	}

	fmt.Println("\nDo you want to show the student details? (y/n)")
	fmt.Scan(&ch)

	if ch == "Y" || ch == "y" {
		for i = 0; i < n; i++ {
			fmt.Println("\n-----------------------------------------")
			fmt.Println("Location : ", p[i].Location)
			fmt.Println("|Temperature | Humidity | Wind Speed |")
			fmt.Printf("| %v         | %v       | %v         |", p[i].Temp, p[i].Humid, p[i].WSpeed)
		}
	} else {
		fmt.Println("\nBye..")
	}
	fmt.Println("\nGood Bye..")
}
