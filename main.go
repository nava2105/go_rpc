package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type Response struct {
	Message string `json:"message"`
}

func sayHelloHandler(w http.ResponseWriter, r *http.Request) {
	// Define la respuesta JSON
	response := Response{"Hola Mundo desde RPC en Go"}

	// Configura la cabecera HTTP para JSON
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)

	// Codifica la respuesta a JSON
	json.NewEncoder(w).Encode(response)
}

func main() {
	http.HandleFunc("/", sayHelloHandler)

	port := ":8080"
	fmt.Println("Servidor escuchando en http://localhost" + port)
	err := http.ListenAndServe(port, nil)
	if err != nil {
		fmt.Println("Error al iniciar el servidor:", err)
	}
}
