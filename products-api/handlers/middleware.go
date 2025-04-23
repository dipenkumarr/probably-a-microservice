package handlers

import (
	"context"
	"fmt"
	"net/http"

	"github.com/dipenkumarr/probably-a-microservice/data"
)

func (p *Products) MiddlerwareProductValidation(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		prod := data.Product{}

		err := prod.FromJSON(r.Body)
		if err != nil {
			http.Error(w, "Unable to decode JSON", http.StatusBadRequest)
			return
		}

		// validate
		err = prod.Validate()
		if err != nil {
			http.Error(w, fmt.Sprintf("Error validating product: %s", err), http.StatusBadRequest)
			return
		}

		ctx := context.WithValue(r.Context(), KeyProduct{}, prod)
		req := r.WithContext(ctx)

		next.ServeHTTP(w, req)
	})
}
