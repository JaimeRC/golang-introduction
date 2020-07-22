package main

import "fmt"

var Calculo func(int,int) int = func(num1 int,num2 int) int{
	return num1 + num2
}


func main()  {
	fmt.Printf("Sumo 5 + 7 = %d \n ",Calculo(5,7))


	Calculo =  func(num1 int , num2 int) int {
		return num1 - num2
	}
	fmt.Printf("Restamos 5 + 7 = %d \n ",Calculo(5,7))


	Calculo =  func(num1 int , num2 int) int {
		return num1 / num2
	}
	fmt.Printf("Dividimo 5 / 7 = %d \n ",Calculo(12,3))

	Operaciones()

	// Closures
	tablaDel := 2
	MiTabla := Tabla(tablaDel)

	for i:=1; i < 11; i++{
		fmt.Println("Clousure: ",MiTabla())
	}
}



func Operaciones(){
	resultado := func() int {
		var a = 23
		var b = 13

		return a + b
	}

	fmt.Println("Operaciones: ",resultado())
}

func Tabla(valor int) func() int{
	numero := valor
	secuencia := 0

	return func()int{
		secuencia += 1
		return secuencia * numero
	}
}