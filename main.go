package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

func processo(min, max int) (int, string) {
	temp := rand.Intn(max-min+1) + min
	time.Sleep(time.Duration(temp) * time.Second)
	texto := fmt.Sprintf("levou %d segundos", temp)
	return temp, texto
}

var t0 = time.Now()
var wg sync.WaitGroup
var mu sync.Mutex

func main() {
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

	t0 := time.Now()
	wg.Add(qtd_processos)
	for i := 0; i < qtd_processos; i++ {
		go func(i int) {
			defer wg.Done()
			processo, texto := processo(min, max)
			mu.Lock()
			fmt.Println("processo #", i+1, texto)
			fmt.Print(TempoProcessos, "+", processo, "=")
			TempoProcessos += processo
			fmt.Print(TempoProcessos, "\n")
			mu.Unlock()

		}(i)
	}

	wg.Wait()
	duracao := time.Since(t0).Seconds()
	fmt.Println("\nProcessos:", TempoProcessos, "segundos")
	fmt.Printf("Duração: %.0f segundos\n", duracao)
	delta := duracao - float64(TempoProcessos)
	//fmt.Printf("dif em segundos: %.2f", delta)

	switch {
	case delta < -0.9:
		fmt.Printf("economizamos %.0f segundos", -delta)
	case delta > 0.9:
		fmt.Printf("Execução levou %.0f segundos a mais que o tempo dos processos ", delta)
	default:
		fmt.Println("empate")
	}

	fmt.Println("\n ")
}
