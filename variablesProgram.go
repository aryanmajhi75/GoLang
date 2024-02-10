package main 

import "fmt"

func main(){
	var a int = 10 //identifier variableName dataType = value
	//var b int = 20 will give error when b is not used in the program
	fmt.Println(a)
	//OR
	var c int
	c=50
	fmt.Println(c)
	//OR
	d:= 100	//not recommmended
	fmt.Printf("%v %T",d,d)
	fmt.Println(d)

	var(
	stu_name string ="Raj"
	stu_marks float32=45.75	//float will ALSO have to use the size i.e., 32 or 64...
	stu_id int=3
	)
	fmt.Printf("%v %T\n",stu_name,stu_name)
	fmt.Printf("%v %T\n",stu_marks,stu_marks)
	fmt.Printf("%v %T\n",stu_id,stu_id)
}
