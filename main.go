package main

import (
	"database/sql"
	"fmt"
	"log"
	"math/rand"
	"sync"
	"time"

	_ "github.com/duckdb/duckdb-go/v2"
)

func processo(min, max int) (int, string) {
	temp := rand.Intn(max-min+1) + min
	time.Sleep(time.Duration(temp) * time.Second)
	texto := fmt.Sprintf("levou %d segundos", temp)
	return temp, texto
}

var t0 = time.Now()
var wg sync.WaitGroup //organiza pra esperar todoas as concorrencias terminarem
var mu sync.Mutex     //Mutual exclusion (evita "bifurcar" variáveis gerando erro na totalização)

func main() {
	db, err := sql.Open("duckdb", "dados.duckdb")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()
	fmt.Println("Conectou ao DuckDB!")
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
	var limite_simultaneos int
	fmt.Print("informe limite de proocessos simultâneos: ")
	fmt.Scanln(&limite_simultaneos)
	for limite_simultaneos < 1 {
		fmt.Print("limite de proocessos simultâneos deve ser maior que zero: ")
		fmt.Scanln(&limite_simultaneos)
	}
	limite := make(chan struct{}, limite_simultaneos)

	fmt.Println()

	t0 := time.Now()
	wg.Add(qtd_processos)
	for i := 0; i < qtd_processos; i++ {

		go func(i int) {
			defer wg.Done()
			limite <- struct{}{}
			processo, texto := processo(min, max)
			mu.Lock()
			fmt.Println("processo #", i+1, texto)
			fmt.Print(TempoProcessos, "+", processo, "=")
			TempoProcessos += processo
			fmt.Print(TempoProcessos, " Timestamp:", time.Since(t0).Seconds(), "\n")
			mu.Unlock()
			<-limite

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
