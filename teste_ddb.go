package main

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/duckdb/duckdb-go/v2"
)

//func main() {

	// ========================================
	// CONEXÃO
	// ========================================

	fmt.Println("========================================")
	fmt.Println("CONECTANDO AO DUCKDB")
	fmt.Println("========================================")

	db, err := sql.Open("duckdb", "crud.db")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	fmt.Println("Conexão realizada com sucesso!")
	fmt.Println()

	// ========================================
	// CREATE
	// ========================================

	fmt.Println("========================================")
	fmt.Println("CREATE - CRIANDO TABELA")
	fmt.Println("========================================")

	_, err = db.Exec(`
		CREATE TABLE IF NOT EXISTS clientes (
			id INTEGER PRIMARY KEY,
			nome VARCHAR,
			idade INTEGER
		)
	`)

	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Tabela 'clientes' criada!")
	fmt.Println()

	// ========================================
	// INSERT
	// ========================================

	fmt.Println("========================================")
	fmt.Println("INSERT - INSERINDO DADOS")
	fmt.Println("========================================")

	clientes := []struct {
		id    int
		nome  string
		idade int
	}{
		{1, "João", 25},
		{2, "Maria", 30},
		{3, "Carlos", 40},
	}

	for _, cliente := range clientes {

		_, err = db.Exec(`
			INSERT INTO clientes (id, nome, idade)
			VALUES (?, ?, ?)
		`,
			cliente.id,
			cliente.nome,
			cliente.idade,
		)

		if err != nil {
			log.Fatal(err)
		}

		fmt.Printf(
			"Inserido -> ID: %d | Nome: %s | Idade: %d\n",
			cliente.id,
			cliente.nome,
			cliente.idade,
		)
	}

	fmt.Println()

	// ========================================
	// SELECT
	// ========================================

	fmt.Println("========================================")
	fmt.Println("SELECT - CONSULTANDO DADOS")
	fmt.Println("========================================")

	rows, err := db.Query(`
		SELECT id, nome, idade
		FROM clientes
		ORDER BY id
	`)

	if err != nil {
		log.Fatal(err)
	}

	defer rows.Close()

	for rows.Next() {

		var (
			id    int
			nome  string
			idade int
		)

		err := rows.Scan(&id, &nome, &idade)
		if err != nil {
			log.Fatal(err)
		}

		fmt.Printf(
			"ID: %d | Nome: %s | Idade: %d\n",
			id,
			nome,
			idade,
		)
	}

	fmt.Println()

	// ========================================
	// UPDATE
	// ========================================

	fmt.Println("========================================")
	fmt.Println("UPDATE - ALTERANDO DADOS")
	fmt.Println("========================================")

	_, err = db.Exec(`
		UPDATE clientes
		SET idade = ?
		WHERE id = ?
	`,
		35,
		2,
	)

	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Cliente ID 2 atualizado para idade 35!")
	fmt.Println()

	// ========================================
	// SELECT APÓS UPDATE
	// ========================================

	fmt.Println("========================================")
	fmt.Println("SELECT - VERIFICANDO UPDATE")
	fmt.Println("========================================")

	rows, err = db.Query(`
		SELECT id, nome, idade
		FROM clientes
		ORDER BY id
	`)

	if err != nil {
		log.Fatal(err)
	}

	for rows.Next() {

		var (
			id    int
			nome  string
			idade int
		)

		err := rows.Scan(&id, &nome, &idade)
		if err != nil {
			log.Fatal(err)
		}

		fmt.Printf(
			"ID: %d | Nome: %s | Idade: %d\n",
			id,
			nome,
			idade,
		)
	}

	rows.Close()

	fmt.Println()

	// ========================================
	// DELETE
	// ========================================

	fmt.Println("========================================")
	fmt.Println("DELETE - REMOVENDO DADOS")
	fmt.Println("========================================")

	_, err = db.Exec(`
		DELETE FROM clientes
		WHERE id = ?
	`, 1)

	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Cliente ID 1 removido!")
	fmt.Println()

	// ========================================
	// SELECT FINAL
	// ========================================

	fmt.Println("========================================")
	fmt.Println("SELECT - ESTADO FINAL DA TABELA")
	fmt.Println("========================================")

	rows, err = db.Query(`
		SELECT id, nome, idade
		FROM clientes
		ORDER BY id
	`)

	if err != nil {
		log.Fatal(err)
	}

	defer rows.Close()

	for rows.Next() {

		var (
			id    int
			nome  string
			idade int
		)

		err := rows.Scan(&id, &nome, &idade)
		if err != nil {
			log.Fatal(err)
		}

		fmt.Printf(
			"ID: %d | Nome: %s | Idade: %d\n",
			id,
			nome,
			idade,
		)
	}

	fmt.Println()
	fmt.Println("========================================")
	fmt.Println("CRUD FINALIZADO")
	fmt.Println("========================================")
}
