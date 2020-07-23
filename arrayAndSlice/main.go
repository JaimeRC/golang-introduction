package main

import "fmt"


func main(){
	venctor_mode1()
	vector_mode2()
	vector_mode3()
	vector_mode4()
	vector_mode5()
}

func venctor_mode1(){
	var tabla [10]int

	tabla[0]= 1
	tabla[5]= 15

	fmt.Println(tabla)
}

func vector_mode2(){
	tabla := [10]int {5,6,98,1,4,9,3,65,8,99}
	fmt.Println(tabla)
}

func vector_mode3(){
	tabla := [10]int{5,6,98,1,4,9,3,65,8,99}

	for i := 0; i < len(tabla); i++{
		fmt.Println(tabla[i])
	}
}
func vector_mode4(){
	var matrix [5][7]int

	matrix[3][5] = 1
	fmt.Println(matrix)
}

func vector_mode5(){
	matrix := []int{2,5,4}
	fmt.Println(matrix)

	variante2()
	variante3()
	variante4()
}

func variante2(){
	elementos := [5]int{1,2,3,4,5}
	porcion := elementos[3:]
	fmt.Println(porcion)
}

func variante3(){
	elementos := make([]int,5,20)

	fmt.Printf("Largo %d, Capacidad %d",len(elementos),cap(elementos))
}

func variante4(){
	nums := make([]int,0,0)

	for i := 0; i < 100; i++{
		nums = append(nums,i)
	}
	fmt.Printf("\n Largo %d, Capacidad %d",len(nums),cap(nums))
}