package main

import (
	"net/http"

	"github.com/go-chi/chi/v5"
)

func main() {
	r := chi.NewRouter()// O Chi facilita a criação das rotas
	r.Get("/{productName}", func(w http.ResponseWriter, r *http.Request){
		
		param := chi.URLParam( r, "productName" )

		w.Write([]byte(param))// O Write é responsável por escrever o conteúdo no body da requisição
	})
	http.ListenAndServe("localhost:8080", r) // O ListenAndServe é responsável por iniciar o servidor
}