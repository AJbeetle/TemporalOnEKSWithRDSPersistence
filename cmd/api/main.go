package main

import (
	"encoding/json"
	"log"
	"net/http"

	"learningTemp/internal/config"
	"learningTemp/internal/order"
	"learningTemp/internal/temporal"
)

func main() {

	cfg := config.Load()
	c, err := temporal.NewClient(cfg)
	if err != nil {
		log.Fatal(err)
	}

	orderService := order.NewService(c)

	http.HandleFunc("/order", func(w http.ResponseWriter, r *http.Request) {

		if r.Method != http.MethodPost {
			http.Error(w, "Only POST allowed", http.StatusMethodNotAllowed)
			return
		}

		var req order.OrderRequest
		if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
			http.Error(w, "Invalid request body", 400)
			return
		}

		err := orderService.CreateOrder(r.Context(), req.OrderID)
		if err != nil {
			http.Error(w, err.Error(), 500)
			return
		}

		w.Write([]byte("Order workflow started"))
	})

	log.Println("API running on :", cfg.Port)
	http.ListenAndServe(cfg.Port, nil)
}
