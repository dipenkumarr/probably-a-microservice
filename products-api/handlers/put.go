package handlers

import (
	"net/http"
	"strconv"

	"github.com/dipenkumarr/probably-a-microservice/data"
	"github.com/gorilla/mux"
)

func (p *Products) UpdateProduct(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, _ := strconv.Atoi(vars["id"])

	p.l.Println("Handle PUT Product ", id)

	prod := r.Context().Value(KeyProduct{}).(data.Product)

	err := data.UpdateProduct(id, &prod)
	if err == data.ErrProductNotFound {
		http.Error(w, "Product not found", http.StatusNotFound)
		return
	}
	if err != nil {
		http.Error(w, "Wrong Id", http.StatusBadRequest)
		return
	}

}
