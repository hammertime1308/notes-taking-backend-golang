package controllers

import (
	"encoding/json"
	"net/http"
)

func (c *controller) pingController(w http.ResponseWriter, r *http.Request) {
	response := map[string]string{"status": "running"}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}
