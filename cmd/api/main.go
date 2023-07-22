package main

import (
	"encoding/json"
	"net/http"

	"github.com/devfullcycle/go-intensivo-jul/internal/entity"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

func main() {
	r := chi.NewRouter()
	r.Use(middleware.Logger) // middleware de log
	r.Get("/order", Order)   // declarar rotas com validações de metodos
	// http.HandleFunc("/order", Order) // declarar rotas manuais sem validações de metodos
	http.ListenAndServe(":8888", r)
}

func Order(w http.ResponseWriter, r *http.Request) {
	// Validações de metodos manuais

	// if r.Method != http.MethodGet {
	// 	w.WriteHeader(http.StatusMethodNotAllowed)
	// 	json.NewEncoder(w).Encode("Method not allowed")
	// 	return
	// }

	order, err := entity.NewOrder("1", 1000, 10)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(err.Error())
		return
	}
	order.CalculateFinalPrice()
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(order)

}
