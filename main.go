package main

import (
	"fmt"
	"math/rand"
	"time"
)

func processo(min, max int) (int, string) {
	temp := rand.Intn(max-min+1) + min
	time.Sleep(time.Duration(temp) * time.Second)
	texto := fmt.Sprintf("levou %d segundos", temp)
	return temp, texto

}

func main() {

	TempoTotal := 0
	var qtd_processos int
	fmt.Print("informe uma quantidade de processos: ")
	fmt.Scanln(&qtd_processos)

	var min int
	fmt.Print("informe tempo mínimo de cada processo em segundos: ")
	fmt.Scanln(&min)

	var max int
	fmt.Print("informe tempo máximo de cada processo em segundos: ")
	fmt.Scanln(&max)

	i := 0
	for i < qtd_processos {
		processo, texto := processo(min, max)
		fmt.Println("processo #", i+1, texto)
		TempoTotal += processo
		i++
	}

	fmt.Println("tempo total:", TempoTotal, "segundos")

}
