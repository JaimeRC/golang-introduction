package main

import (
	"fmt"
	)


func main(){
	interface_mode1()
	interface_mode2()
	interface_mode3()
}

type serVivo interface {
	estaVivo() bool
}

type humano interface {
	respirar()
	pensar()
	comer()
	sexo() string
	estaVivo() bool
}

type animal interface {
	respirar()
	comer()
	esCarnivoro() bool
	estaVivo() bool
}

type vegetal interface {
	clasificacionVegetal() string
}

/* Genero Humano */
type hombre struct {
	edad int
	altura float32
	peso float32
	respirando bool
	pensando bool
	comiendo bool
	esHombre bool
	vivo bool
}

type mujer struct {
	hombre
}

func (h *hombre) respirar(){h.respirando = true}
func (h *hombre) comer(){h.comiendo = true}
func (h *hombre) pensar(){h.pensando = true}
func (h *hombre) sexo() string{
	if h.esHombre == true{
		return "Hombre"
	}else{
		return "Mujer"
	}
}
func (h *hombre) estaVivo() bool{ return true}


func HumanosRespirando(hu humano){
	hu.respirar()
	fmt.Printf("Soy un/a %s, y ya estoy respirando \n", hu.sexo())
}

func interface_mode1(){
	Pedro := new(hombre)
	Pedro.esHombre = true
	HumanosRespirando(Pedro)

	Maria := new(mujer)

	HumanosRespirando(Maria)
}


/* Genero Animal */
type perro struct {
	respirando bool
	comiendo bool
	carnivoro bool
	vivo bool
}

func (p *perro) respirar(){p.respirando = true}
func (p *perro) comer(){p.comiendo = true}
func (p *perro) esCarnivoro() bool{ return p.carnivoro}
func (p *perro) estaVivo() bool{ return true}

func AnimalesRespirar(an animal){
	an.respirar()
	fmt.Printf("Soy un animal y estoy respirando")
}

func AnimalesCarnivoros(an animal) int{
	if an.esCarnivoro() == true {
		return 1
	}
	return 0
}

func interface_mode2(){
	totalCarnivoros := 0
	Dogo := new(perro)
	Dogo.carnivoro = true
	AnimalesRespirar(Dogo)
	totalCarnivoros =+AnimalesCarnivoros(Dogo)

	fmt.Printf("Total carnivoros %d \n",totalCarnivoros)
}

func estoyVIvo(v serVivo) bool{
	return v.estaVivo()
}

func interface_mode3(){
	totalCarnivoros := 0
	Dogo := new(perro)
	Dogo.carnivoro = true
	Dogo.vivo = true
	AnimalesRespirar(Dogo)
	totalCarnivoros =+AnimalesCarnivoros(Dogo)

	fmt.Printf("Total carnivoros %d \n",totalCarnivoros)

	fmt.Printf("Estoy vivo = %t",estoyVIvo(Dogo))
}