package main

import (
	"encoding/json"
	"log"
	"net/http"

	"learningTemp/internal/config"
	"learningTemp/internal/order"
	"learningTemp/internal/order/models"
	"learningTemp/internal/temporal"
)

func main() {

	cfg := config.Load()
	c, err := temporal.NewClient(cfg)
	if err != nil {
		log.Fatal(err)
	}

	orderService := order.NewService(c)

	// Health check — ALB pings this
	http.HandleFunc("/health", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(map[string]string{"status": "ok"})
	})

	http.HandleFunc("/order", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")

		if r.Method != http.MethodPost {
			w.WriteHeader(http.StatusMethodNotAllowed)
			json.NewEncoder(w).Encode(map[string]string{"error": "only POST allowed"})
			return
		}

		var req models.OrderRequest
		if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
			w.WriteHeader(http.StatusBadRequest)
			json.NewEncoder(w).Encode(map[string]string{"error": "invalid request body"})
			return
		}

		// Basic validation
		if req.OrderID == "" {
			w.WriteHeader(http.StatusBadRequest)
			json.NewEncoder(w).Encode(map[string]string{"error": "orderID is required"})
			return
		}

		workflowRun, err := orderService.CreateOrder(r.Context(), req.OrderID)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			json.NewEncoder(w).Encode(map[string]string{"error": err.Error()})
			return
		}

		w.WriteHeader(http.StatusAccepted) // 202 — workflow started, not yet complete
		json.NewEncoder(w).Encode(map[string]string{
			"status":     "accepted",
			"orderID":    req.OrderID,
			"workflowID": workflowRun.GetID(),
			"runID":      workflowRun.GetRunID(),
		})
	})

	log.Println("API running on :", cfg.Port)
	if err := http.ListenAndServe(cfg.Port, nil); err != nil {
		log.Fatal("server failed:", err)
	}
}
