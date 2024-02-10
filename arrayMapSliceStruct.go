package main

import "fmt"

func main(){

var i int=0
var monthlyrecord[12] int

fmt.Println("Enter emission data (in KCO2) ")

for i=0; i<12; i++ {
	fmt.Scanf("%v",&monthlyrecord[i])
}

fmt.Println("+---------------+")

for i=0; i<12; i++ {
	fmt.Printf("| Record %v : %v |\n",i,monthlyrecord[i])
}

fmt.Println("+---------------+")
}
