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
	t0 := time.Now()
	TempoProcessos := 0
	var qtd_processos int
	fmt.Print("informe uma quantidade de processos: ")
	fmt.Scanln(&qtd_processos)
	for qtd_processos < 1 {
		fmt.Print("quantidade de processos deve ser maior que zero: ")
		fmt.Scanln(&qtd_processos)
	}

	var min int
	fmt.Print("informe tempo mínimo de cada processo em segundos: ")
	fmt.Scanln(&min)

	var max int
	for max < min {
		fmt.Print("informe tempo máximo de cada processo em segundos: ")
		fmt.Scanln(&max)
	}
	fmt.Println()
	i := 0
	for i < qtd_processos {
		processo, texto := processo(min, max)
		fmt.Println("processo #", i+1, texto)
		TempoProcessos += processo
		i++
	}

	duracao := time.Since(t0).Seconds()
	fmt.Println("\nProcessos:", TempoProcessos, "segundos")
	fmt.Printf("Duração: %.0f segundos", duracao)
	if duracao > float64(TempoProcessos) {
		fmt.Printf("\nTempo do processo mais %.0f pra rodar tudo", duracao-float64(TempoProcessos))
	} else {
		fmt.Println("Outro cenário de economia do tempo")
	}
	fmt.Println("\n ")

}
