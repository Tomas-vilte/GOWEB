package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
)

// Funciones
func Saludar(nombre string) string {
	return "Hola" + nombre + "desde una funcion"
}

// Estructuras
type Usuarios struct {
	Nombre string
	Edad   int
}

var templates = template.Must(template.New("T").ParseGlob("templates/**/*.html"))
var errorTemplate = template.Must(template.ParseFiles("templates/error/error.html"))

// Handler error
func handleError(rw http.ResponseWriter, status int) {
	rw.WriteHeader(status)
	errorTemplate.Execute(rw, nil)
}

// Funcionde de render template
func renderTemplate(rw http.ResponseWriter, name string, data interface{}) {
	err := templates.ExecuteTemplate(rw, name, data)
	if err != nil {
		//http.Error(rw, "No es posible retornar el template", http.StatusInternalServerError)
		handleError(rw, http.StatusInternalServerError)
	}
}

// Handler
func Index(rw http.ResponseWriter, r *http.Request) {
	//fmt.Fprintln(rw, "Hola mundo")

	//templates := template.Must(template.New("index.html").Funcs(funciones).ParseFiles("index.html", "base.html"))
	usuario := Usuarios{Nombre: "Thomas", Edad: 19}

	//templates.Execute(rw, usuario)
	renderTemplate(rw, "index.html", usuario)
}

func Registro(rw http.ResponseWriter, r *http.Request) {
	renderTemplate(rw, "registro.html", nil)
}

func main() {
	// Mux
	mux := http.NewServeMux()
	mux.HandleFunc("/", Index)
	mux.HandleFunc("/registro", Registro)
	// Servidor
	server := &http.Server{
		Addr:    "localhost:8080",
		Handler: mux,
	}
	fmt.Println("El servidor esta corriendo en el puerto 8080")
	fmt.Println("Run server: http://localhost:8080/")
	log.Fatal(server.ListenAndServe())
}
