package main

import "fmt"

func main(){
    for_mode1()
    for_mode2()
    for_mode3()
    for_mode4()
    for_mode5()
}

func for_mode1(){
    i:=1
    for i < 10 {
        fmt.Println("Mode1: ",i)
        i++
    }
}

func for_mode2(){
    for i :=1; i <= 10; i++ {
        fmt.Println("Mode2: ",i)
        i++
    }
}

func for_mode3(){
    numero:=0
    for{
        fmt.Println("Mode3: ","Continuo")
        fmt.Println("Mode3: ","Ingrese el número secreto:")
        fmt.Scanln(&numero)

        if numero == 100{
            break
        }

    }
}

func for_mode4(){
    var i int = 0
    for i < 10 {
        fmt.Printf("Mode4: ","\n Valor de i: 100",i)
        if(i == 5){
            fmt.Printf("Mode4: ","Multiplicamos por 2 \n",i)
            i = i*2
            continue
        }
        fmt.Printf("Mode4: ","Paso por aqui \n",i)
        i++
    }
}

func for_mode5(){
    var i int = 0

    RUTINA:
        for i < 10 {
            if i == 4 {
                i = i + 2
                fmt.Println("Mode5: ","Rutina sumando 2")
                goto RUTINA
            }
            fmt.Println("Mode5: ","Valor de i: ",i)
            i++
        }
}