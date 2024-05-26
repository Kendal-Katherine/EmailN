package main

import (
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/render"
)

type product struct {
	ID   int
	Name string
}

func main() {
	r := chi.NewRouter() // O Chi facilita a criação das rotas
	r.Get("/{productName}", func(w http.ResponseWriter, r *http.Request) {

		param := chi.URLParam(r, "productName")

		w.Write([]byte(param)) // O Write é responsável por escrever o conteúdo no body da requisição
	})

	r.Post("/product", func(w http.ResponseWriter, r *http.Request) {
		var product product
		render.DecodeJSON(r.Body, &product) // O DecodeJSON é responsável por decodificar o conteúdo do body da requisição em um objeto do tipo product

		product.ID = 5

		render.JSON(w, r, product) // O render é responsável por renderizar o conteúdo no body da requisi
	})

	http.ListenAndServe("localhost:8080", r) // O ListenAndServe é responsável por iniciar o servidor
}
