package handler

import (
	"net/http"

	"github.com/M-2001/GoRestAPI/database"
	"github.com/go-chi/chi"
)

//Handler servira para hacer el ruteo de la metodos
func Handler() {
	r := chi.NewRouter()
	r.Get("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("welcome"))
	})
	r.Get("/products", database.AllProductos)
	r.Post("/products", database.CreateProduct)
	r.Put("/products/{id}", database.UpdateProduct)
	r.Delete("/products/{id}", database.Deleteproduct)
	http.ListenAndServe(":3000", r)
}
