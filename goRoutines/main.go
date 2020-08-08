package main

import (
	"fmt"
	"strings"
	"time"
)

func main(){
	go myNameSlow("Pepito de los Palotes")

	fmt.Println("Next")
	var x string
	fmt.Scanln(&x)
}

func myNameSlow(name string){
	letters := strings.Split(name,"")

	for _, letter := range letters{
		time.Sleep(1000 * time.Millisecond)
		fmt.Println(letter)
	}
}
