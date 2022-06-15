package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/", Index)
	log.Println("Servidor Corriendo...")
	http.ListenAndServe(":80", nil)
}

func Index(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Prueba Base de Go HTTP")
}
