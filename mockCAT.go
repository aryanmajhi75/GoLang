package main

import "fmt"

func buy(amt float64) float64{
  var change float64
  change=amt-40
  return change
}

func main(){
  
  var choice int
  var change,amt float64

  fmt.Println(" [1] Coke  ")
  fmt.Println(" [2] Pepsi ")
  fmt.Println(" [3] Water ")
  fmt.Println("\nEnter your choice : ")
  fmt.Scan(&choice)
  switch(choice){
  case 1:
    fmt.Println("Enter the amount :" )
    fmt.Scan(&amt)
    change=buy(amt)
    if(change>0){
      fmt.Printf("Change : %v",change)
    }

  case 2:
  fmt.Println("Enter the amount : ")
  fmt.Scan(&amt)
  change=buy(amt)
  if(change>0){
      fmt.Printf("Your change : %v",change)
    }
  case 3:
    fmt.Println("No need to pay for water!")
}
}
