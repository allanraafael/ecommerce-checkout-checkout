package main

import (
	"log"
	"net/http"
	"os"

	"github.com/gorilla/mux"
	"github.com/joho/godotenv"
)


var baseUrlProducts string


func init() {
	// Carregando arquivo .env da pasta product
	err := godotenv.Load(".env")

	if err != nil {
		log.Fatalf("Falha ao carregar arquivo .env")
	}

	// Recebe o valor atríbuido a variável PRODUCT_URL=VALOR no arquivo .env 
	baseUrlProducts = os.Getenv("PRODUCT_URL")
}


func displayCheckout(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Estou aqui dentro do checkout"))
}


func main() {
	r := mux.NewRouter()
	r.HandleFunc("/{id}", displayCheckout)
	http.ListenAndServe(":8084", r)
}
