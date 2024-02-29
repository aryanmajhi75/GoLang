/*
You're tasked with creating a program for a highly sensitive government agency that tracks top-secret missions.
The program must manage a collection of mission data, including mission names, objectives, and statuses.
However, there's a catch: the program must be designed to self-destruct (trigger a panic)
if anyone tries to access a mission labeled with a certain codeword.
This requirement is in place to prevent unauthorized access to missions of extreme importance.
Your task is to design and implement this program in Go, ensuring that:
Mission data is stored securely and can only be accessed through proper channels.
The program includes error handling mechanisms to gracefully handle any unexpected issues.
A panic is triggered if the user tries to access a mission labeled with the specified codeword,
but the program should also include a recovery mechanism to prevent the entire program from crashing.
Defer statements are used appropriately to clean up resources and ensure the program runs smoothly under normal conditions.
*/

package main

import (
	"bufio"
	"errors"
	"fmt"
	"os"
	"strings"
)

type mission struct {
	name       string
	objectives string
	status     string
}

func scanInput() string {
	reader := bufio.NewReader(os.Stdin)
	input, _ := reader.ReadString('\n')
	return strings.TrimSpace(input)
}

func showMenu(m []mission, n int) {
	fmt.Println("\nMISSIONS")
	fmt.Println("+", strings.Repeat("-", 30), "+")
	fmt.Println("| [0]       Exit", strings.Repeat(" ", 15), "|")
	fmt.Println("| [Index+1] Add New Mission", strings.Repeat(" ", 4), "|")
	fmt.Println("+", strings.Repeat("-", 30), "+")
	for i := 0; i < n+1; i++ {
		fmt.Printf("\t[%d] %s\n", i+1, m[i].name)
	}
	fmt.Println("\nSelect mission : ")
}

func showSubmenu() {
	fmt.Println("\nMISSIONS OPTIONS \n")
	fmt.Println("[1] Change status to /ongoing/")
	fmt.Println("[2] Change status to /unattended/")
	fmt.Println("[3] Change status to /finished/")
	fmt.Println("\nEnter new status : ")
}

func enterMission() (m []mission, num int) {
	var inputMore string
	for {
		fmt.Printf("Mission name : ")
		name := scanInput()
		fmt.Printf("Mission objective : ")
		objectives := scanInput()
		fmt.Printf("Mission status : ")
		status := scanInput()
		m = append(m, mission{name, objectives, status})
		// fmt.Println(num)
		num++
		fmt.Println("Do you want to continue? (y/n)")
		inputMore = scanInput()
		inputMore = strings.ToLower(inputMore)
		if strings.Contains(inputMore, "n") {
			return m, num - 1
		}
	}
}

func main() {

	var codeword string
	var m, temp []mission
	var n, n1, choice, missionIdx int

	fmt.Println("+", strings.Repeat("-", 41), "+")
	fmt.Println("|  Welcome to the Mission Control System!  |")
	fmt.Println("+", strings.Repeat("-", 41), "+")
	fmt.Println("\nPlease enter the codeword : ")
	fmt.Scan(&codeword)
	if codeword == "deepamaam" {
		for {
			if m == nil {
				err := errors.New("No mission found!\n")
				fmt.Println("\nERROR : ", err)
				m, n = enterMission()
				// fmt.Println(n)
			}
			showMenu(m, n)
			fmt.Scan(&missionIdx)
			if missionIdx == 0 {
				os.Exit(0)
			} else if missionIdx-1 == n+1 {
				temp, n1 = enterMission()
				m = append(m, temp...)

				if n1 == 0 {
					n += 1
				} else {
					n = n + n1
				}
				// fmt.Println(n1)
			} else if missionIdx > n {
				err := errors.New("Wrong Input AGENT\n")
				fmt.Println("\nERROR : ", err)
			} else {
				missionIdx = missionIdx - 1
				showSubmenu()
				fmt.Scan(&choice)
				switch choice {
				case 1:
					m[missionIdx].status = "ongoing"
				case 2:
					m[missionIdx].status = "unattended"
				case 3:
					m[missionIdx].status = "finished"
				default:
					fmt.Println("\nInvalid choice!")
				}
			}
		}
	} else {
		panic("Intruder!!")
	}
}

/*
index Output
[0]   1
[1]   2
[2]   3
[3]   4
-->   5
*/
