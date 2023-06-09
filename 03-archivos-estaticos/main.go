package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
)

// Estructuras
type Usuario struct {
	UserName string
	Edad     int
}

// var templates = template.Must(template.New("T").ParseGlob("templates/*.html"))
var templates = template.Must(template.New("T").ParseGlob("templates/**/*.html"))

// Error template
var errorTemplate = template.Must(template.ParseFiles("templates/error/error.html"))

// Handlear Error
func handlerError(rw http.ResponseWriter, status int) {
	rw.WriteHeader(status)
	errorTemplate.Execute(rw, nil)
}

// Funcion de Render Template
func renderTemplate(rw http.ResponseWriter, name string, data interface{}) {
	err := templates.ExecuteTemplate(rw, name, data)

	if err != nil {
		//http.Error(rw, "No es posible retorna el template", http.StatusInternalServerError)
		//errorTemplate.Execute(rw, nil)
		handlerError(rw, http.StatusInternalServerError)
	}
}

// Handlers
func Index(rw http.ResponseWriter, r *http.Request) {
	usuario := Usuario{"Thomas", 19}
	renderTemplate(rw, "index.html", usuario)
}

func Registro(rw http.ResponseWriter, r *http.Request) {
	renderTemplate(rw, "registro.html", nil)
}

// Función principal
func main() {

	//Archivos statico
	staticFile := http.FileServer(http.Dir("static"))

	//Mux
	mux := http.NewServeMux()
	mux.HandleFunc("/", Index)
	mux.HandleFunc("/registro", Registro)

	//Mux de static file
	mux.Handle("/static/", http.StripPrefix("/static/", staticFile))

	//Server
	server := &http.Server{
		Addr:    "localhost:8080",
		Handler: mux,
	}
	fmt.Println("El servidor esta corriendo en puerto 8080")
	fmt.Println("Run server: http://localhost:8080")
	log.Fatal(server.ListenAndServe())
}
