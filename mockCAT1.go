package main

import "fmt"

func main(){
  var temps[7] float64
  var avgTemp float64
  var threshold = 30.0

  fmt.Println("Enter the temperature of this week")
  for i:=0;i<7;i++{
    fmt.Printf("\nEnter temperature of day %v : ",i+1)
    fmt.Scan(&temps[i])
    avgTemp=avgTemp+temps[i]
  }
  avgTemp=avgTemp/7
  fmt.Printf("\nAverage temperature is : %v `C",avgTemp)
  for i:=0;i<7;i++{
    if(temps[i]>threshold){
      fmt.Printf("\nDay %v has above threshold temperature",i+1)
    }
  }
}
