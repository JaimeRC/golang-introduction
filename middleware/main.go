package main

import (
	"fmt"
)

var result int
func main()  {
	fmt.Println("Start")
	result = operationMidd(add)(2,3)
	fmt.Println(result)
	result = operationMidd(subtract)(10,6)
	fmt.Println(result)
	result = operationMidd(multiply)(2,4)
	fmt.Println(result)
}

func add(a,b int)int  {
	return a+b
}
func subtract(a,b int)int  {
	return a-b
}
func multiply(a,b int )int  {
	return a*b
}


func operationMidd(operation func(int,int) int) func(int,int) int{
	return func(num1 int, num2 int) int {
		return operation(num1,num2)
	}
}