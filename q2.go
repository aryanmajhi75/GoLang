/*
Q2

	You're tasked with building a student information system in GoLang for a school.
	Each student record needs to store details such as student ID, name, age, and grade.
	Define variables to store the information of a single student and assign values accordingly.
	Pay attention to selecting appropriate data types to represent each piece of information.
*/
package main

import "fmt"

type StudentInfo struct {
	Id    int64
	Name  string
	Age   int
	Grade string
}

func main() {
	//variables
	var p [100]StudentInfo
	var i, n int
	var ch string

	fmt.Println("Enter the number of students ")
	fmt.Scan(&n)

	for i = 0; i < n; i++ {
		fmt.Println("\n-----------------------------")
		fmt.Println("Student ", i+1)
		fmt.Println("Enter student name ")
		fmt.Scan(&p[i].Name)
		fmt.Println("Enter student id ")
		fmt.Scan(&p[i].Id)
		fmt.Println("Enter student age ")
		fmt.Scan(&p[i].Age)
		fmt.Println("Enter student grade ")
		fmt.Scan(&p[i].Grade)
	}

	fmt.Println("\nDo you want to show the student details? (y/n)")
	fmt.Scan(&ch)

	if ch == "Y" || ch == "y" {
		for i = 0; i < n; i++ {
			fmt.Println("\n-----------------------------------------")
			fmt.Println("| Id : ", p[i].Id)
			fmt.Println("|Name | Age | Grade |")
			fmt.Printf("| %v  | %v  | %s    |", p[i].Name, p[i].Age, p[i].Grade)
		}
	} else {
		fmt.Println("\nBye..")
	}
	fmt.Println("\nGood Bye..")
}
