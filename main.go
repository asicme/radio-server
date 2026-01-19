package main

import (
	"fmt"
	"net/http"
)

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Servidor de Radio Online OK")
	})

	fmt.Println("Escuchando en el 8080...")
	http.ListenAndServe(":8080", nil)
}
