package main

import "fmt"

func main(){
	fmt.Println(uno(5))

	numero, estado := dos(2)
	fmt.Println("Valor de numero: ",numero)
	fmt.Println("Valor de estado: ",estado)

	fmt.Println(uno(3))


	fmt.Println(Calculo(5,46))
	fmt.Println(Calculo(1,4,56,32,3))
	fmt.Println(Calculo(34,5,23))
	fmt.Println(Calculo(44))

}

func uno(numero int) int {
	return numero * 2
}

func dos(numero int) (int, bool) {
	if(numero == 1) {
		return 5,false
	}else{
		return 10,true
	}
}

func tres(numero int) (z int) {
	z = numero * 2
	return
}

func Calculo(numero ...int) int{
	total :=0

	for _, num :=  range numero{
		total = total + num
	}
	return total
}