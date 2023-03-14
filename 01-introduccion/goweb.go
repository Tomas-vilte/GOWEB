package main

import (
	"fmt"
	"log"
	"net/http"
)

// Handlers
func Hola(rw http.ResponseWriter, r *http.Request) {
	fmt.Println("El metodo es + " + r.Method)
	fmt.Fprintln(rw, "Hola mundo")
}

func PaginaNF(rw http.ResponseWriter, r *http.Request) {
	http.NotFound(rw, r)
}

func Error(rw http.ResponseWriter, r *http.Request) {
	http.Error(rw, "La pagina no funciona", http.StatusNotFound)
}

func Saludar(rw http.ResponseWriter, r *http.Request) {
	fmt.Println(r.URL)
	fmt.Println(r.URL.RawQuery)
	fmt.Println(r.URL.Query())
	name := r.URL.Query().Get("name")
	age := r.URL.Query().Get("age")

	fmt.Fprintf(rw, "Hola %s, tu edad es %s!!", name, age)
}

func main() {
	// Mux
	mux := http.NewServeMux()
	mux.HandleFunc("/", Hola)
	mux.HandleFunc("/page", PaginaNF)
	mux.HandleFunc("/error", Error)
	mux.HandleFunc("/saludar", Saludar)

	// Router
	//http.HandleFunc("/", Hola)
	//http.HandleFunc("/page", PaginaNF)
	//http.HandleFunc("/error", Error)
	//http.HandleFunc("/saludar", Saludar)

	// Crear servidor
	server := &http.Server{
		Addr:    "localhost:8080",
		Handler: mux,
	}
	fmt.Println("El servidor esta corriendo en el puerto 8080")
	fmt.Println("Run server: http://localhost:8080/")
	//log.Fatal(http.ListenAndServe("localhost:8080", mux))
	log.Fatal(server.ListenAndServe())
}
