package main

import (
	"fmt"
	"log"
	"os"
)

func main(){
	//defer_mode()
	panic_mode()
}

func defer_mode(){
	file := "test.txt"

	f, err := os.Open(file)

	defer f.Close()

	if err != nil{
		fmt.Println("error open file")
		os.Exit(1)
	}
}

func panic_mode(){
	defer func() {
		reco := recover()
		if reco != nil {
			log.Fatalf("An error occurred causing Panic --> %v", reco)
		}
	}()
	a:=1
	if a==1{
		panic("Panic is called")
	}
}