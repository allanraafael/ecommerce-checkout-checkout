package main

import (
	"checkout/checkout/queue"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"text/template"

	"github.com/gorilla/mux"
	"github.com/joho/godotenv"
)


type Product struct {
	Uuid string `json:"uuid"`
	Product string `json:"product"`
	Price float64 `json:"price,string"`
}


type Order struct {
	Name string `json:"name"`
	Email string `json:"email"`
	Phone string `json:"phone"`
	ProductId string `json:"product_id"`
}


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
	vars := mux.Vars(r)

	// Enviando requisição para microsserviço de produto
	response, err := http.Get(baseUrlProducts + "/products/" + vars["id"])
	if err != nil {
		fmt.Printf("Falha ao carregar requisição HTTP %s\n", err)
	}
	data, _ := ioutil.ReadAll(response.Body)

	var product Product
	json.Unmarshal(data, &product)

	t := template.Must(template.ParseFiles("templates/checkout.html"))
	t.Execute(w, product)
}


func finish(w http.ResponseWriter, r *http.Request) {
	var order Order
	order.Name = r.FormValue("name")
	order.Email = r.FormValue("email")
	order.Phone = r.FormValue("phone")
	order.ProductId = r.FormValue("product_id")

	data, _ := json.Marshal(order)
	fmt.Println(string(data))

	queue.Connect()
	w.Write([]byte("Em processamento"))
}


func main() {
	r := mux.NewRouter()
	r.HandleFunc("/finish", finish)
	r.HandleFunc("/{id}", displayCheckout)
	http.ListenAndServe(":8084", r)
}
