package main

import "fmt"

func main()  {
	exponential(2)
}

func exponential(numero int){
	if numero > 1000000{
		return
	}
	fmt.Println(numero)
	exponential(numero * 2)
}
