package main

import (
	"fmt"
	"html/template"
	"net/http"
)

//Importamos el gestor de plantillas html

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Mi página principal")
	})
	http.HandleFunc("/productos", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Página Producto")

		fmt.Println("URL", r.URL)
		mapa := r.URL.Query()
		mapa.Add("precio", "10")
		fmt.Println(mapa.Encode())
		fmt.Println("Raw Query(Parametros)", r.URL.RawQuery)
		fmt.Println("Id producto", r.URL.Query().Get("idProducto"))
		fmt.Println("Nombre", r.URL.Query().Get("nombre"))
	})
	http.HandleFunc("/categorias", func(w http.ResponseWriter, r *http.Request) {
		//http.Redirect(w, r, "/", 301)
		template, error := template.ParseFiles("../html/index.html")

		if error != nil {
			panic("Ocurrio un error")
		} else {
			template.Execute(w, nil)
		}
	})
	http.HandleFunc("/noencontro", func(w http.ResponseWriter, r *http.Request) {
		http.NotFound(w, r)
	})
	http.HandleFunc("/error", func(w http.ResponseWriter, r *http.Request) {
		http.Error(w, "Error en el servidor", 500)
	})
	http.ListenAndServe(":8000", nil)
}
