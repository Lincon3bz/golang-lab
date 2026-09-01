package main

import (
	"fmt"
	"math/rand"
)

func processo() (int, string) {
	temp := rand.Intn(9) + 2
	texto := fmt.Sprintf("Esse processo levou %d segundos", temp)
	return temp, texto

}

func main() {

	TempoTotal := 0
	informado := fmt.Scanln("informe um número")
	fmt.Println(informado)
	processo1, texto := processo()
	fmt.Println(texto)
	TempoTotal += processo1

	processo2, texto := processo()
	fmt.Println(texto)
	TempoTotal += processo2

	processo3, texto := processo()
	fmt.Println(texto)
	TempoTotal += processo3

	processo4, texto := processo()
	fmt.Println(texto)
	TempoTotal += processo4

	fmt.Println("tempo total:", TempoTotal, "segundos")

}
