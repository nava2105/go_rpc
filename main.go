package main

import (
	"encoding/json"
	"fmt"
	"net/http"

	httpSwagger "github.com/swaggo/http-swagger"
	_ "go_rpc/docs"
)

// Response JSON response
type Response struct {
	Message string `json:"message"`
}

// sayHelloHandler Returns a welcome message
// @Summary Returns a welcome message
// @Description Return a JSON object with a welcome message from rpc.
// @Tags Welcome message
// @Produce json
// @Success 200 {object} Response
// @Router / [get]
func sayHelloHandler(w http.ResponseWriter, r *http.Request) {
	response := Response{"Hello World from RPC in Go"}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)

	json.NewEncoder(w).Encode(response)
}

// @title RPC API
// @version 1.0
// @description Simple Go API that returns a message in JSON.
// @host localhost:8080
// @BasePath /
func main() {
	http.HandleFunc("/", sayHelloHandler)

	// Path to serve Swagger UI
	http.Handle("/docs/", httpSwagger.WrapHandler)

	port := ":8080"
	fmt.Println("Server running in: http://localhost" + port)
	fmt.Println("Swagger documentation in: http://localhost" + port + "/docs")
	err := http.ListenAndServe(port, nil)
	if err != nil {
		fmt.Println("Error al iniciar el servidor:", err)
	}
}
