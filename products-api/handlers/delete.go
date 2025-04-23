package handlers

import (
	"net/http"
	"strconv"

	"github.com/dipenkumarr/probably-a-microservice/data"
	"github.com/gorilla/mux"
)

func (p *Products) DeleteProduct(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, _ := strconv.Atoi(vars["id"])

	p.l.Println("Handle DELETE Product ", id)

	err := data.DeleteProduct(id)

	if err == data.ErrProductNotFound {
		http.Error(w, "Product not found", http.StatusNotFound)
		return
	}
	if err != nil {
		http.Error(w, "Wrong Id", http.StatusBadRequest)
		return
	}

}
