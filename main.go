package main

import (
	// "fmt"
	"html/template"
	"log"
	"net/http"
)

var templates = template.Must(template.ParseGlob("templates/*"))

func main() {
	http.HandleFunc("/", Index)
	log.Println("Servidor Corriendo...")
	http.ListenAndServe(":80", nil)
}

func Index(w http.ResponseWriter, r *http.Request) {
	// fmt.Fprint(w, "Prueba Base de Go HTTP")
	templates.ExecuteTemplate(w, "app", nil)
}
