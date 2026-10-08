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

func processo(min, max int) int {
	temp := rand.Intn(max-min+1) + min
	time.Sleep(time.Duration(temp) * time.Second)
	return temp
}

type Resultado struct {
	ID      int
	Inicio  int
	Fim     int
	Duracao float64
}

var wg sync.WaitGroup //organiza pra esperar todoas as concorrencias terminarem
var mu sync.Mutex     //Mutual exclusion (evita "bifurcar" variáveis gerando erro na totalização)

func main() {
	db, err := sql.Open("duckdb", "dados.duckdb")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	_, err = db.Exec(`DELETE FROM processos`)
	if err != nil {
		log.Fatal(err)
	}

	_, err = db.Exec(`
		CREATE TABLE IF NOT EXISTS processos(
			db_id INTEGER,
			db_inicio INTEGER,
			db_fim INTEGER,
			db_duracao BIGINT,
			db_tempo_decorrido DOUBLE,
		)
	`)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Conectou ao DuckDB e criou banco de dados!")
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
			TempoProcesso := processo(min, max)
			mu.Lock()
			TempoProcessos += TempoProcesso
			resultado := Resultado{
				ID:      i + 1,
				Inicio:  TempoProcessos - TempoProcesso,
				Duracao: float64(TempoProcesso),
				Fim:     TempoProcessos,
			}
			TempoDecorrido := time.Since(t0).Seconds()
			_, err := db.Exec(`
				INSERT INTO processos(
				db_id,
				db_inicio,
				db_fim,
				db_duracao,
				db_tempo_decorrido)
				VALUES(?,?,?,?,?)
				`,
				resultado.ID,
				resultado.Inicio,
				resultado.Fim,
				int64(resultado.Duracao),
				TempoDecorrido,
			)
			if err != nil {
				fmt.Print("_ - * E R R O  N A L I N H A", resultado.ID, " * - _")
				log.Fatal(err)
			}

			fmt.Println("processo #", resultado.ID, ":", resultado.Inicio, "+", resultado.Duracao, "=", resultado.Fim)
			fmt.Println(" Timestamp:", TempoDecorrido)
			mu.Unlock()
			<-limite

		}(i)
	}

	wg.Wait()
	duracao := time.Since(t0).Seconds()
	fmt.Println("\nProcessos:", TempoProcessos, "segundos")
	fmt.Printf("Duração: %.0f segundos\n", duracao)
	delta := duracao - float64(TempoProcessos)

	switch {
	case delta < -0.9:
		fmt.Printf("economizamos %.0f segundos", -delta)
	case delta > 0.9:
		fmt.Printf("Execução levou %.0f segundos a mais que o tempo dos processos ", delta)
	default:
		fmt.Println("empate")
	}

	_, err = db.Exec(`
	COPY (SELECT
            db_id,
            db_inicio,
            db_fim,
            db_duracao,
            REPLACE(
                CAST(ROUND(db_tempo_decorrido, 3) AS VARCHAR),
                '.',
                ','
            ) AS db_tempo_decorrido
        FROM processos
    )
	TO 'processos.csv'
	(	
		HEADER,
	 	DELIMITER ';'
		)
	`)

	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("\nCSV criado com sucesso!")

	fmt.Println("\n ")
}
