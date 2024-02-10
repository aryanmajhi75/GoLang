package main

import "fmt"

func main(){
	const a int=10
	fmt.Println(a)
	//ERROR: cannot assign to a (declared const)
	//const block
	const(
		d int = iota
		b int = iota
		c int =iota
	)
	fmt.Println(b,c,d)
}
