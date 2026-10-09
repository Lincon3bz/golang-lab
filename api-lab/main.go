package main

import (
	"encoding/json"
	"log"
	"net/http"
)

type Resposta struct {
	Status string `json:"status"`
}

type Request_Soma struct {
	A int `json:"a"`
	B int `json:"b"`
}

type Response_Soma struct {
	Resultado int `json:"resultado"`
}

func status_Handler(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Content-Type", "application/json")

	resposta := Resposta{
		Status: "API funcionando",
	}

	json.NewEncoder(w).Encode(resposta)

}

func soma_Handler(w http.ResponseWriter, r *http.Request) {
	var entrada Request_Soma

	err := json.NewDecoder(r.Body).Decode(&entrada)

	if err != nil {
		http.Error(w, "JSON inválido", http.StatusBadRequest)
		return
	}

	resultado := Response_Soma{
		Resultado: entrada.A + entrada.B,
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(resultado)
}

func main() {
	mux := http.NewServeMux()

	mux.HandleFunc("GET /", func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Path != "/" {
			http.NotFound(w, r)
			return
		}

		http.ServeFile(w, r, "api-lab/index.html")
	})

	mux.HandleFunc("GET /api/status", status_Handler)
	mux.HandleFunc("POST /api/soma", soma_Handler)

	log.Println("Servidor iniciado em http://localhost:8080")

	err := http.ListenAndServe(":8080", mux)

	if err != nil {
		log.Fatal(err)
	}
}
