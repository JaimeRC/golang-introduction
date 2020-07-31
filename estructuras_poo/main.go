package main

import (
	"fmt"
	"time"

	us "./user"
)

func main(){
	poo_mode1()
	poo_mode2()
}

type usuario struct {
	Id int
	Nombre string
	FechaAlta time.Time
	Status bool
}

func poo_mode1(){
	User := new(usuario)
	User.Id = 10
	User.Nombre = "Jaime"

	fmt.Println(User)
}

type alumno struct{
	us.Usuario
}

func poo_mode2(){
	user := new(alumno)

	user.AltaUsuario(11,"Pepe",time.Now(),true)

	fmt.Println(user.Usuario)
}