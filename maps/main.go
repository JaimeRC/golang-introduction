package main

import "fmt"

func main(){
	map_mode1()
	map_mode2()
}

func map_mode1(){
	paises := make(map[string]string)

	fmt.Println(paises)

	paises["Mexico"] = "D.F."

	fmt.Println(paises["Mexico"])

	paises["Argentina"] = "Buenos Aires"

	fmt.Println(paises)
}

func map_mode2(){
	campeanato := map[string]int{
		"Barcelona": 39,
		"Real Madrid": 38,
		"Chivas": 37,
		"Boca Juniors": 30}

	fmt.Println(campeanato)

	campeanato["River Plate"] = 55

	fmt.Println(campeanato)

	campeanato["Chivas"] = 25

	fmt.Println(campeanato)

	for equipo, puntaje := range campeanato{
		fmt.Printf("El equipo %s, tiene un puntaje de : %d \n",equipo,puntaje)
	}

	puntaje, existe := campeanato["Mineiro"]
	fmt.Printf("El puntaje capturado %d, y el equipo existe %t \n",puntaje,existe)

}