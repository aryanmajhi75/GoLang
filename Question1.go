package main

import "fmt"

//function to check and recommend clothes
func recommend(temp float32) string {

//variable to return the recommendation
var recom string

//if-else statements
if (temp<10.0){
	recom="Wear heavy jacket"
}else if (temp>10 && temp<20){
	recom="Wear light jacket"
}else if (temp>=20.0){
	recom="Wear a t-shirt"
}else{
	recom="IDK"
}
return recom	//returning the value "recom"
}

func main(){

//variables for temperature
var tem float32

fmt.Println("Enter the temperature(in Celsius) of your area: ")
fmt.Scanf("%v",&tem)
fmt.Println("Recommendation: ",recommend(tem))//calling the function to recommend
}
