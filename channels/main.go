package main

import (
	"fmt"
	"time"
)

func main(){
	chanel1 := make(chan time.Duration)
	go loop(chanel1)
	fmt.Println("Next")

	msg := <- chanel1
	fmt.Println("Channel: ",msg)

}


func loop(channel chan time.Duration){
	start := time.Now()
	for i := 0; i<1000000000;i++{
	}

	end := time.Now()

	channel <- end.Sub(start)
}