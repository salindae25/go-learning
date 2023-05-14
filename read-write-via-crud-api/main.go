package main

import (
	"encoding/json"
	"io"
	"strconv"

	// "fmt"
	"log"
	"net/http"
	"os"
)

func main() {
	http.HandleFunc("/", sayHello)
	http.HandleFunc("/product/", handlePost)

	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		log.Panic(err)
	}
}

func sayHello(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Hello World!"))
}

func handlePost(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodPost:
		addNewProduct(w, r)
		break
	case http.MethodGet:
		getAllTheProducts(w, r)
		break
	case http.MethodPut:
		updateProductById(w, r)
		break
	case http.MethodDelete:
		deleteProductById(w, r)
		break
	default:
		w.Write([]byte("Method not allowed"))
	}
}

type Product struct {
	Id          int     `json:"id"`
	Name        string  `json:"name"`
	Price       float32 `json:"price"`
	Description string  `json:"description"`
	Category    string  `json:"category"`
}
type ProductList struct {
	Products []Product `json:"products"`
}

func readDb() ProductList {
	data, err := os.ReadFile("./db.json")
	if err != nil {
		log.Printf("%v\n", err)
	}
	pl := ProductList{}
	err = json.Unmarshal(data, &pl)
	if err != nil {
		log.Printf("%v\n", err)
	}
	return pl
}
func (pl *ProductList) writeDb() {
	data, err := json.Marshal(pl)
	if err != nil {
		log.Printf("%v\n", err)
	}
	err = os.WriteFile("./db.json", data, 0644)
	if err != nil {
		log.Printf("%v\n", err)
	}
}
func addNewProduct(w http.ResponseWriter, r *http.Request) {
	var products = readDb()
	data, err := io.ReadAll(r.Body)
	if err != nil {

		log.Printf("%v\n", err)
	}
	newProd := Product{}
	err = json.Unmarshal(data, &newProd)
	if err != nil {
		log.Printf("%v\n", err)
	}
	products.Products = append(products.Products, Product{
		Id:          len(products.Products) + 1,
		Name:        newProd.Name,
		Description: newProd.Description,
		Price:       newProd.Price,
		Category:    newProd.Category,
	})
	products.writeDb()
	w.Header().Add("content-type", "application/json")
	w.Write([]byte(`{"status":"Success"}`))
}
func getAllTheProducts(w http.ResponseWriter, r *http.Request) {
	data, err := os.ReadFile("./db.json")
	if err != nil {
		log.Printf("%v\n", err)
	}
	w.Header().Add("content-type", "application/json")
	w.Write(data)
}
func deleteProductById(w http.ResponseWriter, r *http.Request) {
	var products = readDb()
	id := (r.URL.Path[len("/product/"):])
	idI, err := strconv.Atoi(id)
	if err != nil {
		log.Printf("%v\n", err)
		w.Header().Add("content-type", "application/json")
		w.Write([]byte(`{"status":"Failed"}`))
	}
	indexMap := make(map[int]int)
	for i, item := range products.Products {
		indexMap[item.Id] = i
	}
	arr := products.Products
	if i, ok := indexMap[idI]; ok {
		arr = append(arr[:i], arr[i+1:]...)
	}
	products.Products = arr
	products.writeDb()
	w.Header().Add("content-type", "application/json")
	w.Write([]byte(`{"status":"Success"}`))
}
func updateProductById(w http.ResponseWriter, r *http.Request) {

	var products = readDb()
	data, err := io.ReadAll(r.Body)
	if err != nil {

		log.Printf("%v\n", err)
	}
	updateProd := Product{}
	err = json.Unmarshal(data, &updateProd)
	if err != nil {
		log.Printf("%v\n", err)
	}
	for _, item := range products.Products {
		if item.Id == updateProd.Id {
			if updateProd.Name != "" {
				item.Name = updateProd.Name
			}
			if updateProd.Description != "" {
				item.Description = updateProd.Description
			}
			if updateProd.Price != 0 {
				item.Price = updateProd.Price
			}
			if updateProd.Category != "" {
				item.Category = updateProd.Category
			}
		}
	}
	products.writeDb()
	w.Header().Add("content-type", "application/json")
	w.Write([]byte(`{"status":"Success"}`))
}
